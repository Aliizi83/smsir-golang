package verify

type Parameter struct {
	Name  string
	Value string
}

func NewParam(name, value string) *Parameter {
	return &Parameter{
		name,
		value,
	}
}
