package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/axetroy/gmock/internal/app/function"
	"github.com/axetroy/gmock/internal/lib/mock"
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
		"makeSlice":         function.FuncMakeSlice,
		"makeSliceByLength": function.MakeSliceByLength,
		// math
		"plusInt":    function.PlusInt,
		"plusFloat":  function.PlusFloat,
		"minusInt":   function.MinusInt,
		"minusFloat": function.MinusFloat,
		"timesInt":   function.TimesInt,
		"timesFloat": function.TimesFloat,
		"divInt":     function.DivInt,
		"divFloat":   function.DivFloat,
		// random
		"randomStr":  function.RandomStr,
		"rangeInt":   function.RangeInt,
		"rangeFloat": function.RangeFloat,
	}).Parse(string(input)); err != nil {
		return err
	} else {
		if err := t.Execute(output, context); err != nil {
			return err
		}
	}

	return nil
}

// return file path & content & status code & error
func Render(req *http.Request) (*Schema, error) {
	var (
		result = Schema{}
		buff   = bytes.NewBuffer(nil)
		reader io.Reader
	)
	filepath, routeParams := Lookup(RootDir, req.Method, req.URL)

	context := map[string]interface{}{
		"Request": req,         // The request object
		"Params":  routeParams, // The Params of Route
		"Faker":   mock.Mock{},
	}

	if filepath == nil {
		return nil, os.ErrNotExist
	}

	// if file not exist
	if _, err := os.Stat(*filepath); os.IsNotExist(err) {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}

	if b, err := ioutil.ReadFile(*filepath); err != nil {
		return nil, err
	} else {
		if err := rend(req.URL.Path, context, b, buff); err != nil {
			return nil, err
		}
	}

	if err := json.Unmarshal(buff.Bytes(), &result); err != nil {
		return nil, err
	}

	if str, ok := result.Body.(string); ok {
		// hack file proto
		if strings.HasPrefix(str, "file://") || strings.HasPrefix(str, "template://") {
			isTemplate := strings.HasPrefix(str, "template://")

			var redirect string

			if isTemplate {
				redirect = strings.TrimPrefix(str, "template://")
			} else {
				redirect = strings.TrimPrefix(str, "file://")
			}

			var targetFile string

			if path.IsAbs(redirect) {
				targetFile = redirect
			} else {
				targetFile = path.Join(path.Dir(*filepath), redirect)
			}

			if isTemplate {
				if tpl, err := ioutil.ReadFile(targetFile); err != nil {
					return nil, err
				} else {
					// reset to empty
					buff = bytes.NewBuffer(nil)
					// compile template
					if err := rend(req.URL.Path, context, tpl, buff); err != nil {
						return nil, err
					}
				}
			} else {
				if fileStat, err := os.Open(targetFile); err != nil {
					return nil, err
				} else {
					reader = fileStat
				}
			}

		} else {
			buff = bytes.NewBuffer([]byte(str))
		}
	} else {
		if data, err := json.Marshal(result.Body); err != nil {
			return nil, err
		} else {
			buff = bytes.NewBuffer(data)
		}
	}

	if reader != nil {
		result.Body = &reader
	} else {
		result.Body = buff.Bytes()
	}

	return &result, nil
}
