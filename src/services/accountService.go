package services

import (
	"errors"

	"github.com/rafialg11/rafi_BE_assesment/src/config"
	"github.com/rafialg11/rafi_BE_assesment/src/entities"
	"github.com/rafialg11/rafi_BE_assesment/src/helpers"
	"github.com/rafialg11/rafi_BE_assesment/src/repository"
)

type AccountService interface {
	Register(account *entities.Customer) (*entities.Customer, error)
	Save(account *entities.TransactionRequest) (*entities.Account, error)
	Withdraw(account *entities.TransactionRequest) (*entities.Account, error)
	GetBalance(account *entities.Account) (*entities.Account, error)
}

type accountService struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return &accountService{repository: repository}
}

func (s *accountService) Register(customer *entities.Customer) (*entities.Customer, error) {
	// Input Validation
	if customer.Name == "" || customer.Phone == "" || customer.NIK == "" {
		config.LogError("Name, phone, and nik cannot be empty", map[string]interface{}{
			"name":  customer.Name,
			"phone": customer.Phone,
			"nik":   customer.NIK,
		})
		return nil, errors.New("name, phone, and nik cannot be empty")
	}

	// Check if NIK and Phone already exists
	isCustomerExists, err := s.repository.CheckCustomerNIKandPhone(customer.NIK, customer.Phone)
	if err != nil {
		config.LogError("Failed to check customer", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to check customer")
	}
	if isCustomerExists {
		config.LogError("Customer already exists", map[string]interface{}{
			"nik":   customer.NIK,
			"phone": customer.Phone,
		})
		return nil, errors.New("customer already exists")
	}

	// Save Customer
	_, err = s.repository.SaveCustomer(customer)
	if err != nil {
		config.LogError("Failed to save customer", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to save customer")
	}

	// Generate Account Number using helper
	accountNumber := helpers.GetFormattedAccountNumber(customer.Id)

	// Check Customer Account Number
	isAccountNumberExists, err := s.repository.CheckAccountNumber(accountNumber)
	if err != nil {
		config.LogError("Failed to check account number", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to check account number")
	}
	if isAccountNumberExists {
		config.LogError("Account number already exists", map[string]interface{}{
			"account_number": accountNumber,
		})
		return nil, errors.New("account number already exists")
	}

	// Create Account
	account := &entities.Account{
		AccountNumber: accountNumber,
		Amount:        0,
		CustomerId:    customer.Id,
	}

	// Save Account
	_, err = s.repository.SaveAccount(account)
	if err != nil {
		config.LogError("Failed to save account", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to save account")
	}

	// Preload Customer and Account
	err = s.repository.PreloadCustomerAccount(customer)
	if err != nil {
		config.LogError("Failed to preload customer and account", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to preload customer and account")
	}
	config.LogInfo("Customer registered successfully", map[string]interface{}{
		"customer": customer,
	})
	return customer, err
}

func (s *accountService) Save(transaction *entities.TransactionRequest) (*entities.Account, error) {
	// Input Validation
	if transaction.AccountNumber == "" {
		config.LogError("Account number cannot be empty", map[string]interface{}{
			"account_number": transaction.AccountNumber,
		})
		return nil, errors.New("account number cannot be empty")
	}

	// Amount Validation
	if transaction.Amount <= 0 {
		config.LogError("Amount must be greater than 0", map[string]interface{}{
			"amount": transaction.Amount,
		})
		return nil, errors.New("amount must be greater than 0")
	}

	// Check Account Balance
	checkAccount, err := s.repository.CheckAccountBalance(transaction.AccountNumber)
	if checkAccount == nil {
		config.LogError("Account not found", map[string]interface{}{
			"account_number": transaction.AccountNumber,
		})
		return nil, errors.New("account not found")
	}
	if err != nil {
		config.LogError("Failed to check account balance", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to check account balance")
	}

	// Save Transaction
	trx := &entities.Transaction{
		Amount:          transaction.Amount,
		TransactionType: "Save",
		CustomerId:      checkAccount.CustomerId,
	}

	// Save Transaction
	_, err = s.repository.CreateTransaction(trx)
	if err != nil {
		config.LogError("Failed to save transaction", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to save transaction")
	}

	// Update Account Balance
	checkAccount.Amount += transaction.Amount
	account, err := s.repository.UpdateBalance(checkAccount)
	if err != nil {
		config.LogError("Failed to update account balance", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to update account balance")
	}

	config.LogInfo("Transaction successful", map[string]interface{}{
		"account_number": transaction.AccountNumber,
		"amount":         transaction.Amount,
	})
	return account, err
}

func (s *accountService) Withdraw(transaction *entities.TransactionRequest) (*entities.Account, error) {
	if transaction.AccountNumber == "" {
		config.LogError("Account number cannot be empty", map[string]interface{}{
			"account_number": transaction.AccountNumber,
		})
		return nil, errors.New("account number cannot be empty")
	}

	// Amount Validation
	if transaction.Amount <= 0 {
		config.LogError("Amount must be greater than 0", map[string]interface{}{
			"amount": transaction.Amount,
		})
		return nil, errors.New("amount must be greater than 0")
	}

	// Check Account Balance
	checkAccount, err := s.repository.CheckAccountBalance(transaction.AccountNumber)
	if checkAccount == nil {
		config.LogError("Account not found", map[string]interface{}{
			"account_number": transaction.AccountNumber,
		})
		return nil, errors.New("account not found")
	}
	if err != nil {
		config.LogError("Failed to check account balance", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to check account balance")
	}

	if checkAccount.Amount < transaction.Amount {
		config.LogError("Insufficient balance", map[string]interface{}{
			"account_number": checkAccount.AccountNumber,
			"amount":         checkAccount.Amount,
		})
		return nil, errors.New("insufficient balance")
	}

	// Save Transaction
	trx := &entities.Transaction{
		Amount:          transaction.Amount,
		TransactionType: "Withdraw",
		CustomerId:      checkAccount.CustomerId,
	}

	// Save Transaction
	_, err = s.repository.CreateTransaction(trx)
	if err != nil {
		config.LogError("Failed to save transaction", map[string]interface{}{
			"error": err.Error()})
		return nil, errors.New("failed to save transaction")
	}

	// Update Account Balance
	checkAccount.Amount -= transaction.Amount
	account, err := s.repository.UpdateBalance(checkAccount)
	if err != nil {
		config.LogError("Failed to update account balance", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to update account balance")
	}

	config.LogInfo("Transaction successful", map[string]interface{}{
		"account_number": account.AccountNumber,
		"amount":         account.Amount,
	})
	return account, err
}

func (s *accountService) GetBalance(account *entities.Account) (*entities.Account, error) {
	if account.AccountNumber == "" {
		config.LogError("Account number cannot be empty", map[string]interface{}{
			"account_number": account.AccountNumber,
		})
		return nil, errors.New("account number cannot be empty")
	}

	// Check Account Balance
	checkAccount, err := s.repository.CheckAccountBalance(account.AccountNumber)
	if checkAccount == nil {
		config.LogError("Account not found", map[string]interface{}{
			"account_number": account.AccountNumber,
		})
		return nil, errors.New("account not found")
	}
	if err != nil {
		config.LogError("Failed to check account balance", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, errors.New("failed to check account balance")
	}

	// Log transaksi berhasil
	config.LogInfo("Transaction successful", map[string]interface{}{
		"account_number": checkAccount.AccountNumber,
		"amount":         checkAccount.Amount,
	})

	return checkAccount, nil
}
