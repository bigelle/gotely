package internal

import "encoding/json"

type Validable interface {
	Validate() error
}

type Executable interface {
	// TODO: create an interface for returned by API types
	Validable
	json.Marshaler
}
