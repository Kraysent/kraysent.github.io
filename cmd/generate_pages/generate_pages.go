package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"pages/internal/generator"
	"pages/internal/templates"
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
	generators := []generator.Generator{
		generator.NewConstantGenerator(templates.IndexTemplate, "index.md"),
		generator.NewConstantGenerator(templates.ConfigTemplate, "_config.yml"),
		generator.NewConstantGenerator(templates.GemfileTemplate, "Gemfile"),
	}

	for _, g := range generators {
		s, err := g.Generate()
		if err != nil {
			return err
		}

		filePath := fmt.Sprintf("%s/%s", outputDir, g.Name())
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
