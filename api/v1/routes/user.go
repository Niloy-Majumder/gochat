package routes

import (
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/controllers"
)

func UserRouter(r fiber.Router) {
	userRouter := r.Group("/user/")

	authRouter(userRouter)
}

func authRouter(r fiber.Router) {
	userRouter := r.Group("/auth/")

	userRouter.Post("/register/", controllers.Register)
	userRouter.Get("/login/", controllers.Login)
}
