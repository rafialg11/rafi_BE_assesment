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
	app.Get("/saldo/:account_number", accountHandler.GetBalance)
}

func (h *AccountHandler) Register(c *fiber.Ctx) error {
	customer := new(entities.Customer)
	if err := c.BodyParser(customer); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "FAILED",
			Data:    nil,
			Error:   "Invalid request body",
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
	transaction := new(entities.TransactionRequest)
	if err := c.BodyParser(transaction); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "FAILED",
			Data:    nil,
			Error:   "Invalid request body",
		})
	}

	account, err := a.accountService.Save(transaction)
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
			"account_number": account.AccountNumber,
			"amount":         account.Amount,
		},
		Error: nil,
	})
}

func (a *AccountHandler) Withdraw(c *fiber.Ctx) error {
	transaction := new(entities.TransactionRequest)
	if err := c.BodyParser(transaction); err != nil {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "FAILED",
			Data:    nil,
			Error:   "Invalid request body",
		})
	}

	account, err := a.accountService.Withdraw(transaction)
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
			"account_number": account.AccountNumber,
			"amount":         account.Amount,
		},
		Error: nil,
	})
}
func (a *AccountHandler) GetBalance(c *fiber.Ctx) error {
	accountNumber := c.Params("account_number")
	if accountNumber == "" {
		return c.JSON(utils.ApiResponse{
			Status:  fiber.StatusBadRequest,
			Message: "FAILED",
			Data:    nil,
			Error:   "account number cannot be empty",
		})
	}

	account, err := a.accountService.GetBalance(&entities.Account{
		AccountNumber: accountNumber,
	})
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
			"account_number": account.AccountNumber,
			"amount":         account.Amount,
		},
		Error: nil,
	})
}
