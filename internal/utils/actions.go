package utils

import (
	"fmt"
	types "internal/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func ExecAction(action types.Action, destination string) error {
	out, _, err := Shellout(action["run"])
	fmt.Println(out)

	if err != nil {
		return err
	}

	return nil
}

func ExecOptionAction(action types.Action, destination string) error {
	options := []string{}
	for _, option := range strings.Split(action["options"], ",") {
		options = append(options, strings.Trim(option, " "))
	}

	selectedOption := ""
	prompt := &survey.Select{
		Message: action["prompt"],
		Options: options,
	}
	survey.AskOne(prompt, &selectedOption)

	if action["option-"+selectedOption] == "" {
		return nil
	}

	out, _, err := Shellout(action["option-"+selectedOption])
	fmt.Println(out)

	if err != nil {
		return err
	}

	return nil
}

func DeleteAction(action types.Action, destination string) error {
	files := strings.Split(action["files"], ",")

	for _, filePattern := range files {
		filesToDelete, err := filepath.Glob(strings.Trim(filePattern, " "))
		if err != nil {
			return err
		}

		for _, f := range filesToDelete {
			if err := os.Remove(f); err != nil {
				return err
			}
		}
	}

	return nil
}
