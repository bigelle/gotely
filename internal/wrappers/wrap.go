package wrappers

import (
	"encoding/json"
)

func MarshalWithConstTypeField[T any](v T, Type string) ([]byte, error) {
	return json.Marshal(&struct {
		Type string
		Data T
	}{
		Type: Type,
		Data: v,
	})
}
