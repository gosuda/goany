package goanygin

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rabbitprincess/goany/goany"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.POST("/hello", WithAny(func(req *goany.Request, res *goany.Response) error {
		name := req.Path("user.name").String()
		age := req.Get("user").Get("age").Int()

		res.Set("message", name).Set("age", age)
		return nil
	}))

	payload := []byte(`{"user":{"name":"Alice", "age":30}}`)
	req := httptest.NewRequest("POST", "/hello", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	require.NotNil(t, rec)
	require.Equal(t, http.StatusOK, rec.Code)

	// validate the response
	respond := goany.NewRequest(rec.Body.Bytes())
	require.Equal(t, "Alice", respond.Path("message").String())
	require.Equal(t, "30", respond.Path("age").String())
}
