package generator

type constantGenerator struct {
	name     string
	template string
}

var _ Generator = &constantGenerator{}

func NewConstantGenerator(template string, outputName string) *constantGenerator {
	return &constantGenerator{template: template, name: outputName}
}

func (c *constantGenerator) Name() string {
	return c.name
}

func (c *constantGenerator) Generate() (string, error) {
	return c.template, nil
}
