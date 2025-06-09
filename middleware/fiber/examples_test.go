package goanyfiber

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rabbitprincess/goany/goany"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	app := fiber.New()

	app.Post("/hello", WithAny(func(req goany.Request, res goany.Response) error {
		name := req.Path("user.name").String()
		age := req.Get("user").Get("age").Int()

		res.Set("message", name).Set("age", age)
		return nil
	}))

	payload := []byte(`{"user":{"name":"Alice", "age":30}}`)
	req := httptest.NewRequest("POST", "/hello", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	require.NoError(t, err)

	respond := goany.NewRequest(buf.Bytes())
	require.Equal(t, "Alice", respond.Path("message").String())
	require.Equal(t, "30", respond.Path("age").String())
}
