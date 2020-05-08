package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

// server's root dir
var RootDir string

type handler struct {
}

type Schema struct {
	Status  *int                 `json:"status"`  // 返回的状态码
	Body    interface{}          `json:"body"`    // 请求体
	Headers *map[string][]string `json:"headers"` // 返回头
}

func (h handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var (
		err        error
		statusCode = 200
	)

	defer func() {
		if err != nil {
			_, _ = res.Write([]byte(err.Error()))
		}
	}()

	filepath := Lookup(RootDir, req.Method, req.URL)

	// if file not exist
	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		statusCode = 404
		err = errors.New(http.StatusText(http.StatusNotFound))
		return
	}

	bytes, er := ioutil.ReadFile(filepath)

	if er != nil {
		statusCode = http.StatusInternalServerError
		err = er
		return
	}

	data := Schema{}

	if er := json.Unmarshal(bytes, &data); err != nil {
		statusCode = http.StatusInternalServerError
		err = er
		return
	}

	body, er := json.Marshal(data.Body)

	if er != nil {
		statusCode = http.StatusInternalServerError
		err = er
		return
	}

	if data.Status != nil {
		statusCode = *data.Status
	}

	if data.Headers != nil {
		for key, values := range *data.Headers {
			for _, value := range values {
				res.Header().Add(key, value)
			}
		}
	}

	res.WriteHeader(statusCode)

	_, _ = res.Write(body)
}

func Server(addr string, targetDir string) error {
	// if root path is relative
	if !path.IsAbs(targetDir) {
		cwd, _ := os.Getwd()
		targetDir = path.Join(cwd, RootDir)
	}

	RootDir = targetDir

	s := &http.Server{
		Addr:           addr,
		Handler:        handler{},
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 10M
	}

	log.Printf("Root Dir: %s", RootDir)

	log.Printf("Listen on %s\n.", addr)

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
