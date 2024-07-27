package main

import (
	"context"

	"pages/internal/commands"
)

func main() {
	var command commands.GeneratePagesCommand

	command.Init()

	if err := command.Run(context.Background()); err != nil {
		panic(err)
	}
}
