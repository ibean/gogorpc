// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protocol/protocol.proto

/*
	Package protocol is a generated protocol buffer package.

	It is generated from these files:
		protocol/protocol.proto

	It has these top-level messages:
		Greeting
		RequestHeader
		ResponseHeader
		Heartbeat
*/
package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MessageType int32

const (
	MESSAGE_REQUEST   MessageType = 0
	MESSAGE_RESPONSE  MessageType = 1
	MESSAGE_HEARTBEAT MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "MESSAGE_REQUEST",
	1: "MESSAGE_RESPONSE",
	2: "MESSAGE_HEARTBEAT",
}
var MessageType_value = map[string]int32{
	"MESSAGE_REQUEST":   0,
	"MESSAGE_RESPONSE":  1,
	"MESSAGE_HEARTBEAT": 2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

type ErrorCode int32

const (
	ERROR_NO              ErrorCode = 0
	ERROR_CHANNEL_BUSY    ErrorCode = 1
	ERROR_NOT_FOUND       ErrorCode = 2
	ERROR_BAD_REQUEST     ErrorCode = 3
	ERROR_NOT_IMPLEMENTED ErrorCode = 4
	ERROR_INTERNAL_SERVER ErrorCode = 5
	ERROR_USER_DEFINED    ErrorCode = 256
)

var ErrorCode_name = map[int32]string{
	0:   "ERROR_NO",
	1:   "ERROR_CHANNEL_BUSY",
	2:   "ERROR_NOT_FOUND",
	3:   "ERROR_BAD_REQUEST",
	4:   "ERROR_NOT_IMPLEMENTED",
	5:   "ERROR_INTERNAL_SERVER",
	256: "ERROR_USER_DEFINED",
}
var ErrorCode_value = map[string]int32{
	"ERROR_NO":              0,
	"ERROR_CHANNEL_BUSY":    1,
	"ERROR_NOT_FOUND":       2,
	"ERROR_BAD_REQUEST":     3,
	"ERROR_NOT_IMPLEMENTED": 4,
	"ERROR_INTERNAL_SERVER": 5,
	"ERROR_USER_DEFINED":    256,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

type Greeting struct {
	Channel   Greeting_Channel `protobuf:"bytes,1,opt,name=channel" json:"channel"`
	Handshake []byte           `protobuf:"bytes,2,opt,name=handshake,proto3" json:"handshake,omitempty"`
}

func (m *Greeting) Reset()                    { *m = Greeting{} }
func (m *Greeting) String() string            { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()               {}
func (*Greeting) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0} }

func (m *Greeting) GetChannel() Greeting_Channel {
	if m != nil {
		return m.Channel
	}
	return Greeting_Channel{}
}

func (m *Greeting) GetHandshake() []byte {
	if m != nil {
		return m.Handshake
	}
	return nil
}

type Greeting_Channel struct {
	Id                 []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timeout            int32  `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	IncomingWindowSize int32  `protobuf:"varint,3,opt,name=incoming_window_size,json=incomingWindowSize,proto3" json:"incoming_window_size,omitempty"`
	OutgoingWindowSize int32  `protobuf:"varint,4,opt,name=outgoing_window_size,json=outgoingWindowSize,proto3" json:"outgoing_window_size,omitempty"`
}

func (m *Greeting_Channel) Reset()                    { *m = Greeting_Channel{} }
func (m *Greeting_Channel) String() string            { return proto.CompactTextString(m) }
func (*Greeting_Channel) ProtoMessage()               {}
func (*Greeting_Channel) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{0, 0} }

func (m *Greeting_Channel) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Greeting_Channel) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Greeting_Channel) GetIncomingWindowSize() int32 {
	if m != nil {
		return m.IncomingWindowSize
	}
	return 0
}

func (m *Greeting_Channel) GetOutgoingWindowSize() int32 {
	if m != nil {
		return m.OutgoingWindowSize
	}
	return 0
}

type RequestHeader struct {
	TraceId        []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SequenceNumber int32  `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	ServiceName    string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	MethodName     string `protobuf:"bytes,4,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
}

func (m *RequestHeader) Reset()                    { *m = RequestHeader{} }
func (m *RequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()               {}
func (*RequestHeader) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{1} }

func (m *RequestHeader) GetTraceId() []byte {
	if m != nil {
		return m.TraceId
	}
	return nil
}

func (m *RequestHeader) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *RequestHeader) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RequestHeader) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

type ResponseHeader struct {
	SequenceNumber int32 `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	ErrorCode      int32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (m *ResponseHeader) Reset()                    { *m = ResponseHeader{} }
func (m *ResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*ResponseHeader) ProtoMessage()               {}
func (*ResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{2} }

func (m *ResponseHeader) GetSequenceNumber() int32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *ResponseHeader) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

type Heartbeat struct {
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptorProtocol, []int{3} }

func init() {
	proto.RegisterType((*Greeting)(nil), "Greeting")
	proto.RegisterType((*Greeting_Channel)(nil), "Greeting.Channel")
	proto.RegisterType((*RequestHeader)(nil), "RequestHeader")
	proto.RegisterType((*ResponseHeader)(nil), "ResponseHeader")
	proto.RegisterType((*Heartbeat)(nil), "Heartbeat")
	proto.RegisterEnum("MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("ErrorCode", ErrorCode_name, ErrorCode_value)
}
func (m *Greeting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Greeting) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintProtocol(dAtA, i, uint64(m.Channel.Size()))
	n1, err := m.Channel.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Handshake) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Handshake)))
		i += copy(dAtA[i:], m.Handshake)
	}
	return i, nil
}

