package gmock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/axetroy/gmock/function"
	"github.com/axetroy/gmock/lib/mock"
	"github.com/pkg/errors"
)

type Schema struct {
	Status  *int        `json:"status"`  // 返回的状态码
	Body    interface{} `json:"body"`    // 请求体
	Headers interface{} `json:"headers"` // 返回头, 可以是 map[string]string 类型，也可以是 map[string][]string
}

func rend(templateName string, context map[string]interface{}, input []byte, output *bytes.Buffer) error {
	t := template.New(templateName)

	if t, err := t.Funcs(template.FuncMap{
		// slice
		"MakeSlice":         function.MakeSlice,
		"MakeSliceByLength": function.MakeSliceByLength,
		// encoding
		"Base64Encoding": function.Base64Encoding,
		"Base64Decoding": function.Base64Decoding,
		"Escape":         function.Escape,
		// math
		"Plus":  function.Plus,
		"Minus": function.Minus,
		"Times": function.Times,
		"Div":   function.Div,
		// random
		"RandomStr":  function.RandomStr,
		"RangeInt":   function.RangeInt,
		"RangeFloat": function.RangeFloat,
	}).Parse(string(input)); err != nil {
		return err
	} else {
		if err := t.Execute(output, context); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// return file path & content & status code & error
func Render(req *http.Request) (*Schema, string, error) {
	var (
		result      = Schema{}
		buff        = bytes.NewBuffer(nil)
		reader      io.Reader
		contentType = "text/plain"
	)
	filepath, routeParams := Lookup(RootDir, req.Method, req.URL)

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, contentType, errors.WithStack(err)
	}

	var bodyMap map[string]interface{}

	_ = json.Unmarshal(body, &bodyMap)

	var bodyStr string = string(body)

	bodyStr = strings.Replace(bodyStr, `"`, `\"`, -1)

	context := map[string]interface{}{
		"Request":    req,         // The request object
		"Body":       body,        // The request body Bytes
		"BodyString": bodyStr,     // The request body String
		"BodyMap":    bodyMap,     // The request body Map
		"Params":     routeParams, // The Params of Route
		"Faker":      mock.Mock{},
	}

	if filepath == nil {
		return nil, contentType, os.ErrNotExist
	}

	// if file not exist
	if _, err := os.Stat(*filepath); os.IsNotExist(err) {
		return nil, contentType, errors.New(http.StatusText(http.StatusNotFound))
	}

	if b, err := ioutil.ReadFile(*filepath); err != nil {
		return nil, contentType, errors.WithStack(err)
	} else {
		if err := rend(req.URL.Path, context, b, buff); err != nil {
			return nil, contentType, errors.WithStack(err)
		}
	}

	if err := json.Unmarshal(buff.Bytes(), &result); err != nil {
		err = errors.Wrapf(err, "invalid JSON file for '%s'", *filepath)
		err = errors.Wrapf(err, "with file content\n=====Content Start=====\n%s\n=====Content End=====\n", buff.String())
		return nil, contentType, err
	}

	if str, ok := result.Body.(string); ok {
		// hack file proto
		reg, err := regexp.Compile(`^\w+:\/\/.+`)

		if err != nil {
			return nil, contentType, errors.WithStack(err)
		}

		if reg.MatchString(str) {
			chars := strings.Split(str, "://")

			proto := chars[0]
			params := strings.Join(chars[1:], "://")

			switch proto {
			case "file":
				var targetFile string
				if path.IsAbs(params) {
					targetFile = params
				} else {
					targetFile = path.Join(path.Dir(*filepath), params)
				}
				contentType = mime.TypeByExtension(path.Ext(targetFile))

				if fileStat, err := os.Open(targetFile); err != nil {
					return nil, contentType, errors.WithStack(err)
				} else {
					reader = fileStat
				}
			case "template":
				var targetFile string
				if path.IsAbs(params) {
					targetFile = params
				} else {
					targetFile = path.Join(path.Dir(*filepath), params)
				}

				contentType = mime.TypeByExtension(path.Ext(targetFile))

				if tpl, err := ioutil.ReadFile(targetFile); err != nil {
					return nil, contentType, errors.WithStack(err)
				} else {
					// reset to empty
					buff = bytes.NewBuffer(nil)
					// compile template
					if err := rend(req.URL.Path, context, tpl, buff); err != nil {
						return nil, contentType, errors.WithStack(err)
					}
				}
			default:
				return nil, contentType, errors.WithStack(errors.New(fmt.Sprintf("Invalid body proto '%s'", proto)))
			}

		} else {
			buff = bytes.NewBuffer([]byte(str))
		}
	} else {
		if data, err := json.Marshal(result.Body); err != nil {
			return nil, contentType, errors.WithStack(err)
		} else {
			contentType = mime.TypeByExtension(".json")
			buff = bytes.NewBuffer(data)
		}
	}

	if reader != nil {
		result.Body = &reader
	} else {
		result.Body = buff.Bytes()
	}

	return &result, contentType, nil
}
