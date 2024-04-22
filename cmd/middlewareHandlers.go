package cmd

import (
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/middlewares"
)

func HandleMiddlewares(isProd bool, app *fiber.App) {
	if isProd {
		middlewares.LogToFile(app)
	} else {
		middlewares.LogToFileAndConsole(app)
	}
}
