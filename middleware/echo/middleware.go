package goanyecho

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rabbitprincess/goany/goany"
)

func WithAny(fn goany.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bodyBytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to read request body"})
		}
		defer c.Request().Body.Close()

		req := goany.NewRequest(bodyBytes)
		res := goany.NewResponse()

		if err := fn(req, res); err != nil {
			return c.JSON(res.HTTPStatus(err), map[string]string{"error": err.Error()})
		}

		return c.JSON(res.HTTPStatus(nil), res)
	}
}
