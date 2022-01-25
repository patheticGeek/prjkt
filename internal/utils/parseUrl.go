package utils

import (
	"fmt"
	"regexp"

	"github.com/go-git/go-git/v5/plumbing"
)

func ParseUrl(url string) (string, plumbing.ReferenceName) {
	var finalUrl string
	var ref plumbing.ReferenceName

	urlPattern := regexp.MustCompile(`(?:(?:https:\/\/)?([^:/]+\.[^:/]+)\/|git@([^:/]+)[:/]|([^/]+):)?([^/\s]+)\/([^/\s@#]+)(?:#(.+))?(?:@(.+))?`)

	result := urlPattern.FindStringSubmatch(url)
	// { fullUrl, empty, website, empty, username, repo, tag, branch }

	var site string = "https://github.com"
	if result[2] != "" {
		site = result[2]
	}

	finalUrl = site + "/" + result[4] + "/" + result[5]

	if result[6] != "" {
		fmt.Println(result[6])
		ref = plumbing.NewTagReferenceName(result[6])
	}

	if result[7] != "" {
		fmt.Println(result[6])
		ref = plumbing.NewBranchReferenceName(result[7])
	}

	return finalUrl, ref
}
