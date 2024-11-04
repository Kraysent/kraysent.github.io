package generator

import "context"

type Generator interface {
	Generate(ctx context.Context) (string, error)
	Name() string
}
