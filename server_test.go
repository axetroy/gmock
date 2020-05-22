package main_test

import (
	"encoding/json"
	"github.com/axetroy/gmock"
	"github.com/axetroy/mocker"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {
	cwd, _ := os.Getwd()

	{
		main.RootDir = path.Join(cwd, "__test__", "hello_world")
		mock := mocker.New(main.Handler{})

		r := mock.Get("/hello", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)
		assert.Equal(t, "hello world!", string(body))
	}

	{
		main.RootDir = path.Join(cwd, "__test__", "params")
		mock := mocker.New(main.Handler{})

		{
			r := mock.Get("/user/123", nil, nil)

			assert.Equal(t, http.StatusOK, r.Code)
			body, err := ioutil.ReadAll(r.Body)

			assert.Nil(t, err)

			b, _ := json.Marshal(map[string]interface{}{
				"id":       123,
				"username": "root",
			})

			assert.Equal(t, string(b), string(body))
		}

		{
			r := mock.Get("/zoo_id/detail", nil, nil)

			assert.Equal(t, http.StatusOK, r.Code)
			body, err := ioutil.ReadAll(r.Body)

			assert.Nil(t, err)

			b, _ := json.Marshal(map[string]interface{}{
				"id":        123,
				"card_name": "This is card name",
			})

			assert.Equal(t, string(b), string(body))
		}

		{
			r := mock.Get("/zoo_id/id", nil, nil)

			assert.Equal(t, http.StatusOK, r.Code)
			body, err := ioutil.ReadAll(r.Body)

			assert.Nil(t, err)

			assert.Equal(t, "you id is: zoo_id", string(body))
		}
	}

	{
		main.RootDir = path.Join(cwd, "__test__", "user_context_with_template")
		mock := mocker.New(main.Handler{})

		r := mock.Get("/hello_id", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		assert.Equal(t, "/hello_id", string(body))
	}

	{
		main.RootDir = path.Join(cwd, "__test__", "status_code")
		mock := mocker.New(main.Handler{})

		r := mock.Get("/error", nil, nil)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		assert.Equal(t, "error", string(body))
	}

	{
		main.RootDir = path.Join(cwd, "__test__", "loop_output")
		mock := mocker.New(main.Handler{})

		r := mock.Get("/array", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		arr := []int{1, 2, 3}

		b, _ := json.Marshal(arr)

		assert.Equal(t, b, body)
	}

	{
		main.RootDir = path.Join(cwd, "__test__", "loop_output")
		mock := mocker.New(main.Handler{})

		r := mock.Get("/function", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		arr := []int{1, 2, 3}

		b, _ := json.Marshal(arr)

		assert.Equal(t, b, body)
	}
}

func TestServerExample(t *testing.T) {
	cwd, _ := os.Getwd()
	main.RootDir = path.Join(cwd, "example")

	mock := mocker.New(main.Handler{})

	// GET /hello
	{
		r := mock.Get("/hello", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)
		assert.Equal(t, "hello world!", string(body))
	}

	//// GET /
	//{
	//	r := mock.Get("/", nil, nil)
	//
	//	assert.Equal(t, http.StatusOK, r.Code)
	//	body, err := ioutil.ReadAll(r.Body)
	//
	//	assert.Nil(t, err)
	//	assert.Equal(t, "root path", string(body))
	//}

	// GET /home
	{
		r := mock.Get("/home", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		b, err := ioutil.ReadFile(path.Join(main.RootDir, "home.html"))

		assert.Nil(t, err)
		assert.Equal(t, string(b), string(body))
	}

	// GET /template
	{
		r := mock.Get("/template", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)
		assert.Equal(t, `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Template</title>
</head>
<body>
  <p>Your request URL path: /template</p>
</body>
</html>`, strings.ReplaceAll(string(body), "\r\n", "\n"))
	}

	// GET /avatar
	{
		r := mock.Get("/avatar", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "image/jpeg", r.Header().Get("Content-Type"))
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		b, err := ioutil.ReadFile(path.Join(main.RootDir, "avatar.jpeg"))

		assert.Nil(t, err)
		assert.Equal(t, b, body)
	}

	// GET /avatar_download
	{
		r := mock.Get("/avatar_download", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "image/jpeg", r.Header().Get("Content-Type"))
		assert.Equal(t, "attachment;filename=avatar.jpeg", r.Header().Get("Content-Disposition"))
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		b, err := ioutil.ReadFile(path.Join(main.RootDir, "avatar.jpeg"))

		assert.Nil(t, err)
		assert.Equal(t, b, body)
	}
}
