package methods

import (
	"github.com/bigelle/tele.go/types"
)

type SetPassportDataErrors struct {
	UserId int
	Errors []types.PassportElementError
}

func (s SetPassportDataErrors) Validate() error {
	if s.UserId < 1 {
		return types.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(s.Errors) < 1 {
		return types.ErrInvalidParam("types.parameter can't be empty")
	}
	return nil
}
