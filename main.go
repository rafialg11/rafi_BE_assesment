package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rafialg11/rafi_BE_assesment/src/config"
	handler "github.com/rafialg11/rafi_BE_assesment/src/handlers"
	"github.com/rafialg11/rafi_BE_assesment/src/repository"
	"github.com/rafialg11/rafi_BE_assesment/src/services"
)

func main() {
	app := fiber.New()
	config.InitDB()
	config.InitLogger()

	v1 := app.Group("/api/v1")

	//Initialize Account Handler
	accountRepo := repository.NewAccountRepository(config.Database)
	accountService := services.NewAccountService(accountRepo)
	handler.NewAccountHandler(v1, accountService)

	//Start the server
	app.Listen(":3000")
}
