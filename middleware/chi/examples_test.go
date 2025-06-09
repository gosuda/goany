package goanychi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rabbitprincess/goany/goany"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/hello", WithAny(func(req *goany.Request, res *goany.Response) error {
		name := req.Path("user.name").String()  // path
		age := req.Get("user").Get("age").Int() // get

		res.Set("message", name).Set("age", age)

		return nil
	}))

	payload := []byte(`{"user":{"name":"Alice", "age":30}}`)
	req := httptest.NewRequest("POST", "/hello", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	require.NotNil(t, rec, "Response recorder should not be nil")
	require.Equal(t, rec.Code, http.StatusOK, "Expected status code 200 OK")

	// validate the response
	var respond = goany.NewRequest(rec.Body.Bytes())
	require.Equal(t, "Alice", respond.Path("message").String(), "Unexpected message in response")
	require.Equal(t, "30", respond.Path("age").String(), "Unexpected age in response")
}
