package app

import (
	"io"
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

func allowCORS(res http.ResponseWriter, req *http.Request) (skip bool) {
	res.Header().Set("Access-Control-Allow-Origin", res.Header().Get("Origin"))
	res.Header().Set("Access-Control-Allow-Credentials", res.Header().Get("true"))
	res.Header().Set("Access-Control-Allow-Methods", res.Header().Get(strings.Join([]string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}, ",")))

	if req.Method == http.MethodOptions {
		res.WriteHeader(http.StatusNoContent)
		_, _ = res.Write(nil)
		skip = true
		return
	}

	return skip
}

func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var (
		err        error
		statusCode = 200
		data       *Schema
	)

	if skip := allowCORS(res, req); skip {
		return
	}

	defer func() {

		if err != nil {
			if statusCode == http.StatusOK {
				statusCode = http.StatusInternalServerError
			}
			res.WriteHeader(statusCode)
			_, _ = res.Write([]byte(err.Error()))
		} else {
			res.WriteHeader(statusCode)
			if data != nil {
				if b, ok := data.Body.([]byte); ok {
					_, _ = res.Write(b)
				} else if reader, ok := data.Body.(*io.Reader); ok {
					_, _ = io.Copy(res, *reader)
				} else if str, ok := data.Body.(string); ok {
					_, _ = res.Write([]byte(str))
				} else {
					_, _ = res.Write(nil)
				}
			} else {
				_, _ = res.Write(nil)
			}
		}
	}()

	if data, err = Render(req); err != nil {
		if os.IsNotExist(err) {
			statusCode = http.StatusNotFound
		}
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

	log.Printf("Listen on %s.\n", addr)

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
