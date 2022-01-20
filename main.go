package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
)

func createProject(c *cli.Context) error {
	url := c.String("url")

	fmt.Println("Cloning repo: " + url)

	_, err := git.PlainClone("./", false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println("An error occurred cloning the repo")
		return err
	}

	return nil
}

func main() {
	app := &cli.App{
		Name:  "prjkt",
		Usage: "Project creation simplified",
		Commands: []*cli.Command{
			{
				Name:        "create",
				Description: "Create a project with repo",
				Aliases:     []string{"c"},
				Action:      createProject,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "url",
						Aliases:  []string{"u"},
						Required: true,
						Usage:    "The url of the repo to create project from",
					},
					&cli.BoolFlag{
						Name:    "preserve-git",
						Aliases: []string{"pg"},
						Value:   false,
						Usage:   "Clone repo normally instead of getting the latest code and deleting .git",
					},
					&cli.BoolFlag{
						Name:    "no-actions",
						Aliases: []string{"a"},
						Value:   false,
						Usage:   "If you don't want to run any actions",
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
