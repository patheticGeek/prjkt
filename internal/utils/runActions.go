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

	// Run commands in the project dir
	wd, _ := os.Getwd()
	projectPath := filepath.Join(wd, destination)
	os.Chdir(projectPath)

	for i, action := range actions {
		fmt.Println(strconv.Itoa(i) + ". " + action["name"])

		var err error

		switch action["type"] {
		case "exec":
			err = ExecAction(action, destination)
		default:
			fmt.Println("Invalid action type: " + action["type"])
		}

		if err != nil && action["continue"] != "true" {
			return err
		}
	}

	fmt.Println("")

	return nil
}
