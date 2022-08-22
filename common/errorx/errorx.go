package errorx

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const defaultCode = 1001

var (
	ErrWrongPassword      = NewCodeError(1002, "wrong password")
	ErrUsernameNotFound   = NewCodeError(1003, "username not found")
	ErrDuplicateUsername  = NewCodeError(1004, "username had been registered")
	ErrDuplicateSubscribe = NewCodeError(1005, "item had been subscribed by the same user")
	ErrWrongUrlFormat     = NewCodeError(1006, "wrong url format")
)

func NewDefaultError(msg string) error {
	return status.Error(defaultCode, "error: "+msg)
}

func NewCodeError(code codes.Code, msg string) error {
	return status.Error(code, "error: "+msg)
}
