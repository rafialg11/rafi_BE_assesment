package services

import "github.com/rafialg11/rafi_BE_assesment/src/entities"

type AccountService interface {
	Register(account *entities.Account) (*entities.Account, error)
	Save(account *entities.Account) (*entities.Account, error)
	Withdraw(account *entities.Account) (*entities.Account, error)
	GetBalance(account *entities.Account) (*entities.Account, error)
}

type accountService struct {
}

func NewAccountService() AccountService {
	return &accountService{}
}

func (a *accountService) Register(account *entities.Account) (*entities.Account, error) {
	return account, nil
}

func (a *accountService) Save(account *entities.Account) (*entities.Account, error) {
	return account, nil
}

func (a *accountService) Withdraw(account *entities.Account) (*entities.Account, error) {
	return account, nil
}

func (a *accountService) GetBalance(account *entities.Account) (*entities.Account, error) {
	return account, nil
}
