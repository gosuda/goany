package goanychi

import (
	"io"
	"net/http"

	_ "github.com/go-chi/chi/v5/middleware"
	jsoniter "github.com/json-iterator/go"
	"github.com/rabbitprincess/goany/goany"
)

func WithAny(fn func(req goany.Request, res goany.Response)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := io.ReadAll(r.Body)
		_ = r.Body.Close()

		In := goany.NewRequest(bodyBytes)
		Out := goany.NewResponse()

		fn(In, Out)

		w.Header().Set("Content-Type", "application/json")
		b, _ := jsoniter.Marshal(Out)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	}
}
