package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"io"
	"log"
	"os"
)

func LogToFile(app *fiber.App) {
	// Logging - Method and Status
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output: file,
	}))
}

func LogToFileAndConsole(app *fiber.App) {
	// Logging - Method and Status
	file, err := os.OpenFile("logfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		Output: io.MultiWriter(file, os.Stdout),
	}))
}
