package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gochat/api/v1/services"
	zaplogger "gochat/config/logger"
	"gochat/types/dto/response"
)

var loggerStruct = &zaplogger.Logger{}
var logger = loggerStruct.New("User Controller")

func Login(ctx *fiber.Ctx) error {

	logger.Info("Login Consumed", string(ctx.Body()))

	userService := services.UserService{}
	token, err := userService.LoginUser(ctx.Body())

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.Auth{Status: "Success", Token: token})
}

func Register(ctx *fiber.Ctx) error {

	logger.Info("Registered User", string(ctx.Body()))

	userService := services.UserService{}
	err := userService.CreateUser(ctx.Body())
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(response.BaseResponse{Status: "Success", Message: "User Created"})
}
