package app_test

import (
	"encoding/json"
	"github.com/axetroy/gmock/internal/app"
	"github.com/axetroy/mocker"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"testing"
)

func TestServer(t *testing.T) {
	cwd, _ := os.Getwd()

	{
		app.RootDir = path.Join(cwd, "__test__", "hello_world")
		mock := mocker.New(app.Handler{})

		r := mock.Get("/hello", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)
		assert.Equal(t, "hello world!", string(body))
	}

	{
		app.RootDir = path.Join(cwd, "__test__", "params")
		mock := mocker.New(app.Handler{})

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
		app.RootDir = path.Join(cwd, "__test__", "user_context_with_template")
		mock := mocker.New(app.Handler{})

		r := mock.Get("/hello_id", nil, nil)

		assert.Equal(t, http.StatusOK, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		assert.Equal(t, "/hello_id", string(body))
	}

	{
		app.RootDir = path.Join(cwd, "__test__", "status_code")
		mock := mocker.New(app.Handler{})

		r := mock.Get("/error", nil, nil)

		assert.Equal(t, http.StatusInternalServerError, r.Code)
		body, err := ioutil.ReadAll(r.Body)

		assert.Nil(t, err)

		assert.Equal(t, "error", string(body))
	}
}
