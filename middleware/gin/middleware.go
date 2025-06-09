package goanygin

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rabbitprincess/goany/goany"
)

func WithAny(fn goany.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
			return
		}

		req := goany.NewRequest(bodyBytes)
		res := goany.NewResponse()

		if err := fn(req, res); err != nil {
			c.JSON(res.HTTPStatus(err), gin.H{"error": err.Error()})
			return
		}

		c.JSON(res.HTTPStatus(nil), res)
	}
}
