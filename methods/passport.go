package methods

import (
	"fmt"
	"io"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/objects"
)

// /Informs a user that some of the Telegram Passport elements they provided contains errors.
// The user will not be able to re-submit their Passport to you until the errors are fixed
// (the contents of the field for which you returned the error must change).
// Returns True on success.
//
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason.
// For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc.
// Supply some details in the error message to make sure the user knows how to correct the issues.
type SetPassportDataErrors struct {
	// REQUIRED:
	// User identifier
	UserId int `json:"user_id"`
	// REQUIRED:
	// A JSON-serialized array describing the errors
	Errors []objects.PassportElementError `json:"errors"`
}

func (s SetPassportDataErrors) Validate() error {
	var err gotely.ErrFailedValidation
	if s.UserId < 1 {
		err = append(err, fmt.Errorf("user_id parameter can't be empty"))
	}
	if len(s.Errors) < 1 {
		err = append(err, fmt.Errorf("objects.parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SetPassportDataErrors) Endpoint() string {
	return "setPassportDataErrors"
}

func (s SetPassportDataErrors) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetPassportDataErrors) ContentType() string {
	return "application/json"
}
