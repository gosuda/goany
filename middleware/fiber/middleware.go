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
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		b, err := res.MarshalJSON()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("failed to encode response")
		}

		return c.Status(fiber.StatusOK).Send(b)
	}
}
