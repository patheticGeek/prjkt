package utils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func CloneWithGit(destination, url string, preserveGit bool) error {
	fmt.Println("🔃 Cloning repo")
	// finalUrl, ref := ParseUrl(url)

	var depth int
	if !preserveGit {
		depth = 1
	}

	// Clone the repo
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Depth:    depth,
	})

	if err != nil {
		return err
	}

	if !preserveGit {
		// Delete the git folder
		err = os.RemoveAll(destination + "/.git")
		if err != nil {
			return err
		}
	}

	return nil
}
