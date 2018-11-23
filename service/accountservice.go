package service

import (
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
)

var instance = AccountService{}

type AccountService struct {
}

func (service *AccountService) Verify(uuid string, password string) (*models.Accounts, error) {
	var account models.Accounts

	sucess, err := database.GetDBEngine().Where("uuid = ? AND password = SHA1(?)", uuid, password).Get(&account)

	if sucess {
		return &account, err
	} else {
		return nil, err
	}
}

func (service *AccountService) VerifyOnly(uuid string, password string) (bool, error) {
	var account models.Accounts

	return database.GetDBEngine().Where("uuid = ? AND password = SHA1(?)", uuid, password).Exist(&account)
}

func GetAccountService() *AccountService {
	return &instance
}
