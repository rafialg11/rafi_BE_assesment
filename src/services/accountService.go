package services

import (
	"errors"
	"log"

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
		return nil, errors.New("name, phone, and nik cannot be empty")
	}

	// Check if NIK and Phone already exists
	isCustomerExists, err := s.repository.CheckCustomerNIKandPhone(customer.NIK, customer.Phone)
	if err != nil {
		return nil, errors.New("failed to check customer")
	}
	if isCustomerExists {
		return nil, errors.New("customer already exists")
	}

	// Save Customer
	log.Println("Saving customer:", customer)
	_, err = s.repository.SaveCustomer(customer)
	if err != nil {
		return nil, errors.New("failed to save customer")
	}

	// Generate Account Number using helper
	accountNumber := helpers.GetFormattedAccountNumber(customer.Id)

	// Check Customer Account Number
	isAccountNumberExists, err := s.repository.CheckAccountNumber(accountNumber)
	if err != nil {
		return nil, errors.New("failed to check account number")
	}
	if isAccountNumberExists {
		return nil, errors.New("account number already exists")
	}

	// Create Account
	account := &entities.Account{
		AccountNumber: accountNumber,
		Amount:        0,
		CustomerId:    customer.Id,
	}

	// Save Account
	log.Println("Saving account for customer:", customer.Id)
	_, err = s.repository.SaveAccount(account)
	if err != nil {
		return nil, errors.New("failed to save account")
	}

	// Preload Customer and Account
	err = s.repository.PreloadCustomerAccount(customer)
	if err != nil {
		return nil, errors.New("failed to preload customer and account")
	}

	return customer, err
}

func (s *accountService) Save(transaction *entities.TransactionRequest) (*entities.Account, error) {
	// Input Validation
	if transaction.AccountNumber == "" {
		return nil, errors.New("account number cannot be empty")
	}

	// Amount Validation
	if transaction.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	// Check Account Balance
	checkAccount, err := s.repository.CheckAccountBalance(transaction.AccountNumber)
	if checkAccount == nil {
		return nil, errors.New("account not found")
	}
	if err != nil {
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
		return nil, errors.New("failed to save transaction")
	}

	// Update Account Balance
	checkAccount.Amount += transaction.Amount
	account, err := s.repository.UpdateBalance(checkAccount)
	if err != nil {
		return nil, errors.New("failed to update account balance")
	}

	return account, err
}

func (s *accountService) Withdraw(transaction *entities.TransactionRequest) (*entities.Account, error) {
	if transaction.AccountNumber == "" {
		return nil, errors.New("account number cannot be empty")
	}

	// Amount Validation
	if transaction.Amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	// Check Account Balance
	checkAccount, err := s.repository.CheckAccountBalance(transaction.AccountNumber)
	if checkAccount == nil {
		return nil, errors.New("account not found")
	}
	if err != nil {
		return nil, errors.New("failed to check account balance")
	}

	if checkAccount.Amount < transaction.Amount {
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
		return nil, errors.New("failed to save transaction")
	}

	// Update Account Balance
	checkAccount.Amount -= transaction.Amount
	account, err := s.repository.UpdateBalance(checkAccount)
	if err != nil {
		return nil, errors.New("failed to update account balance")
	}

	return account, err
}

func (s *accountService) GetBalance(account *entities.Account) (*entities.Account, error) {
	if account.AccountNumber == "" {
		return nil, errors.New("account number cannot be empty")
	}

	// Check Account Balance
	checkAccount, err := s.repository.CheckAccountBalance(account.AccountNumber)
	if checkAccount == nil {
		return nil, errors.New("account not found")
	}
	if err != nil {
		return nil, errors.New("failed to check account balance")
	}

	return checkAccount, nil
}
