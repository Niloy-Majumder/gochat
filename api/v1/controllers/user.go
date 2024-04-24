package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/services"
	zaplogger "gochat/config/logger"
)

var loggerStruct = &zaplogger.Logger{}
var logger = loggerStruct.New("User Controller")

func Login(ctx *fiber.Ctx) error {

	logger.Info("Login Consumed", string(ctx.Body()))

	return ctx.SendString("User LoggedIn")
}

func Register(ctx *fiber.Ctx) error {

	logger.Info("Registered User", string(ctx.Body()))

	userService := services.UserService{}
	err := userService.CreateUser(ctx.Body())
	if err != nil {
		return err
	}

	return ctx.Status(200).SendString("User Registered")
}
