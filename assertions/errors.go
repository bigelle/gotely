package assertions

import "fmt"

type ErrFailedRequest struct {
	code    int
	message string
}

func (e ErrFailedRequest) Error() string {
	return fmt.Sprintf("error %d: %s", e.code, e.message)
}
