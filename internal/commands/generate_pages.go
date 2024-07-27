package commands

import (
	"context"
	"flag"
	"fmt"
	"os"

	"pages/internal/common"
	"pages/internal/generator"
	"pages/internal/templates"

	"github.com/google/go-github/github"
)

type GeneratePagesCommand struct {
	githubClient *github.Client
	generators   []generator.Generator

	outputDir string
}

func (c *GeneratePagesCommand) Init() {
	outputDir := flag.String("output", "gen", "output directory")

	flag.Parse()

	c.outputDir = common.PtrFrom(outputDir)

	c.githubClient = github.NewClient(nil)
	c.generators = []generator.Generator{
		generator.NewConstantGenerator(templates.IndexTemplate, "index.md"),
		generator.NewConstantGenerator(templates.ConfigTemplate, "_config.yml"),
		generator.NewConstantGenerator(templates.GemfileTemplate, "Gemfile"),
	}
}

func (c *GeneratePagesCommand) Run(ctx context.Context) error {
	os.Mkdir(c.outputDir, os.ModePerm)

	for _, g := range c.generators {
		s, err := g.Generate()
		if err != nil {
			return err
		}

		filePath := fmt.Sprintf("%s/%s", c.outputDir, g.Name())

		fmt.Printf("Writing %s\n", filePath)

		if common.WriteFile(filePath, s); err != nil {
			return err
		}
	}

	return nil
}