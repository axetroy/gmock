package app

import (
	"net/http"
	"path"
	"strings"
)

func Lookup(rootDir string, req *http.Request) string {
	pathArr := strings.Split(req.URL.Path, "/")

	filepath := rootDir

	for _, p := range pathArr {
		filepath = path.Join(filepath, p)
	}

	filepath = filepath + "." + strings.ToLower(req.Method) + ".json"

	return filepath
}
