package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/rafialg11/rafi_BE_assesment/src/config"
	handler "github.com/rafialg11/rafi_BE_assesment/src/handlers"
	"github.com/rafialg11/rafi_BE_assesment/src/repository"
	"github.com/rafialg11/rafi_BE_assesment/src/services"
)

func main() {
	host := flag.String("host", "localhost", "Host for the server")
	port := flag.String("port", "3000", "Port for the server")

	flag.Parse()

	fmt.Printf("Starting server on %s:%s\n", *host, *port)

	config.InitDB()
	config.InitLogger()

	app := fiber.New()

	v1 := app.Group("/api/v1")

	accountRepo := repository.NewAccountRepository(config.Database)
	accountService := services.NewAccountService(accountRepo)
	handler.NewAccountHandler(v1, accountService)

	err := app.Listen(fmt.Sprintf("%s:%s", *host, *port))
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
