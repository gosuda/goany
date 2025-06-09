package goanychi

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rabbitprincess/goany/goany"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	r := chi.NewRouter()
	r.Post("/hello", WithAny(func(ctx *AnyCtx) {
		// parse req
		message := "Hello my name is " + ctx.In.Path("user.name").String()
		age := ctx.In.Path("user.age").Int()

		// res json
		ctx.JSON("message", message)
		ctx.JSON("age", age)
	}))

	payload := []byte(`{"user":{"name":"Alice", "age":30}}`)
	req := httptest.NewRequest("POST", "/hello", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)
	require.NotNil(t, rec, "Response recorder should not be nil")
	require.Equal(t, rec.Code, http.StatusOK, "Expected status code 200 OK")

	// validate the response
	fmt.Println("Response:", rec.Body.String())
	var respond = goany.Wrap(rec.Body.Bytes())
	require.Equal(t, "Hello my name is Alice", respond.Path("message").String(), "Unexpected message in response")
	require.Equal(t, "30", respond.Path("age").String(), "Unexpected age in response")
}
