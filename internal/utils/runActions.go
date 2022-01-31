package utils

import (
	"fmt"
	"internal/types"
	"strconv"
)

// func ExecAction(action map[string]string, destination string) {}

// func ReplaceAction(action map[string]string, destination string) {}

func RunActions(actions types.Actions, destination string) error {
	if len(actions) == 0 {
		return nil
	}

	fmt.Print("\n🔧 Running actions\n\n")

	for i, action := range actions {
		fmt.Println(strconv.Itoa(i) + ". " + action["name"])

		switch action["type"] {
		default:
			fmt.Println("Invalid action type: " + action["type"])
		}
	}

	fmt.Println("")

	return nil
}
