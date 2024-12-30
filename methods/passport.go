package methods

import (
	"github.com/bigelle/tele.go/objects"
)

type SetPassportDataErrors struct {
	UserId int
	Errors []objects.PassportElementError
}

func (s SetPassportDataErrors) Validate() error {
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(s.Errors) < 1 {
		return objects.ErrInvalidParam("objects.parameter can't be empty")
	}
	return nil
}
