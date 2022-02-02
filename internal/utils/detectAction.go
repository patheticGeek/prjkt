package utils

import (
	"os"
	"path"
)

/**
Detect if a default action can be run.
This is called if there is no prjkt.yaml file and no action specified.

Returns the name of default action which can be run.
*/
func DetectAction(destination string) string {
	// Check for a node project
	// Check if package.json exists
	if _, err := os.Stat(path.Join(destination, "package.json")); err == nil {
		return "node"
	}

	// Check for a python project
	// Check if requirements.txt exists
	if _, err := os.Stat(path.Join(destination, "requirements.txt")); err == nil {
		return "python"
	}

	// Check for a go project
	// Check if requirements.txt exists
	if _, err := os.Stat(path.Join(destination, "go.mod")); err == nil {
		return "go"
	}

	return ""
}
