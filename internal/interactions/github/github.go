package github

import (
	"context"

	"pages/internal/common"

	"github.com/google/go-github/github"
)

type githubClient struct {
	client *github.Client
}

var _ Client = &githubClient{}

func NewClient(client *github.Client) *githubClient {
	return &githubClient{client: client}
}

func (c *githubClient) ListRepositories(ctx context.Context, user string) ([]Repository, error) {
	// TODO pagination
	repos, _, err := c.client.Repositories.List(ctx, user, &github.RepositoryListOptions{})
	if err != nil {
		return nil, err
	}

	var result []Repository

	for _, repo := range repos {
		result = append(result, Repository{
			Name:          common.PtrFrom(repo.Name),
			Author:        user,
			Description:   common.PtrFrom(repo.Description),
			StarNumber:    common.PtrFrom(repo.StargazersCount),
			WatcherNumber: common.PtrFrom(repo.WatchersCount),
			ForksNumber:   common.PtrFrom(repo.ForksCount),
			Website:       common.PtrFrom(repo.Homepage),
			Tags:          repo.Topics,
		})
	}

	return result, nil
}
