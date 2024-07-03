package cmd

import (
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/routes"
	"gochat/api/v1/ws"
)

func HandleRoutes(app *fiber.App) {
	// v1 Router

	v1 := app.Group("/v1/")

	// User Router
	routes.UserRouter(v1)

	// SocketIO
	ws.SocketIORouter(v1)
}
