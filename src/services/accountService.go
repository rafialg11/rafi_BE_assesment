package services

import (
	"errors"
	"log"

	"github.com/rafialg11/rafi_BE_assesment/src/entities"
	"github.com/rafialg11/rafi_BE_assesment/src/helpers"
	"github.com/rafialg11/rafi_BE_assesment/src/repository"
	"gorm.io/gorm"
)

type AccountService interface {
	Register(account *entities.Customer) (*entities.Customer, error)
	Save(account *entities.Account) (*entities.Account, error)
	Withdraw(account *entities.Account) (*entities.Account, error)
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
	_, err := s.repository.CheckCustomerNIKandPhone(customer.NIK, customer.Phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // Jika error BUKAN karena data tidak ditemukan, return error
		return nil, errors.New("failed to check NIK and phone")
	}
	if err == nil { // Jika tidak ada error (berarti data ditemukan), maka return error karena duplikasi
		return nil, errors.New("nik and phone already exists")
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
	_, err = s.repository.CheckAccountNumber(accountNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to check account number")
	}
	if err == nil {
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

func (s *accountService) Save(account *entities.Account) (*entities.Account, error) {
	return account, nil
}

func (s *accountService) Withdraw(account *entities.Account) (*entities.Account, error) {
	return account, nil
}

func (s *accountService) GetBalance(account *entities.Account) (*entities.Account, error) {
	return account, nil
}
