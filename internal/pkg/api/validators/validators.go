package validators

type Validator interface {
	Validate(interface{}) error
}
