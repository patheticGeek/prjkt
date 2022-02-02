package main

import (
	"log"
	"os"

	commands "internal/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "prjkt",
		Usage:   "Project creation simplified",
		Version: "v0.3.0",
		Authors: []*cli.Author{
			{
				Name:  "Pathetic Geek",
				Email: "geekpathetic@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Value: false,
				Usage: "Show debug logs",
			},
		},
		Commands: []*cli.Command{
			&commands.CreateProjectCommand,
			&commands.ListActionsCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
