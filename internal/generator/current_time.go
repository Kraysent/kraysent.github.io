package generator

import (
	"context"
	"strings"
	"time"
)

type currentTimeGenerator struct{}

var _ Generator = &currentTimeGenerator{}

func NewCurrentTimeGenerator() *currentTimeGenerator {
	return &currentTimeGenerator{}
}

func (c *currentTimeGenerator) Filename() string {
	return "current_time.md"
}

func (c *currentTimeGenerator) Generate(ctx context.Context) (string, error) {
	t := time.Now().Format(time.RFC850)

	builder := strings.Builder{}

	builder.WriteString("### This doc was built on ")
	builder.WriteString(t)

	return builder.String(), nil
}
