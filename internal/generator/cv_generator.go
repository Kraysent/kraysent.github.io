package generator

import (
	"strings"
	"text/template"

	"pages/internal/templates"
)

type cvGenerator struct {
	cvPath string
}

var _ Generator = &cvGenerator{}

func NewCVGenerator() *cvGenerator {
	return &cvGenerator{}
}

func (g *cvGenerator) Generate() (string, error) {
	t, err := template.New("example").Parse(templates.CVTemplate)
	if err != nil {
		return "", err
	}

	var builder strings.Builder

	if err := t.Execute(&builder, nil); err != nil {
		return "", err
	}

	return builder.String(), err
}

func (g *cvGenerator) Name() string {
	return "cv.md"
}
