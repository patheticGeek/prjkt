package utils

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func CloneRepo(destination, url string, preserveGit bool) error {
	finalUrl, ref := ParseUrl(url)

	err := CloneWithGit(destination, finalUrl, ref, preserveGit)
	if err != nil {
		return err
	}

	return nil
}

func CloneWithGit(destination, url string, ref plumbing.ReferenceName, preserveGit bool) error {
	fmt.Println("🔃 Cloning repo")

	var depth int
	if !preserveGit {
		depth = 1
	}

	// Clone the repo
	_, err := git.PlainClone(destination, false, &git.CloneOptions{
		URL:           url,
		Depth:         depth,
		ReferenceName: ref,
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
