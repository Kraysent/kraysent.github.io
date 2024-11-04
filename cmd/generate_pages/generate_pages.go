package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"pages/internal/generator"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	outputDirectory := flag.String("output", "gen", "output directory")

	flag.Parse()

	os.Mkdir(*outputDirectory, os.ModePerm)

	err := generate(*outputDirectory)
	if err != nil {
		panic(err)
	}
}

func generate(outputDir string) error {
	ctx := context.Background()

	githubToken, ok := os.LookupEnv("GITHUB_TOKEN")
	if !ok {
		return fmt.Errorf("GITHUB_TOKEN environment variable is not set")
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	generators := []generator.Generator{
		generator.NewCurrentTimeGenerator(),
		generator.NewGithubProjectsGenerator(client),
	}

	for _, g := range generators {
		s, err := g.Generate(ctx)
		if err != nil {
			return err
		}

		filePath := fmt.Sprintf("%s/%s", outputDir, g.Filename())
		fmt.Printf("Writing %s\n", filePath)
		err = writeFile(filePath, s)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeFile(filePath, content string) error {
	if err := os.MkdirAll(filepath.Dir(filePath), 0o770); err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
