package gmock

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

// server's root dir
var RootDir string

type Handler struct {
}

func allowCORS(res http.ResponseWriter, req *http.Request) (skip bool) {
	origin := req.Header.Get("Origin")

	if origin == "" {
		origin = "*"
	}

	res.Header().Set("Access-Control-Allow-Origin", origin)
	res.Header().Set("Access-Control-Allow-Credentials", "true")
	res.Header().Set("Access-Control-Allow-Headers", "*")
	res.Header().Set("Access-Control-Allow-Methods", strings.Join([]string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}, ","))

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
		err         error
		statusCode  = 200
		data        *Schema
		contentType string
	)

	res.Header().Set("X-Powered-By", "github.com/axetroy/gmock")

	if skip := allowCORS(res, req); skip {
		return
	}

	defer func() {
		if err != nil {
			if statusCode == http.StatusOK {
				statusCode = http.StatusInternalServerError
			}
			res.WriteHeader(statusCode)
			_, _ = res.Write([]byte(fmt.Sprintf("%+v", err)))
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

	if data, contentType, err = Render(req); err != nil {
		if os.IsNotExist(err) {
			statusCode = http.StatusNotFound
		}
		return
	}

	if data.Status != nil {
		statusCode = *data.Status
	}

	headers := res.Header()

	if data.Headers != nil {
		v := reflect.ValueOf(data.Headers)

		if v.Kind() == reflect.Map {
			for _, key := range v.MapKeys() {
				strct := v.MapIndex(key)

				k := fmt.Sprintf("%v", key.Interface())

				if val, ok := (strct.Interface()).(string); ok {
					headers.Set(k, val)
				} else if values, ok := (strct.Interface()).([]interface{}); ok {
					for _, value := range values {
						headers.Add(k, fmt.Sprintf("%v", value))
					}
				}

			}
		} else {
			// invalid format for header
			// invalid header
			statusCode = http.StatusInternalServerError
			err = fmt.Errorf("invalid headers: `%v`", data.Headers)
			return
		}
	}

	// if not set the response's content type
	// Then automatically detect the type
	if headers != nil && len(headers.Values("Content-Type")) == 0 {
		headers.Set("Content-Type", contentType)
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
