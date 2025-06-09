package goanyhttp

import (
	"io"
	"net/http"

	"github.com/rabbitprincess/goany/goany"
)

func WithAnyNetHTTP(fn goany.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		req := goany.NewRequest(bodyBytes)
		res := goany.NewResponse()

		if err := fn(req, res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		b, _ := res.MarshalJSON()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	}
}
