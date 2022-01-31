package utils

import (
	"fmt"
	"internal/types"
	"os"
	"path/filepath"
	"strconv"
)

func RunActions(actions types.ActionsMap, destination string) error {
	if len(actions) == 0 {
		return nil
	}

	fmt.Print("\n🔧 Running actions\n\n")

	// Change current dir to destination dir so commands are executed there
	wd, _ := os.Getwd()
	projectPath := filepath.Join(wd, destination)
	os.Chdir(projectPath)

	var err error

	for i, action := range actions {
		fmt.Println(strconv.Itoa(i) + ". " + action["name"])

		switch action["type"] {
		case "exec":
			err = ExecAction(action, destination)
		case "exec-option":
			err = ExecOptionAction(action, destination)
		case "delete":
			err = DeleteAction(action, destination)
		case "replace":
			err = ReplaceAction(action, destination)
		default:
			fmt.Println("Invalid action type: " + action["type"])
		}

		if err != nil && action["continue"] != "true" {
			return err
		}
	}

	fmt.Println("")

	return err
}
