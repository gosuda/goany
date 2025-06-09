package goanyfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rabbitprincess/goany/goany"
)

func WithAny(fn goany.HandlerFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Body()
		req := goany.NewRequest(body)
		res := goany.NewResponse()

		if err := fn(req, res); err != nil {
			return c.Status(res.HTTPStatus(err)).SendString(err.Error())
		}

		b, _ := res.MarshalJSON()
		return c.Status(res.HTTPStatus(nil)).Send(b)
	}
}
