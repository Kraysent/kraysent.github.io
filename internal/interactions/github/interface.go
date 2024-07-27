package github

import "context"

type GithubClient interface {
	ListRepositories(ctx context.Context, user string) ([]Repository, error)
}
