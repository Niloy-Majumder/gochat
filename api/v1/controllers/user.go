package controllers

import (
	"github.com/gofiber/fiber/v2"
	zaplogger "gochat/config/logger"
)

var loggerStruct = &zaplogger.Logger{}
var logger = loggerStruct.New("User Controller")

func Login(ctx *fiber.Ctx) error {

	logger.Info("Login Consumed", string(ctx.Body()))

	return ctx.SendString("User LoggedIn")
}

func Register(ctx *fiber.Ctx) error {

	logger.Info("Registered Consumed", string(ctx.Body()))

	return ctx.Status(200).SendString("User Registered")
}
