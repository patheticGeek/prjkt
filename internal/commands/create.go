package commands

import (
	"fmt"
	"internal/utils"

	"github.com/urfave/cli/v2"
)

var CreateProjectCommand = cli.Command{
	Name:    "create",
	Usage:   "Create a project with repo",
	Aliases: []string{"c"},
	Action:  CreateProject,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "url",
			Aliases:  []string{"u"},
			Required: true,
			Usage:    "The url of the repo to create project from",
		},
		&cli.StringFlag{
			Name:    "destination",
			Aliases: []string{"d"},
			Usage:   "Path of the destination folder to clone to",
			Value:   "new-prjkt",
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
}

func CreateProject(c *cli.Context) error {
	url := c.String("url")
	destination := c.String("destination")
	preserveGit := c.Bool("preserve-git")

	err := utils.CloneWithGit(destination, url, preserveGit)
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("✨ Happy hacking!")

	return nil
}
