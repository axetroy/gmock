package app

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

func NameToRegExp(name string) *regexp.Regexp {

	paramsRegExp := regexp.MustCompile("\\[[\\w\\d-]+\\]")

	name = paramsRegExp.ReplaceAllString(name, "gmock_params")

	// escapeRegExp := regexp.MustCompile("[.*+\\-?^${}()|[\\]\\\\]")
	escapes := []string{"\\", ".", "*", "+", "-", "?", "^", "$", "{", "}", "[", "]"}

	for _, e := range escapes {
		name = strings.ReplaceAll(name, e, "\\"+e)
	}

	name = strings.ReplaceAll(name, "gmock_params", "[\\w\\d-]+")

	return regexp.MustCompile("^" + name + "$")
}

func MatchFileName(rule, name string) bool {
	return NameToRegExp(rule).MatchString(name)
}

func Lookup(rootDir string, method string, u *url.URL) *string {
	if u.Path == "/favicon.ico" {
		return nil
	}
	method = strings.ToLower(method)
	pathArr := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")

	currentDir := rootDir

	for index, p := range pathArr {
		filepath := path.Join(currentDir, p)
		isLastElement := len(pathArr)-1 == index

		// If it is the last element, then the file should have a extension name
		if isLastElement {
			filepath = filepath + "." + method + ".json"
		} else {
			currentDir = path.Join(currentDir, p)
		}

		// if file not exist
		if f, err := os.Stat(filepath); os.IsNotExist(err) {
			// if path not found. try read the params route
			currentDir = path.Dir(filepath)
			files, _ := ioutil.ReadDir(currentDir)

			// Find matching items in a directory of the same level
		flatten:
			for _, f := range files {
				f2, _ := os.Stat(path.Join(currentDir, f.Name()))

				if !isLastElement && f2.IsDir() {
					if MatchFileName(f.Name(), p) {
						currentDir = path.Join(currentDir, f.Name())
						break flatten
					}
				} else {
					if MatchFileName(f.Name(), p+"."+method+".json") {
						filepath = path.Join(currentDir, f.Name())
						return &filepath
					}
				}

			}
		} else {
			if !f.IsDir() {
				return &filepath
			}
		}
	}

	return nil
}
