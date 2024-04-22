package routes

import (
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/controllers"
)

func UserRouter(r fiber.Router) {
	userRouter := r.Group("/user/")

	userRouter.Get("/login/", controllers.Login)
	userRouter.Post("/register/", controllers.Register)
}