func (m *Greeting_Channel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Greeting_Channel) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Timeout != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.Timeout))
	}
	if m.IncomingWindowSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.IncomingWindowSize))
	}
	if m.OutgoingWindowSize != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.OutgoingWindowSize))
	}
	return i, nil
}

func (m *RequestHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TraceId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.TraceId)))
		i += copy(dAtA[i:], m.TraceId)
	}
	if m.SequenceNumber != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SequenceNumber))
	}
	if len(m.ServiceName) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.ServiceName)))
		i += copy(dAtA[i:], m.ServiceName)
	}
	if len(m.MethodName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(len(m.MethodName)))
		i += copy(dAtA[i:], m.MethodName)
	}
	return i, nil
}

func (m *ResponseHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseHeader) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.SequenceNumber != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.SequenceNumber))
	}
	if m.ErrorCode != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtocol(dAtA, i, uint64(m.ErrorCode))
	}
	return i, nil
}

func (m *Heartbeat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Heartbeat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintProtocol(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Greeting) Size() (n int) {
	var l int
	_ = l
	l = m.Channel.Size()
	n += 1 + l + sovProtocol(uint64(l))
	l = len(m.Handshake)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	return n
}

func (m *Greeting_Channel) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.Timeout != 0 {
		n += 1 + sovProtocol(uint64(m.Timeout))
	}
	if m.IncomingWindowSize != 0 {
		n += 1 + sovProtocol(uint64(m.IncomingWindowSize))
	}
	if m.OutgoingWindowSize != 0 {
		n += 1 + sovProtocol(uint64(m.OutgoingWindowSize))
	}
	return n
}

func (m *RequestHeader) Size() (n int) {
	var l int
	_ = l
	l = len(m.TraceId)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	if m.SequenceNumber != 0 {
		n += 1 + sovProtocol(uint64(m.SequenceNumber))
	}
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	l = len(m.MethodName)
	if l > 0 {
		n += 1 + l + sovProtocol(uint64(l))
	}
	return n
}

func (m *ResponseHeader) Size() (n int) {
	var l int
	_ = l
	if m.SequenceNumber != 0 {
		n += 1 + sovProtocol(uint64(m.SequenceNumber))
	}
	if m.ErrorCode != 0 {
		n += 1 + sovProtocol(uint64(m.ErrorCode))
	}
	return n
}

