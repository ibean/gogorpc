package transport

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/let-z-go/toolkit/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/let-z-go/gogorpc/internal/proto"
)

func TestOptions(t *testing.T) {
	type PureOptions struct {
		HandshakeTimeout      time.Duration
		MaxHandshakeSize      int
		MinInputBufferSize    int
		MaxInputBufferSize    int
		MaxIncomingPacketSize int
		MaxOutgoingPacketSize int
	}
	makePureOptions := func(opts *Options) PureOptions {
		return PureOptions{
			HandshakeTimeout:      opts.HandshakeTimeout,
			MaxHandshakeSize:      opts.MaxHandshakeSize,
			MinInputBufferSize:    opts.MinInputBufferSize,
			MaxInputBufferSize:    opts.MaxInputBufferSize,
			MaxIncomingPacketSize: opts.MaxIncomingPacketSize,
			MaxOutgoingPacketSize: opts.MaxOutgoingPacketSize,
		}
	}
	{
		opts1 := Options{}
		opts1.Normalize()
		opts2 := Options{
			HandshakeTimeout:      defaultHandshakeTimeout,
			MaxHandshakeSize:      defaultMaxHandshakeSize,
			MinInputBufferSize:    defaultMinInputBufferSize,
			MaxInputBufferSize:    defaultMaxInputBufferSize,
			MaxIncomingPacketSize: defaultMaxPacketSize,
			MaxOutgoingPacketSize: defaultMaxPacketSize,
		}
		assert.Equal(t, makePureOptions(&opts2), makePureOptions(&opts1))
	}
	{
		opts1 := Options{
			HandshakeTimeout:      -1,
			MaxHandshakeSize:      -1,
			MinInputBufferSize:    -1,
			MaxInputBufferSize:    -1,
			MaxIncomingPacketSize: -1,
			MaxOutgoingPacketSize: -1,
		}
		opts1.Normalize()
		opts2 := Options{
			HandshakeTimeout:      minHandshakeTimeout,
			MaxHandshakeSize:      minMaxHandshakeSize,
			MinInputBufferSize:    minInputBufferSize,
			MaxInputBufferSize:    minInputBufferSize,
			MaxIncomingPacketSize: minInputBufferSize,
			MaxOutgoingPacketSize: minMaxPacketSize,
		}
		assert.Equal(t, makePureOptions(&opts2), makePureOptions(&opts1))
	}
	{
		opts1 := Options{
			HandshakeTimeout:      math.MaxInt64,
			MaxHandshakeSize:      math.MaxInt32,
			MinInputBufferSize:    math.MaxInt32,
			MaxInputBufferSize:    math.MaxInt32,
			MaxIncomingPacketSize: math.MaxInt32,
			MaxOutgoingPacketSize: math.MaxInt32,
		}
		opts1.Normalize()
		opts2 := Options{
			HandshakeTimeout:      maxHandshakeTimeout,
			MaxHandshakeSize:      maxMaxHandshakeSize,
			MinInputBufferSize:    maxInputBufferSize,
			MaxInputBufferSize:    maxInputBufferSize,
			MaxIncomingPacketSize: maxMaxPacketSize,
			MaxOutgoingPacketSize: maxMaxPacketSize,
		}
		assert.Equal(t, makePureOptions(&opts2), makePureOptions(&opts1))
	}
}

func TestHandshake1(t *testing.T) {
	testSetup(
		t,
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(&Options{Logger: &logger}, false, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{
				CbSizeHandshake: func() int {
					return len("hello")
				},
				CbEmitHandshake: func(buf []byte) error {
					copy(buf, "hello")
					return nil
				},
				CbHandleHandshake: func(ctx context.Context, rh []byte) (bool, error) {
					if string(rh) != "nick to meet you" {
						return false, nil
					}
					return true, nil
				},
			}.Init())
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.True(t, ok)
			tp.Prepare(DummyTrafficDecrypter{})
		},
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(&Options{Logger: &logger}, true, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{
				CbHandleHandshake: func(ctx context.Context, rh []byte) (bool, error) {
					if string(rh) != "hello" {
						return false, nil
					}
					return true, nil
				},
				CbSizeHandshake: func() int {
					return len("nick to meet you")
				},
				CbEmitHandshake: func(buf []byte) error {
					copy(buf, "nick to meet you")
					return nil
				},
			}.Init())
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.True(t, ok)
			tp.Prepare(DummyTrafficDecrypter{})
		},
	)
}

func TestHandshake2(t *testing.T) {
	testSetup(
		t,
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(&Options{HandshakeTimeout: -1}, false, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{}.Init())
			if !assert.Regexp(t, "i/o timeout", err) {
				t.FailNow()
			}
			assert.False(t, ok)
			tp.Prepare(DummyTrafficDecrypter{})
		},
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(&Options{HandshakeTimeout: -1}, true, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{
				CbHandleHandshake: func(ctx context.Context, rh []byte) (bool, error) {
					<-ctx.Done()
					<-time.After(10 * time.Millisecond)
					return true, ctx.Err()
				},
			}.Init())
			if !assert.EqualError(t, err, "context deadline exceeded") {
				t.FailNow()
			}
			assert.False(t, ok)
			tp.Prepare(DummyTrafficDecrypter{})
		},
	)
}

