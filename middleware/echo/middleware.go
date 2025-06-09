package goanyecho

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rabbitprincess/goany/goany"
)

func WithAny(fn func(req goany.Request, res goany.Response) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		bodyBytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to read request body"})
		}
		defer c.Request().Body.Close()

		req := goany.NewRequest(bodyBytes)
		res := goany.NewResponse()

		if err := fn(req, res); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, res)
	}
}
