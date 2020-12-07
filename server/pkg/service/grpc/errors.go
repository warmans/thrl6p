package grpc

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func ErrInvalidRequestField(field string, reason string) *status.Status {
	s, err := status.New(codes.InvalidArgument, http.StatusText(http.StatusBadRequest)).WithDetails(
		&errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{Field: field, Description: reason},
			},
		},
	)
	if err != nil {
		return status.New(codes.Internal, "failed to create error")
	}
	return s
}

func ErrInternal(err error) *status.Status {
	if err == nil {
		return status.New(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}
	// do not wrap existing grpc errors
	if sta, ok := status.FromError(err); ok {
		return sta
	}
	s, err := status.New(codes.Internal, http.StatusText(http.StatusInternalServerError)).WithDetails(
		&errdetails.DebugInfo{
			Detail: err.Error(),
		},
	)
	if err != nil {
		return status.New(codes.Internal, "failed to create error")
	}
	return s
}
