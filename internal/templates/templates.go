package templates

import (
	_ "embed"
)

//go:embed index.md
var IndexTemplate string

//go:embed _config.yml
var ConfigTemplate string

//go:embed example.md
var ExampleTemplate string
