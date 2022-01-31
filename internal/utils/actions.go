package utils

import (
	"fmt"
	types "internal/types"
	"os"
	"path/filepath"
	"strings"
)

func ExecAction(action types.Action, destination string) error {
	out, _, err := Shellout(action["run"])
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
