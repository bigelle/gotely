package api

import (
	"encoding/json"
	"io"
	"mime/multipart"
)

func WriteJSONToForm(mw *multipart.Writer, key string, value any) error {
	r, err := mw.CreateFormField(key)
	if err != nil {
		return err
	}
	return json.NewEncoder(r).Encode(value)
}

func EncodeJSON(body any) (io.Reader, error) {
	pr, pw := io.Pipe()
	enc := json.NewEncoder(pw)
	go func() {
		defer pw.Close()
		if err := enc.Encode(body); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()
	return pr, nil
}
