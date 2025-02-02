package repository

import (
	"github.com/rafialg11/rafi_BE_assesment/src/entities"
	"gorm.io/gorm"
)

type AccountRepository interface {
	SaveCustomer(customer *entities.Customer) (*entities.Customer, error)
	SaveAccount(account *entities.Account) (*entities.Account, error)
	CheckCustomerNIKandPhone(nik string, phone string) (bool, error)
	CheckAccountNumber(accountNumber string) (bool, error)
	PreloadCustomerAccount(customer *entities.Customer) error
	CheckAccountBalance(accountNumber string) (*entities.Account, error)

	CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error)
	UpdateBalance(account *entities.Account) (*entities.Account, error)
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

func (r *accountRepository) CheckCustomerNIKandPhone(nik string, phone string) (bool, error) {
	var count int64
	err := r.db.Model(&entities.Customer{}).Where("nik = ? AND phone = ?", nik, phone).Count(&count).Error
	return count > 0, err
}

func (r *accountRepository) CheckAccountNumber(accountNumber string) (bool, error) {
	var count int64
	err := r.db.Model(&entities.Account{}).Where("account_number = ?", accountNumber).Count(&count).Error
	return count > 0, err
}

func (r *accountRepository) PreloadCustomerAccount(customer *entities.Customer) error {
	return r.db.Preload("Account").First(customer, customer.Id).Error
}

func (r *accountRepository) CheckAccountBalance(accountNumber string) (*entities.Account, error) {
	var account entities.Account
	err := r.db.Where("account_number = ?", accountNumber).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) CreateTransaction(transaction *entities.Transaction) (*entities.Transaction, error) {
	return transaction, r.db.Save(transaction).Error
}

func (r *accountRepository) UpdateBalance(account *entities.Account) (*entities.Account, error) {
	updateBalance := &entities.Account{
		Amount: account.Amount,
	}
	if err := r.db.Model(&account).
		Where("account_number = ?", account.AccountNumber).
		Updates(updateBalance).
		First(&account, account.Id).Error; err != nil {
		return nil, err
	} else {
		return account, nil
	}
}