func TestHandshake3(t *testing.T) {
	testSetup(
		t,
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(&Options{HandshakeTimeout: -1}, false, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{
				CbHandleHandshake: func(ctx context.Context, rh []byte) (bool, error) {
					<-ctx.Done()
					<-time.After(10 * time.Millisecond)
					return false, ctx.Err()
				},
			}.Init())
			if !assert.EqualError(t, err, "context deadline exceeded") {
				t.FailNow()
			}
			assert.False(t, ok)
			tp.Prepare(DummyTrafficDecrypter{})
		},
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(&Options{HandshakeTimeout: -1}, true, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{}.Init())
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.True(t, ok)
			tp.Prepare(DummyTrafficDecrypter{})
		},
	)
}

func TestHandshake4(t *testing.T) {
	testSetup2(
		t,
		&Options{},
		&Options{},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, defaultMaxPacketSize, tp.maxIncomingPacketSize)
			assert.Equal(t, defaultMaxPacketSize, tp.maxOutgoingPacketSize)
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, defaultMaxPacketSize, tp.maxIncomingPacketSize)
			assert.Equal(t, defaultMaxPacketSize, tp.maxOutgoingPacketSize)
		},
	)
	testSetup2(
		t,
		&Options{
			MaxOutgoingPacketSize: minMaxPacketSize + 1,
		},
		&Options{
			MaxIncomingPacketSize: minMaxPacketSize + 100,
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxOutgoingPacketSize)
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxIncomingPacketSize)
		},
	)
	testSetup2(
		t,
		&Options{
			MaxOutgoingPacketSize: minMaxPacketSize + 100,
		},
		&Options{
			MaxIncomingPacketSize: minMaxPacketSize + 1,
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxOutgoingPacketSize)
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxIncomingPacketSize)
		},
	)
	testSetup2(
		t,
		&Options{
			MaxIncomingPacketSize: minMaxPacketSize + 100,
		},
		&Options{
			MaxOutgoingPacketSize: minMaxPacketSize + 1,
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxIncomingPacketSize)
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxOutgoingPacketSize)
		},
	)
	testSetup2(
		t,
		&Options{
			MaxIncomingPacketSize: minMaxPacketSize + 1,
		},
		&Options{
			MaxOutgoingPacketSize: minMaxPacketSize + 100,
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxIncomingPacketSize)
		},
		func(ctx context.Context, tp *Transport) {
			assert.Equal(t, minMaxPacketSize+1, tp.maxOutgoingPacketSize)
		},
	)
}

func TestSendAndReceivePackets(t *testing.T) {
	const N = 1000
	makeEventType := func(i int) proto.EventType {
		switch i % 4 {
		case 0:
			return proto.EVENT_RESPONSE
		case 1:
			return proto.EVENT_REQUEST
		case 2:
			return proto.EVENT_HANGUP
		default:
			return proto.EVENT_KEEPALIVE
		}
	}
	opts1 := Options{}
	opts2 := Options{}
	cb1 := func(ctx context.Context, tp *Transport) {
		m := 1
		for i := 0; i < N; i++ {
			msg := fmt.Sprintf("this packet %d", i)
			err := tp.Write(&Packet{
				Header: proto.PacketHeader{
					EventType: makeEventType(i),
				},
				PayloadSize: len(msg),
			}, func(buf []byte) error {
				copy(buf, msg)
				return nil
			})
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			if i%m == 0 {
				err = tp.Flush(ctx, 0, DummyTrafficEncrypter{})
				if !assert.NoError(t, err, i) {
					t.FailNow()
				}
				m = (m+1)%13 + 1
			}
		}
		err := tp.Flush(ctx, 0, DummyTrafficEncrypter{})
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		tp.Close()
	}
	cb2 := func(ctx context.Context, tp *Transport) {
		pk := Packet{}
		i := 0
		for {
			err := tp.Peek(ctx, 0, DummyTrafficDecrypter{}, &pk)
			if err != nil {
				if !assert.EqualError(t, err, "gogorpc/transport: network: EOF", i) {
					t.FailNow()
				}
				if !assert.Equal(t, i, N) {
					t.FailNow()
				}
				break
			}
			if !assert.Equal(t, makeEventType(i), pk.Header.EventType, i) {
				t.FailNow()
			}
			if !assert.Equal(t, fmt.Sprintf("this packet %d", i), string(pk.Payload)) {
				t.FailNow()
			}
			i++
			for {
				ok, err := tp.PeekNext(&pk)
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				if !ok {
					break
				}
				if !assert.Equal(t, makeEventType(i), pk.Header.EventType, i) {
					t.FailNow()
				}
				if !assert.Equal(t, fmt.Sprintf("this packet %d", i), string(pk.Payload)) {
					t.FailNow()
				}
				i++
			}
		}
	}
	testSetup2(t, &opts1, &opts2, cb1, cb2)
	testSetup2(t, &opts1, &opts2, cb2, cb1)
}

