package generator

import (
	"context"

	"pages/internal/interactions/github"
)

type projectsGenerator struct {
	githubClient github.Client
}

var _ Generator = &projectsGenerator{}

func NewProjectsGenerator(githubClient github.Client) *projectsGenerator {
	return &projectsGenerator{githubClient: githubClient}
}

func (*projectsGenerator) Generate(ctx context.Context) (string, error) {
	return "", nil
}

func (*projectsGenerator) Name() string {
	return "projects.md"
}
