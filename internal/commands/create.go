package commands

import (
	"errors"
	"fmt"
	"internal/utils"

	"github.com/AlecAivazis/survey/v2"
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
prjkt create user/repo@branch
prjkt create user/repo@branch test-project
`

var CreateProjectCommand = cli.Command{
	Name:        "create",
	Usage:       "Create a project with repo",
	Description: description,
	Aliases:     []string{"c"},
	Action:      CreateProject,
	ArgsUsage:   "[url] [destination]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "action",
			Aliases: []string{"a"},
			Value:   "",
			Usage:   "Name of the predefined action you want to run",
		},
		&cli.BoolFlag{
			Name:    "preserve-git",
			Aliases: []string{"pg"},
			Value:   false,
			Usage:   "Preserve git in the created folder",
		},
		&cli.BoolFlag{
			Name:    "no-actions",
			Aliases: []string{"na"},
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
	if c.Args().Len() == 0 {
		return errors.New("❌ No url passed in to create from")
	}

	url := c.Args().Get(0)
	destination := "new-prjkt"
	if c.Args().Len() > 1 {
		arg := c.Args().Get(1)
		// Check if this is not a flag (-a, --action, etc)
		if arg[0] != '-' {
			destination = c.Args().Get(1)
		}
	}
	preserveGit := c.Bool("preserve-git")
	noActions := c.Bool("no-actions")
	action := c.String("action")

	err := utils.CloneRepo(destination, url, preserveGit)
	if err != nil {
		fmt.Println("🚨 An error occurred cloning the repo")
		return err
	}

	var fileData []byte

	// If user passed in a default action get that otherwise get repo's prjkt.yaml
	if action != "" {
		fileData, err = utils.GetDefaultAction(action)
	} else {
		fileData, err = utils.ReadRepoPrjktYAML(destination)
	}

	// If there was no error and length is 0, that means there's no file/default action
	// And if yser hasn't specified no-actions, try and detect if a default action can be run
	// If a default action can be run ask user if they want to run it otherwise exit
	if err == nil && len(fileData) == 0 && !noActions {
		action = utils.DetectAction(destination)

		if action != "" {
			fmt.Println("")
			runAction := false
			prompt := &survey.Confirm{
				Message: "✨ " + action + " project detected, do you want to run default actions for it?",
			}
			survey.AskOne(prompt, &runAction)

			if runAction {
				fileData, err = utils.GetDefaultAction(action)
			} else {
				fmt.Println("")
				fmt.Println("ℹ️ No prjkt.yaml or default action found skipping")
				printDefaultSuccessMessage()
				return nil
			}
		}
	} else if err != nil {
		// If there was an error it means the file/default action was not fond
		fmt.Println("🚨 An error occurred reading actions file")
		return err
	}

	// Parse the prjkt.yaml file data
	result, err := utils.ParsePrjktYAML(fileData)
	if err != nil {
		fmt.Println("🚨 An error occurred parsing actions file")
		return err
	}

	if result.Welcome_message != "" {
		fmt.Println("")
		fmt.Println(result.Welcome_message)
	}

	// If user hasn't set no-actions flag then run actions if any
	if !noActions {
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
	} else {
		if result.No_actions_message != "" {
			fmt.Println(result.No_actions_message)
		} else {
			printDefaultSuccessMessage()
		}
	}

	return nil
}
