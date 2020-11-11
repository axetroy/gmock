package gmock

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

const ROUTER_PARAMS_PLACEHOLDER = "__ROUTER_PARAMS_PLACEHOLDER__"

func NameToRegExp(name string) *regexp.Regexp {

	paramsRegExp := regexp.MustCompile(`\[[\w\d-]+\]`)

	name = paramsRegExp.ReplaceAllString(name, ROUTER_PARAMS_PLACEHOLDER)

	// escapeRegExp := regexp.MustCompile("[.*+\\-?^${}()|[\\]\\\\]")
	escapes := []string{"\\", ".", "*", "+", "-", "?", "^", "$", "{", "}", "[", "]"}

	for _, e := range escapes {
		name = strings.ReplaceAll(name, e, "\\"+e)
	}

	name = strings.ReplaceAll(name, ROUTER_PARAMS_PLACEHOLDER, "[\\w\\d-]+")

	return regexp.MustCompile("^" + name + "$")
}

func MatchFileName(rule, name string) bool {
	return NameToRegExp(rule).MatchString(name)
}

func GetRealFileName(fileName string) string {
	fileName = strings.TrimSuffix(fileName, ".json")
	fileName = strings.TrimSuffix(fileName, path.Ext(fileName))

	return fileName
}

func ExtractParamsFromFileName(fileName string, urlPath string) map[string]string {
	fileName = GetRealFileName(fileName)

	params := map[string]string{}

	paramsRegExp := regexp.MustCompile(`\[([\w\d-]+)\]`)

	matchers := paramsRegExp.FindAllStringSubmatch(fileName, -1)

	if len(matchers) < 1 {
		return params
	}

	for _, m := range matchers {
		paramsName := m[1]
		params[paramsName] = urlPath
	}

	return params
}

// Lookup router file and params context
func Lookup(rootDir string, method string, u *url.URL) (*string, map[string]string) {
	method = strings.ToLower(method)

	if u.Path == "/" {
		target := path.Join(rootDir, "."+method+".json")
		if _, err := os.Stat(target); os.IsNotExist(err) {
			return nil, nil
		} else {
			return &target, map[string]string{}
		}
	}

	if u.Path == "/favicon.ico" {
		return nil, nil
	}

	pathArr := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")

	routeParams := map[string]string{}

	currentDir := rootDir

	for index, pathName := range pathArr {
		filepath := path.Join(currentDir, pathName)
		isLastElement := len(pathArr)-1 == index

		// If it is the last element, then the file should have a extension name
		if isLastElement {
			filepath = filepath + "." + method + ".json"
		} else {
			currentDir = path.Join(currentDir, pathName)
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
					if MatchFileName(f.Name(), pathName) {
						currentDir = path.Join(currentDir, f.Name())
						paramName := strings.TrimPrefix(f.Name(), "[")
						paramName = strings.TrimSuffix(paramName, "]")
						routeParams[paramName] = pathName
						break flatten
					}
				} else {
					if MatchFileName(f.Name(), pathName+"."+method+".json") {
						filepath = path.Join(currentDir, f.Name())

						paramName := strings.TrimSuffix(f.Name(), ".json")
						paramName = strings.TrimSuffix(paramName, "."+method)
						paramName = strings.TrimPrefix(paramName, "[")
						paramName = strings.TrimSuffix(paramName, "]")
						routeParams[paramName] = pathName
						return &filepath, routeParams
					}
				}

			}
		} else {
			if !f.IsDir() {
				return &filepath, routeParams
			}
		}
	}

	return nil, nil
}
