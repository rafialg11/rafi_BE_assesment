package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rafialg11/rafi_BE_assesment/src/config"
	handler "github.com/rafialg11/rafi_BE_assesment/src/handlers"
)

func main() {
	app := fiber.New()
	config.InitDB()

	v1 := app.Group("/api/v1")

	//Initialize Account Handler
	handler.NewAccountHandler(v1)

	//Start the server
	app.Listen(":3000")
}
