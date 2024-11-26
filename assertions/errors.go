package assertions

import "fmt"

type ErrFailedRequest struct {
	Code    int
	Message string
}

func (e ErrFailedRequest) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}
