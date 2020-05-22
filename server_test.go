package gmock_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/axetroy/gmock"
	"github.com/axetroy/mocker"
	"github.com/stretchr/testify/assert"
)

func TestServerExample(t *testing.T) {
	cwd, _ := os.Getwd()
	gmock.RootDir = path.Join(cwd, "example")

	mock := mocker.New(gmock.Handler{})

	// GET /hello
	{
		r := mock.Get("/hello", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)
		assert.Equal(t, "hello world!", string(body))
	}

	// GET /
	{
		r := mock.Get("/", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)
		assert.Equal(t, "root path", string(body))
	}

	// GET /home
	{
		r := mock.Get("/home", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		b, err := ioutil.ReadFile(path.Join(gmock.RootDir, "home.html"))

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
</html>`, string(body))
	}

	// GET /avatar
	{
		r := mock.Get("/avatar", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "image/jpeg", r.Header().Get("Content-Type"))
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		b, err := ioutil.ReadFile(path.Join(gmock.RootDir, "avatar.jpeg"))

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

		b, err := ioutil.ReadFile(path.Join(gmock.RootDir, "avatar.jpeg"))

		assert.Nil(t, err)
		assert.Equal(t, b, body)
	}

	// GET /user/:user_id
	{
		r := mock.Get("/user/321", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		assert.Equal(t, "application/json", r.Header().Get("Content-Type"))
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		b, err := json.Marshal(map[string]interface{}{
			"uid": 123,
		})

		assert.Nil(t, err)
		assert.Equal(t, string(b), string(body))
	}
}
