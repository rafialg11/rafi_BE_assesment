package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafialg11/rafi_BE_assesment/src/entities"
	"github.com/rafialg11/rafi_BE_assesment/src/services"
	"github.com/rafialg11/rafi_BE_assesment/src/utils"
)

type AccountHandler struct {
	accountService services.AccountService
}

func NewAccountHandler(app fiber.Router, accountService services.AccountService) {
	accountHandler := AccountHandler{accountService: accountService}
	app.Post("/daftar", accountHandler.Register)
	app.Post("/tabung", accountHandler.Save)
	app.Post("/tarik", accountHandler.Withdraw)
	app.Get("/saldo/:id", accountHandler.GetBalance)
}

func (h *AccountHandler) Register(c *fiber.Ctx) error {
	customer := new(entities.Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "FAILED",
			Data:    nil,
			Error:   "All fields (name, phone, NIK) are required",
		})
	}

	customer, err := h.accountService.Register(customer)
	if err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "FAILED",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(utils.ApiResponse{
		Status:  fiber.StatusOK,
		Message: "SUCCESS",
		Data: map[string]interface{}{
			"account_number": customer.Account.AccountNumber,
		},
		Error: nil,
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
