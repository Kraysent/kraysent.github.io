package github

import "context"

type Client interface {
	ListRepositories(ctx context.Context, user string) ([]Repository, error)
}
