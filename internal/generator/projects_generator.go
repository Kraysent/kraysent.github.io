package generator

import (
	"context"
	"strings"
	"text/template"

	"pages/internal/interactions/github"
	"pages/internal/templates"
)

type projectsGenerator struct {
	githubClient github.Client
}

var _ Generator = &projectsGenerator{}

func NewProjectsGenerator(githubClient github.Client) *projectsGenerator {
	return &projectsGenerator{githubClient: githubClient}
}

type placeholderProject struct {
	Name   string
	Author string
}

type placeholder struct {
	Projects []placeholderProject
}

func (g *projectsGenerator) Generate(ctx context.Context) (string, error) {
	repos, err := g.githubClient.ListRepositories(ctx, "Kraysent")
	if err != nil {
		return "", err
	}

	fileTemplate, err := template.New("projects").Parse(templates.ProjectsTemplate)
	if err != nil {
		return "", err
	}

	var data placeholder

	for _, repo := range repos {
		data.Projects = append(data.Projects, placeholderProject{
			Name:   repo.Name,
			Author: repo.Author,
		})
	}

	var builder strings.Builder

	if err := fileTemplate.Execute(&builder, data); err != nil {
		return "", nil
	}

	return builder.String(), nil
}

func (g *projectsGenerator) Name() string {
	return "projects.md"
}
