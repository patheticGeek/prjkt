package commands

import (
	"errors"
	"fmt"
	"internal/utils"
	"io/ioutil"
	"os"
	"path/filepath"

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

func printDefaultSuccessMessage() {
	fmt.Println("")
	fmt.Println("✨ Happy hacking!")
}

func CreateProject(c *cli.Context) error {
	url := c.String("url")
	destination := c.String("destination")
	preserveGit := c.Bool("preserve-git")

	err := utils.CloneRepo(destination, url, preserveGit)
	if err != nil {
		fmt.Println("🚨 An error occurred cloning the repo")
		return err
	}

	prjktYAMLPath := filepath.Join(destination, "/prjkt.yaml")

	// Check if prjkt.yaml exists, if it doesn't exit
	if _, err := os.Stat(prjktYAMLPath); errors.Is(err, os.ErrNotExist) {
		printDefaultSuccessMessage()
		return nil
	}

	fileData, err := ioutil.ReadFile(prjktYAMLPath)

	if err != nil {
		fmt.Println("🚨 An error occurred reading actions file")
		return err
	}

	result, err := utils.ParsePrjktYAML(fileData)
	if err != nil {
		fmt.Println("🚨 An error occurred parsing actions file")
		return err
	}

	if result.Welcome_message != "" {
		fmt.Println("")
		fmt.Println(result.Welcome_message)
	}

	err = utils.RunActions(result.Actions, destination)

	if err != nil {
		fmt.Println("🚨 An error occurred while running the actions")
		if result.Error_message != "" {
			fmt.Println(result.Error_message)
		}
		return err
	}

	if result.Success_message != "" {
		fmt.Println(result.Success_message)
	} else {
		printDefaultSuccessMessage()
	}

	return nil
}
