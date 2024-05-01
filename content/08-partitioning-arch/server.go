package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/api"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/config"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/repository"
	"github.com/midoon/tiktok-playground/content/08-partitioning-arch/service"
)

func main() {
	dbConnection := config.GetDBOpenConenction()
	userRepository := repository.NewUserRepository(dbConnection)
	userService := service.NewUserService(userRepository)

	app := fiber.New()
	api.NewUserApi(app, userService)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
