package generator

import (
	"strings"
	"text/template"
)

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
	t, err := template.New("example").Parse(c.template)
	if err != nil {
		return "", err
	}

	var builder strings.Builder

	if err := t.Execute(&builder, nil); err != nil {
		return "", err
	}

	return builder.String(), err
}
