package channel

import (
	"fmt"

	"github.com/let-z-go/gogorpc/internal/proto"
)

const (
	RPCErrorBadRequest      = proto.RPC_ERROR_BAD_REQUEST
	RPCErrorUnauthorized    = proto.RPC_ERROR_UNAUTHORIZED
	RPCErrorForbidden       = proto.RPC_ERROR_FORBIDDEN
	RPCErrorNotFound        = proto.RPC_ERROR_NOT_FOUND
	RPCErrorTooManyRequests = proto.RPC_ERROR_TOO_MANY_REQUESTS

	RPCErrorInternalServer     = proto.RPC_ERROR_INTERNAL_SERVER
	RPCErrorNotImplemented     = proto.RPC_ERROR_NOT_IMPLEMENTED
	RPCErrorBadGateway         = proto.RPC_ERROR_BAD_GATEWAY
	RPCErrorServiceUnavailable = proto.RPC_ERROR_SERVICE_UNAVAILABLE
	RPCErrorGatewayTimeout     = proto.RPC_ERROR_GATEWAY_TIMEOUT
)

type RPCError proto.RPCError

func (re *RPCError) Error() string {
	return fmt.Sprintf("gogorpc/channel: rpc error: %d - %s", re.Type, re.Code)
}

func (re *RPCError) Equals(err error) bool {
	if other, ok := err.(*RPCError); ok {
		return re.Code == other.Code
	}

	return false
}

func (re *RPCError) Describe(desc string) *RPCError {
	return &RPCError{
		Type: re.Type,
		Code: re.Code,
		Desc: desc,
	}
}

type RPCErrorType = proto.RPCErrorType

func NewRPCError(rpcErrorType RPCErrorType, rpcErrorName string) *RPCError {
	return &RPCError{rpcErrorType, rpcErrorName, ""}
}

var (
	RPCErrBadRequest      = NewRPCError(RPCErrorBadRequest, "BadRequest")
	RPCErrUnauthorized    = NewRPCError(RPCErrorUnauthorized, "Unauthorized")
	RPCErrForbidden       = NewRPCError(RPCErrorForbidden, "Forbidden")
	RPCErrNotFound        = NewRPCError(RPCErrorNotFound, "NotFound")
	RPCErrTooManyRequests = NewRPCError(RPCErrorTooManyRequests, "TooManyRequests")

	RPCErrInternalServer     = NewRPCError(RPCErrorInternalServer, "InternalServer")
	RPCErrNotImplemented     = NewRPCError(RPCErrorNotImplemented, "NotImplemented")
	RPCErrBadGateway         = NewRPCError(RPCErrorBadGateway, "BadGateway")
	RPCErrServiceUnavailable = NewRPCError(RPCErrorServiceUnavailable, "ServiceUnavailable")
	RPCErrGatewayTimeout     = NewRPCError(RPCErrorGatewayTimeout, "GatewayTimeout")
)
