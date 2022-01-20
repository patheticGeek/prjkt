package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
)

func createProject(c *cli.Context) error {
	url := c.String("url")
	destination := c.String("destination")
	preserveGit := c.Bool("preserve-git")

	fmt.Println("Cloning repo:", url)
	fmt.Println("Destination:", destination)
	fmt.Println("Preserve git:", preserveGit)
	fmt.Println("")

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

	if !preserveGit {
		// Delete the git folder
		os.RemoveAll(destination + "/.git")
	}

	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("✨ Done")

	return nil
}

func main() {
	app := &cli.App{
		Name:  "prjkt",
		Usage: "Project creation simplified",
		Commands: []*cli.Command{
			{
				Name:        "create",
				Description: "Create a project with repo",
				Aliases:     []string{"c"},
				Action:      createProject,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "url",
						Aliases:  []string{"u"},
						Required: true,
						Usage:    "The url of the repo to create project from",
					},
					&cli.StringFlag{
						Name:    "destination",
						Aliases: []string{"d"},
						Usage:   "Path of the destination folder to clone to",
						Value:   "new-prjkt",
					},
					&cli.BoolFlag{
						Name:    "preserve-git",
						Aliases: []string{"pg"},
						Value:   false,
						Usage:   "Clone repo normally instead of getting the latest code and deleting .git",
					},
					&cli.BoolFlag{
						Name:    "no-actions",
						Aliases: []string{"a"},
						Value:   false,
						Usage:   "If you don't want to run any actions",
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
