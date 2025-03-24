package gotely

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

// TODO: all of the possible error types
// TODO: collection of errors (used in validation)
// TODO: maybe here also should be interfaces, not sure

type ErrInvalidParam string

func (e ErrInvalidParam) Error() string {
	return string(e)
}

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

func DecodeJSON(source io.Reader, dest any) error {
	return json.NewDecoder(source).Decode(dest)
}

func WriteJSONToForm(mw *multipart.Writer, key string, value any) error {
	r, err := mw.CreateFormField(key)
	if err != nil {
		return err
	}
	return json.NewEncoder(r).Encode(value)
}

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
