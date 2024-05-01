package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/dto"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/helper"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/model"
)

type userApi struct {
	userService model.UserService
}

func NewUserApi(app *fiber.App, userService model.UserService) {
	handler := userApi{
		userService: userService,
	}

	app.Post("/api/register", handler.register)
	app.Post("/api/login", handler.login)
}

func (u *userApi) login(fctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := fctx.BodyParser(&req); err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	userData, err := u.userService.Login(fctx.Context(), &req)
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	res := dto.LoginResponse{
		Status:  true,
		Message: "Success login",
		Data:    userData,
	}

	return fctx.Status(200).JSON(res)
}

func (u *userApi) register(fctx *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := fctx.BodyParser(&req); err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	err := u.userService.Register(fctx.Context(), &req)
	if err != nil {
		return fctx.Status(helper.HttpStatusErr(err)).JSON(dto.BasicResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	res := dto.BasicResponse{
		Status:  true,
		Message: "Success registration",
	}
	return fctx.Status(200).JSON(res)
}
