package repository

import (
	"github.com/rafialg11/rafi_BE_assesment/src/entities"
	"gorm.io/gorm"
)

type AccountRepository interface {
	SaveCustomer(customer *entities.Customer) (*entities.Customer, error)
	SaveAccount(account *entities.Account) (*entities.Account, error)
	CheckCustomerNIKandPhone(nik string, phone string) (*entities.Customer, error)
	CheckAccountNumber(accountNumber string) (*entities.Account, error)
	PreloadCustomerAccount(customer *entities.Customer) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) SaveCustomer(customer *entities.Customer) (*entities.Customer, error) {
	return customer, r.db.Save(customer).Error
}

func (r *accountRepository) SaveAccount(account *entities.Account) (*entities.Account, error) {
	return account, r.db.Save(account).Error
}

func (r *accountRepository) CheckCustomerNIKandPhone(nik string, phone string) (*entities.Customer, error) {
	var customer entities.Customer
	err := r.db.Where("nik = ? AND phone = ?", nik, phone).First(&customer).Error
	return &customer, err
}

func (r *accountRepository) CheckAccountNumber(accountNumber string) (*entities.Account, error) {
	var account entities.Account
	err := r.db.Where("account_number = ?", accountNumber).First(&account).Error
	return &account, err
}

func (r *accountRepository) PreloadCustomerAccount(customer *entities.Customer) error {
	return r.db.Preload("Account").First(customer, customer.Id).Error
}