func TestSendBadPacket1(t *testing.T) {
	opts1 := Options{MaxOutgoingPacketSize: -1}
	opts2 := Options{MaxOutgoingPacketSize: -1}
	cb1 := func(ctx context.Context, tp *Transport) {
		err := tp.Write(&Packet{
			Header: proto.PacketHeader{
				EventType: proto.EVENT_REQUEST,
			},
			PayloadSize: minMaxPacketSize,
		}, func([]byte) error { return nil })
		if !assert.EqualError(t, err, ErrPacketTooLarge.Error()) {
			t.FailNow()
		}
	}
	cb2 := func(ctx context.Context, tp *Transport) {
	}
	testSetup2(t, &opts1, &opts2, cb1, cb2)
	testSetup2(t, &opts1, &opts2, cb2, cb1)
}

func TestSendBadPacket2(t *testing.T) {
	errTest := errors.New("my test")
	opts1 := Options{MaxOutgoingPacketSize: -1}
	opts2 := Options{MaxOutgoingPacketSize: -1}
	cb1 := func(ctx context.Context, tp *Transport) {
		err := tp.Write(&Packet{
			Header: proto.PacketHeader{
				EventType: proto.EVENT_REQUEST,
			},
			PayloadSize: 0,
		}, func([]byte) error { return errTest })
		if !assert.EqualError(t, err, errTest.Error()) {
			t.FailNow()
		}
	}
	cb2 := func(ctx context.Context, tp *Transport) {
	}
	testSetup2(t, &opts1, &opts2, cb1, cb2)
	testSetup2(t, &opts1, &opts2, cb2, cb1)
}

type testHandshaker struct {
	CbHandleHandshake func(context.Context, []byte) (bool, error)
	CbSizeHandshake   func() int
	CbEmitHandshake   func([]byte) error
}

func (s testHandshaker) Init() testHandshaker {
	if s.CbHandleHandshake == nil {
		s.CbHandleHandshake = func(context.Context, []byte) (bool, error) { return true, nil }
	}
	if s.CbSizeHandshake == nil {
		s.CbSizeHandshake = func() int { return 0 }
	}
	if s.CbEmitHandshake == nil {
		s.CbEmitHandshake = func([]byte) error { return nil }
	}
	return s
}

func (s testHandshaker) HandleHandshake(ctx context.Context, rh []byte) (bool, error) {
	return s.CbHandleHandshake(ctx, rh)
}

func (s testHandshaker) SizeHandshake() int {
	return s.CbSizeHandshake()
}

func (s testHandshaker) EmitHandshake(buf []byte) error {
	return s.CbEmitHandshake(buf)
}

func testSetup(
	t *testing.T,
	cb1 func(ctx context.Context, conn net.Conn),
	cb2 func(ctx context.Context, conn net.Conn),
) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	defer l.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", l.Addr().String())
		if !assert.NoError(t, err) {
			t.FailNow()
		}
		cb1(ctx, conn)
	}()
	defer wg.Wait()
	conn, err := l.Accept()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	cb2(ctx, conn)
}

func testSetup2(
	t *testing.T,
	opts1 *Options,
	opts2 *Options,
	cb1 func(ctx context.Context, tp *Transport),
	cb2 func(ctx context.Context, tp *Transport),
) {
	testSetup(
		t,
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(opts1, false, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{
				CbSizeHandshake: func() int {
					return 0
				},
				CbEmitHandshake: func(buf []byte) error {
					return nil
				},
				CbHandleHandshake: func(ctx context.Context, rh []byte) (bool, error) {
					return true, nil
				},
			}.Init())
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			if !assert.True(t, ok) {
				t.FailNow()
			}
			tp.Prepare(DummyTrafficDecrypter{})
			cb1(ctx, tp)
		},
		func(ctx context.Context, conn net.Conn) {
			tp := new(Transport).Init(opts2, true, uuid.UUID{})
			defer tp.Close()
			ok, err := tp.Establish(ctx, conn, testHandshaker{
				CbHandleHandshake: func(ctx context.Context, rh []byte) (bool, error) {
					return true, nil
				},
				CbSizeHandshake: func() int {
					return 0
				},
				CbEmitHandshake: func(buf []byte) error {
					return nil
				},
			}.Init())
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			if !assert.True(t, ok) {
				t.FailNow()
			}
			tp.Prepare(DummyTrafficDecrypter{})
			cb2(ctx, tp)
		},
	)
}

var logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
