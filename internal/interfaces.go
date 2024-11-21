package internal

type Validable interface {
	Validate() error
}
