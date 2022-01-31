package utils

import (
	types "internal/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func ExecAction(action types.Action, destination string) error {
	err := Shellout(action["run"])

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

	err := Shellout(action["option-"+selectedOption])

	if err != nil {
		return err
	}

	return nil
}

func ReplaceAction(action types.Action, destination string) error {
	answer := ""
	prompt := &survey.Input{
		Message: action["prompt"],
	}
	survey.AskOne(prompt, &answer)

	filePatterns := strings.Split(action["files"], ",")
	for _, filePattern := range filePatterns {
		matchingFiles, err := filepath.Glob(strings.Trim(filePattern, " "))
		if err != nil {
			return err
		}

		for _, file := range matchingFiles {
			read, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}

			newContents := strings.Replace(string(read), action["to_replace"], answer, -1)

			err = ioutil.WriteFile(file, []byte(newContents), 0)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func DeleteAction(action types.Action, destination string) error {
	filePatterns := strings.Split(action["files"], ",")

	for _, filePattern := range filePatterns {
		filesToDelete, err := filepath.Glob(strings.Trim(filePattern, " "))
		if err != nil {
			return err
		}

		for _, file := range filesToDelete {
			if err := os.Remove(file); err != nil {
				return err
			}
		}
	}

	return nil
}
