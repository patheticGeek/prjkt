package commands

import (
	"encoding/json"
	"fmt"
	"internal/types"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/urfave/cli/v2"
)

var ListActionsCommand = cli.Command{
	Name:    "list-actions",
	Usage:   "List all available default actions",
	Aliases: []string{"la"},
	Action:  ListActions,
}

func ListActions(c *cli.Context) error {
	fmt.Println("🔃 Getting the list of actions")

	// Templates are stored here: https://github.com/patheticGeek/prjkt-templates/defaults
	url := "https://api.github.com/repos/patheticGeek/prjkt-templates/contents/defaults?ref=main"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("🚨 An error occurred while fetching actions list")
	}
	defer resp.Body.Close()

	fileData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data types.GitHubFilesListResp
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		return err
	}

	fmt.Println("\nAvailable actions:")
	for _, item := range data {
		fmt.Println(strings.TrimRight(item.Name, ".yaml"))
	}

	return nil
}
