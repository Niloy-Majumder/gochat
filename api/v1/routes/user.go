package routes

import (
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/controllers"
	"gochat/api/v1/middlewares"
)

func UserRouter(r fiber.Router) {
	userRouter := r.Group("/user/")

	authRouter(userRouter)
}

func authRouter(r fiber.Router) {
	userRouter := r.Group("/auth/")

	userRouter.Post("/register/", middlewares.AuthLimiter, controllers.Register)

	userRouter.Post("/login/", middlewares.AuthLimiter, controllers.Login)
}
