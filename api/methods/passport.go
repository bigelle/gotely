package methods

import (
	"github.com/bigelle/gotely/api/objects"
)

type SetPassportDataErrors struct {
	UserId int                            `json:"user_id"`
	Errors []objects.PassportElementError `json:"errors"`
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
