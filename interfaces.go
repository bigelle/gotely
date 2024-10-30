package telego

type Validator interface {
	Validate() error
}
