package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

var AuthLimiter = limiter.New(limiter.Config{
	Max: 7,
	LimitReached: func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusGatewayTimeout, "Too many requests")
	},
})
