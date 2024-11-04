package generator

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

type githubProjectsGenerator struct {
	client *github.Client
}

var _ Generator = &githubProjectsGenerator{}

func NewGithubProjectsGenerator(client *github.Client) *githubProjectsGenerator {
	return &githubProjectsGenerator{
		client: client,
	}
}

func (g *githubProjectsGenerator) Filename() string {
	return "github_projects.md"
}

func buildRepoMarkdown(repo *github.Repository) string {
	builder := strings.Builder{}

	name := fmt.Sprintf("%s/%s", repo.Owner.GetLogin(), repo.GetName())
	url := *repo.HTMLURL
	description := repo.GetDescription()
	stars := repo.GetStargazersCount()
	forks := repo.GetForksCount()
	language := repo.GetLanguage()
	isForked := repo.GetFork()
	license := repo.GetLicense().GetName()
	homepage := repo.GetHomepage()
	updatedAt := repo.GetUpdatedAt().Format("2006-01-02")
	openIssues := repo.GetOpenIssuesCount()
	topics := repo.Topics

	builder.WriteString(fmt.Sprintf("## [%s](%s)\n", name, url))
	if isForked {
		builder.WriteString("**Forked Repository**\n\n")
	}
	if description != "" {
		description = strings.Trim(description, "*")
		description = strings.Trim(description, "-")
		builder.WriteString(fmt.Sprintf("*%s*\n\n", description))
	}
	if stars != 0 {
		builder.WriteString(fmt.Sprintf("- Starred by %d people\n\n", stars))
	}
	if forks != 0 {
		builder.WriteString(fmt.Sprintf("- Forked by %d people\n\n", forks))
	}
	if language != "" {
		builder.WriteString(fmt.Sprintf("- Written in %s\n\n", language))
	}
	if license != "" {
		builder.WriteString(fmt.Sprintf("- Licensed with %s\n\n", license))
	}
	if homepage != "" {
		builder.WriteString(fmt.Sprintf("- [**%s**](%s)\n\n", homepage, homepage))
	}
	if updatedAt != "" {
		builder.WriteString(fmt.Sprintf("- Last updated at %s\n\n", updatedAt))
	}
	if openIssues != 0 {
		builder.WriteString(fmt.Sprintf("- %d open issues\n\n", openIssues))
	}
	if len(topics) > 0 {
		builder.WriteString("- **Topics:** ")
		for i, topic := range topics {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(fmt.Sprintf("`%s`", topic))
		}
		builder.WriteString("\n\n")
	}

	builder.WriteString("---\n\n")

	return builder.String()
}

func (g *githubProjectsGenerator) Generate(ctx context.Context) (string, error) {
	repos, _, err := g.client.Repositories.List(ctx, "", &github.RepositoryListOptions{
		Visibility: "public",
		Sort:       "pushed",
		Direction:  "desc",
	})
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}

	builder.WriteString("# Github Projects\n\n")

	for _, repo := range repos {
		builder.WriteString(buildRepoMarkdown(repo))
	}

	return builder.String(), nil
}
