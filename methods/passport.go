package methods

import (
	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/types"
)

type SetPassportDataErrors struct {
	UserId int
	Errors []types.PassportElementError
}

func (s SetPassportDataErrors) Validate() error {
	if s.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(s.Errors) < 1 {
		return errors.ErrInvalidParam("errors parameter can't be empty")
	}
	return nil
}
