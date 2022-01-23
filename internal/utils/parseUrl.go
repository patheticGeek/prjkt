package utils

import (
	"fmt"
	"regexp"
)

func ParseUrl(url string) (string, string) {
	var finalUrl string
	var ref string

	urlPattern := regexp.MustCompile(`^(?:(?:https:\/\/)?([^:/]+\.[^:/]+)\/|git@([^:/]+)[:/]|([^/]+):)?([^/\s]+)\/([^/\s#]+)(?:((?:\/[^/\s#]+)+))?(?:\/)?(?:#(.+))?`)

	fmt.Println(urlPattern.MatchString(url))

	return finalUrl, ref
}
