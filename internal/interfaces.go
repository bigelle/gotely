package internal

type Validable interface {
	Validate() error
}

type Executable interface {
	Validable
	ToRequestBody() ([]byte, error)
}
