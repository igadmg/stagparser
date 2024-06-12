package stagparser

type Definition struct {
	Name       string
	Attributes map[string]any
}

func newDefinition(name string, attributes map[string]any) Definition {
	return Definition{
		Name:       name,
		Attributes: attributes,
	}
}