func (m *Heartbeat) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovProtocol(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtocol(x uint64) (n int) {
	return sovProtocol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Greeting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Greeting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Greeting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Channel.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handshake", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handshake = append(m.Handshake[:0], dAtA[iNdEx:postIndex]...)
			if m.Handshake == nil {
				m.Handshake = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Greeting_Channel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Channel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Channel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			m.Timeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timeout |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncomingWindowSize", wireType)
			}
			m.IncomingWindowSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IncomingWindowSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutgoingWindowSize", wireType)
			}
			m.OutgoingWindowSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutgoingWindowSize |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RequestHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RequestHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceId = append(m.TraceId[:0], dAtA[iNdEx:postIndex]...)
			if m.TraceId == nil {
				m.TraceId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MethodName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtocol
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MethodName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResponseHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResponseHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorCode", wireType)
			}
			m.ErrorCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrorCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Heartbeat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Heartbeat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Heartbeat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipProtocol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtocol
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProtocol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtocol
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProtocol
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtocol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtocol
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipProtocol(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthProtocol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtocol   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protocol/protocol.proto", fileDescriptorProtocol) }

var fileDescriptorProtocol = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcd, 0x4e, 0xdb, 0x4e,
	0x14, 0xc5, 0x33, 0x21, 0xfc, 0x43, 0x6e, 0xf2, 0x07, 0x33, 0x85, 0x12, 0x50, 0x1b, 0x68, 0x36,
	0x45, 0x48, 0x84, 0x7e, 0x6c, 0xba, 0x4d, 0xc8, 0x40, 0x22, 0x11, 0x87, 0x8e, 0x9d, 0x7e, 0x6d,
	0x2c, 0xc7, 0xbe, 0x75, 0xac, 0x62, 0x4f, 0x6a, 0x8f, 0x8b, 0xca, 0xaa, 0x8f, 0xd0, 0x1d, 0x2f,
	0xd1, 0x07, 0x61, 0xc9, 0x13, 0x54, 0x2d, 0x7d, 0x86, 0xee, 0x2b, 0x8f, 0x63, 0x52, 0x55, 0xec,
	0xee, 0x3d, 0xe7, 0x77, 0xe6, 0x9e, 0xcd, 0xc0, 0xc6, 0x34, 0x12, 0x52, 0x38, 0xe2, 0xec, 0x20,
	0x1f, 0x5a, 0x6a, 0xd8, 0xda, 0xf7, 0x7c, 0x39, 0x49, 0xc6, 0x2d, 0x47, 0x04, 0x07, 0x9e, 0xf0,
	0x44, 0xe6, 0x8f, 0x93, 0xf7, 0x6a, 0x53, 0x8b, 0x9a, 0x32, 0xbc, 0xf9, 0x9b, 0xc0, 0xd2, 0x71,
	0x84, 0x28, 0xfd, 0xd0, 0xa3, 0x4f, 0xa1, 0xec, 0x4c, 0xec, 0x30, 0xc4, 0xb3, 0x3a, 0xd9, 0x21,
	0xbb, 0xd5, 0x67, 0xab, 0xad, 0xdc, 0x6b, 0x1d, 0x66, 0x46, 0xa7, 0x74, 0xf5, 0x7d, 0xbb, 0xc0,
	0x73, 0x8e, 0x3e, 0x80, 0xca, 0xc4, 0x0e, 0xdd, 0x78, 0x62, 0x7f, 0xc0, 0x7a, 0x71, 0x87, 0xec,
	0xd6, 0xf8, 0x5c, 0xd8, 0xba, 0x24, 0x50, 0x9e, 0x05, 0xe9, 0x32, 0x14, 0x7d, 0x57, 0xbd, 0x5b,
	0xe3, 0x45, 0xdf, 0xa5, 0x75, 0x28, 0x4b, 0x3f, 0x40, 0x91, 0x48, 0x95, 0x5b, 0xe4, 0xf9, 0x4a,
	0x9f, 0xc0, 0x9a, 0x1f, 0x3a, 0x22, 0xf0, 0x43, 0xcf, 0x3a, 0xf7, 0x43, 0x57, 0x9c, 0x5b, 0xb1,
	0x7f, 0x81, 0xf5, 0x05, 0x85, 0xd1, 0xdc, 0x7b, 0xad, 0x2c, 0xc3, 0xbf, 0xc0, 0x34, 0x21, 0x12,
	0xe9, 0x89, 0x7f, 0x13, 0xa5, 0x2c, 0x91, 0x7b, 0xf3, 0x44, 0xf3, 0x92, 0xc0, 0xff, 0x1c, 0x3f,
	0x26, 0x18, 0xcb, 0x1e, 0xda, 0x2e, 0x46, 0x74, 0x13, 0x96, 0x64, 0x64, 0x3b, 0x68, 0xdd, 0xb6,
	0x2c, 0xab, 0xbd, 0xef, 0xd2, 0xc7, 0xb0, 0x12, 0xa7, 0x6c, 0xe8, 0xa0, 0x15, 0x26, 0xc1, 0x18,
	0xa3, 0x59, 0xe5, 0xe5, 0x5c, 0xd6, 0x95, 0x4a, 0x1f, 0x41, 0x2d, 0xc6, 0xe8, 0x93, 0x9f, 0x72,
	0x76, 0x90, 0x35, 0xae, 0xf0, 0xea, 0x4c, 0xd3, 0xed, 0x00, 0xe9, 0x36, 0x54, 0x03, 0x94, 0x13,
	0xe1, 0x66, 0x44, 0x49, 0x11, 0x90, 0x49, 0x29, 0xd0, 0x7c, 0x03, 0xcb, 0x1c, 0xe3, 0xa9, 0x08,
	0x63, 0x9c, 0x35, 0xbb, 0xe3, 0x3c, 0xb9, 0xf3, 0xfc, 0x43, 0x00, 0x8c, 0x22, 0x11, 0x59, 0x8e,
	0x70, 0x71, 0x56, 0xb1, 0xa2, 0x94, 0x43, 0xe1, 0x62, 0xb3, 0x0a, 0x95, 0x1e, 0xda, 0x91, 0x1c,
	0xa3, 0x2d, 0xf7, 0x86, 0x50, 0x1d, 0x60, 0x1c, 0xdb, 0x1e, 0x9a, 0x9f, 0xa7, 0x48, 0xef, 0xc1,
	0xca, 0x80, 0x19, 0x46, 0xfb, 0x98, 0x59, 0x9c, 0xbd, 0x1c, 0x31, 0xc3, 0xd4, 0x0a, 0x74, 0x0d,
	0xb4, 0xb9, 0x68, 0x9c, 0x0e, 0x75, 0x83, 0x69, 0x84, 0xae, 0xc3, 0x6a, 0xae, 0xf6, 0x58, 0x9b,
	0x9b, 0x1d, 0xd6, 0x36, 0xb5, 0xe2, 0xde, 0x37, 0x02, 0x15, 0x96, 0xdf, 0xa2, 0x35, 0x58, 0x62,
	0x9c, 0x0f, 0xb9, 0xa5, 0x0f, 0xb5, 0x02, 0xbd, 0x0f, 0x34, 0xdb, 0x0e, 0x7b, 0x6d, 0x5d, 0x67,
	0x27, 0x56, 0x67, 0x64, 0xbc, 0xd5, 0x48, 0x7a, 0x35, 0xa7, 0x4c, 0xeb, 0x68, 0x38, 0xd2, 0xbb,
	0x5a, 0x31, 0x7d, 0x3f, 0x13, 0x3b, 0xed, 0xee, 0x6d, 0x99, 0x05, 0xba, 0x09, 0xeb, 0x73, 0xb6,
	0x3f, 0x38, 0x3d, 0x61, 0x03, 0xa6, 0x9b, 0xac, 0xab, 0x95, 0xe6, 0x56, 0x5f, 0x37, 0x19, 0xd7,
	0xdb, 0x27, 0x96, 0xc1, 0xf8, 0x2b, 0xc6, 0xb5, 0x45, 0xba, 0x91, 0x5f, 0x1e, 0x19, 0x8c, 0x5b,
	0x5d, 0x76, 0xd4, 0xd7, 0x59, 0x57, 0xfb, 0x52, 0xec, 0xbc, 0xb8, 0xfe, 0xd9, 0x28, 0x5c, 0xdd,
	0x34, 0xc8, 0xf5, 0x4d, 0x83, 0xfc, 0xb8, 0x69, 0x90, 0xaf, 0xbf, 0x1a, 0x85, 0x77, 0xcd, 0xbf,
	0x7e, 0xcf, 0x19, 0xca, 0xfd, 0x8b, 0xfd, 0xf4, 0x07, 0x8d, 0xa3, 0xa9, 0x73, 0xfb, 0xcf, 0xc6,
	0xff, 0xa9, 0xe9, 0xf9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x9c, 0xa1, 0x2d, 0x83, 0x03,
	0x00, 0x00,
}
