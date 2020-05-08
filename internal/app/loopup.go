package app

import (
	"net/url"
	"path"
	"strings"
)

func Lookup(rootDir string, method string, u *url.URL) string {
	pathArr := strings.Split(u.Path, "/")

	filepath := rootDir

	for _, p := range pathArr {
		filepath = path.Join(filepath, p)
	}

	filepath = filepath + "." + strings.ToLower(method) + ".json"

	return filepath
}
