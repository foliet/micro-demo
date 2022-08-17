package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultCode = 1001

type ErrorResponse struct {
	Code codes.Code
	Msg  string
}

func NewDefaultError(msg string) error {
	return status.Error(defaultCode, msg)
}
