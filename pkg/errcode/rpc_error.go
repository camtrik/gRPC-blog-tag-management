package errcode

import (
	pb "github.com/camtrik/gRPC-blog-tag-management/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Status struct {
	*status.Status
}

func TogRPCError(err *Error) error {
	s, _ := status.New(ToRPCCode(err.Code()), err.Msg()).WithDetails(&pb.Error{Code: int32(err.Code()), Message: err.Msg()})
	return s.Err()
}

func FromError(err error) *Status {
	s, _ := status.FromError(err)
	return &Status{s}
}

func ToRPCCode(code int) codes.Code {
	var statusCode codes.Code
	switch code {
	case Fail.Code():
		statusCode = codes.Internal
	case InvalidParams.Code():
		statusCode = codes.InvalidArgument
	case Unautorized.Code():
		statusCode = codes.Unauthenticated
	case NotFound.Code():
		statusCode = codes.NotFound
	case Unknown.Code():
		statusCode = codes.Unknown
	case DeadlineExceeded.Code():
		statusCode = codes.DeadlineExceeded
	case AccessDenied.Code():
		statusCode = codes.PermissionDenied
	case LimitExceed.Code():
		statusCode = codes.ResourceExhausted
	case MethodNotAllow.Code():
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}

	return statusCode
}
