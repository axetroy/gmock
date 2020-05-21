package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// server's root dir
var RootDir string

type Handler struct {
}

type Schema struct {
	Status  *int                 `json:"status"`  // 返回的状态码
	Body    interface{}          `json:"body"`    // 请求体
	Headers *map[string][]string `json:"headers"` // 返回头
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var (
		err        error
		statusCode = 200
		filepath   string
		fileBytes  []byte
	)

	defer func() {
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(err.Error()))
		}
	}()

	if filepath, fileBytes, statusCode, err = Render(req); err != nil {
		return
	}

	data := Schema{}

	if er := json.Unmarshal(fileBytes, &data); err != nil {
		statusCode = http.StatusInternalServerError
		err = er
		return
	}

	var body []byte

	if str, ok := data.Body.(string); ok {
		// hack file proto
		if strings.HasPrefix(str, "file://") {
			redirect := strings.TrimPrefix(str, "file://")

			var target string

			if path.IsAbs(redirect) {
				target = redirect
			} else {
				target = path.Join(path.Dir(filepath), redirect)
			}

			if b, err := ioutil.ReadFile(target); err != nil {
				if os.IsExist(err) {
					statusCode = http.StatusNotFound
				} else {
					statusCode = http.StatusInternalServerError
				}
			} else {
				body = b
			}
		} else {
			body = []byte(str)
		}
	} else {
		if b, err := json.Marshal(data.Body); err != nil {
			statusCode = http.StatusInternalServerError
		} else {
			body = b
		}
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
		targetDir = path.Join(cwd, targetDir)
	}

	RootDir = targetDir

	s := &http.Server{
		Addr:           addr,
		Handler:        Handler{},
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
