package commands

import (
	"fmt"
	"internal/utils"

	"github.com/urfave/cli/v2"
)

const description string = `
The value of url can be as following:

https://github.com/user/repo
https://bitbucket.org/user/repo
https://github.com/user/repo#tag
https://github.com/user/repo@branch

Shorthands: (https://github.com is assumed)
user/repo
user/repo#tag
user/repo@branch

Example:
prjkt create -u user/repo@branch -d repo-with-branch
`

var CreateProjectCommand = cli.Command{
	Name:        "create",
	Usage:       "Create a project with repo",
	Description: description,
	Aliases:     []string{"c"},
	Action:      CreateProject,
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
			Usage:   "Preserve git in the created folder",
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

	err := utils.CloneRepo(destination, url, preserveGit)
	if err != nil {
		fmt.Println("")
		fmt.Println("🚨 Some error occurred")
		return err
	}

	fmt.Println("")
	fmt.Println("✨ Happy hacking!")

	return nil
}
