package internal

type Validable interface {
	Validate() error
}

type Executable interface {
	// TODO: create an interface for returned by API types
	Validable
	ToRequestBody() ([]byte, error)
}
