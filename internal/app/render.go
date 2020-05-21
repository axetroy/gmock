package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/axetroy/gmock/internal/app/function"
	"github.com/axetroy/gmock/internal/lib/mock"
)

type RenderStat struct {
	filepath string
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

			var target string

			if path.IsAbs(redirect) {
				target = redirect
			} else {
				target = path.Join(path.Dir(*filepath), redirect)
			}

			if tpl, err := ioutil.ReadFile(target); err != nil {
				return nil, err
			} else {
				if isTemplate {
					// reset to empty
					buff = bytes.NewBuffer(nil)
					// compile template
					if err := rend(req.URL.Path, context, tpl, buff); err != nil {
						return nil, err
					}
				} else {
					buff = bytes.NewBuffer(tpl)
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

	result.Body = buff.Bytes()

	return &result, nil
}
