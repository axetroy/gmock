package app

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

func Render(req *http.Request) ([]byte, int, error) {
	filepath, routeParams := Lookup(RootDir, req.Method, req.URL)

	if filepath == nil {
		return nil, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound))
	}

	// if file not exist
	if _, err := os.Stat(*filepath); os.IsNotExist(err) {
		return nil, http.StatusNotFound, errors.New(http.StatusText(http.StatusNotFound))
	}

	b, err := ioutil.ReadFile(*filepath)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	t := template.New(req.URL.Path)

	if t, err = t.Parse(string(b)); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var buff bytes.Buffer

	err = t.Execute(&buff, map[string]interface{}{
		"Request": req,         // The request object
		"Params":  routeParams, // The Params of Route
	})

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return buff.Bytes(), http.StatusOK, nil
}
