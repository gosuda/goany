package goanyecho

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rabbitprincess/goany/goany"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	e := echo.New()

	e.POST("/hello", WithAny(func(req goany.Request, res goany.Response) error {
		name := req.Path("user.name").String()
		age := req.Get("user").Get("age").Int()
		res.Set("message", name).Set("age", age)
		return nil
	}))

	payload := []byte(`{"user":{"name":"Alice", "age":30}}`)
	req := httptest.NewRequest(http.MethodPost, "/hello", bytes.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	respond := goany.NewRequest(rec.Body.Bytes())
	require.Equal(t, "Alice", respond.Path("message").String())
	require.Equal(t, "30", respond.Path("age").String())
}
