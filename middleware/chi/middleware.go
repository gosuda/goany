package goanychi

import (
	"io"
	"net/http"

	_ "github.com/go-chi/chi/v5/middleware"
	jsoniter "github.com/json-iterator/go"
	"github.com/rabbitprincess/goany/goany"
)

type AnyCtx struct {
	In  goany.Any
	Out map[string]any
}

func (c *AnyCtx) JSON(key string, val any) {
	c.Out[key] = val
}

func WithAny(fn func(*AnyCtx)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := io.ReadAll(r.Body)
		_ = r.Body.Close()

		var parsed any
		_ = jsoniter.Unmarshal(bodyBytes, &parsed)

		ctx := &AnyCtx{
			In:  goany.Wrap(parsed),
			Out: make(map[string]any),
		}

		fn(ctx)

		w.Header().Set("Content-Type", "application/json")
		b, _ := jsoniter.Marshal(ctx.Out)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(b)
	}
}
