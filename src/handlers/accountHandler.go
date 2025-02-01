package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafialg11/rafi_BE_assesment/src/utils"
)

type AccountHandler struct {
}

func NewAccountHandler(app fiber.Router) {
	accountHandler := AccountHandler{}
	app.Post("/daftar", accountHandler.Register)
	app.Post("/tabung", accountHandler.Register)
	app.Post("/tarik", accountHandler.Register)
	app.Get("/saldo/:id", accountHandler.Register)
}

func (a *AccountHandler) Register(c *fiber.Ctx) error {
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "SUCCESS",
		Data:    "this is Register",
		Error:   nil,
	})
}

func (a *AccountHandler) Save(c *fiber.Ctx) error {
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "SUCCESS",
		Data:    "this is Save",
		Error:   nil,
	})
}
func (a *AccountHandler) Withdraw(c *fiber.Ctx) error {
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "SUCCESS",
		Data:    "this is Withdraw",
		Error:   nil,
	})
}
func (a *AccountHandler) GetBalance(c *fiber.Ctx) error {
	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "SUCCESS",
		Data:    "this is GetBalance",
		Error:   nil,
	})
}
