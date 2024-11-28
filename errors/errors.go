package errors

import "fmt"

type ErrFailedRequest struct {
	Code    int
	Message string
}

func (e ErrFailedRequest) Error() string {
	return fmt.Sprintf("request failed with code %d: %s", e.Code, e.Message)
}

type ErrInvalidParam string

func (e ErrInvalidParam) Error() string {
	return string(e)
}
