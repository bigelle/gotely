package gotely

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
)

// TODO: implement it everywhere
// ErrFailedValidation contains a list of errors
// that occurred during validation.
type ErrFailedValidation []error

func (e ErrFailedValidation) Error() string {
	var sb strings.Builder
	for i, err := range e {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(err.Error())
	}
	return sb.String()
}

func (e ErrFailedValidation) Is(target error) bool {
	_, ok := target.(ErrFailedValidation)
	return ok
}

func (e ErrFailedValidation) Unwrap() []error {
	return e
}

// ErrTelegramAPIFailedRequest describes the reason
// why the request to the Telegram Bot API was unsuccessful.
type ErrTelegramAPIFailedRequest struct {
	Code               int
	Description        string
	ResponseParameters *ResponseParameters
}

func (e ErrTelegramAPIFailedRequest) Error() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("error %d: %s", e.Code, e.Description))
	if e.ResponseParameters != nil {
		if e.ResponseParameters.MigrateToChatId != nil {
			sb.WriteString(fmt.Sprintf(", the group has been migrated to supergroup with id=%d", *e.ResponseParameters.MigrateToChatId))
		}
		if e.ResponseParameters.RetryAfter != nil {
			sb.WriteString(fmt.Sprintf(", retry after %d seconds", *e.ResponseParameters.RetryAfter))
		}
	}
	return sb.String()
}

func (e ErrTelegramAPIFailedRequest) Is(target error) bool {
	_, ok := target.(ErrTelegramAPIFailedRequest)
	return ok
}

func (e ErrTelegramAPIFailedRequest) Unwrap() error {
	return e
}

// DecodeExactField reads the contents of source, searches for the specified field
// and writes it's value to dest
func DecodeExactField(source io.Reader, field string, dest any) error {
	dec := json.NewDecoder(source)
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}

		if key, ok := tok.(string); ok && key == field {
			return dec.Decode(dest)
		}
	}
	return fmt.Errorf("unknown field: %s", field)
}

// DecodeJSON reads JSON from source and decodes it into dest
func DecodeJSON(source io.Reader, dest any) error {
	return json.NewDecoder(source).Decode(dest)
}

// WriteJSONToForm creates multipart form field with the given key
// and writes the JSON-encoded value into it
func WriteJSONToForm(mw *multipart.Writer, key string, value any) error {
	r, err := mw.CreateFormField(key)
	if err != nil {
		return err
	}
	return json.NewEncoder(r).Encode(value)
}

// EncodeJSON encodes the given value into JSON and returns as [io.Reader]
func EncodeJSON(body any) io.Reader {
	pr, pw := io.Pipe()
	enc := json.NewEncoder(pw)
	go func() {
		defer pw.Close()
		if err := enc.Encode(body); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()
	return pr
}
