package service

import (
	"errors"
	"fmt"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/utils"
	"time"
)

var instance = AccountService{}

type AccountService struct {
}

var randomNames = []string{
	"Darq", "Deyst", "Drac", "Drol",
	"Eango", "Eashy", "Eati", "Eendi", "Ehoni",
	"Gharr", "Iatho", "Iawa", "Idrae",
	"Iri", "Issz", "Itani",
	"Laen", "Lauk", "Lorz",
	"Oalei", "Odaru", "Oeti", "Orothi", "Oshyu",
	"Queq", "Radph", "Rayr", "Ril", "Rilr", "Risrr",
	"Saylt", "Scheev", "Sek", "Serl", "Seus",
	"Tal", "Tiar", "Uoro", "Urake", "Utanu",
	"Vorck", "Vorv", "Yangu", "Yimi", "Zhiar",
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

func (service *AccountService) CreateChest(accounts *models.Accounts) (int64, error) {
	return database.GetDBEngine().Insert(&models.Vaults{Accid: int(accounts.Id), Items: "-1, -1, -1, -1, -1, -1, -1, -1"})
}

func (service *AccountService) CreateChests(accounts *models.Accounts, amount int) (int64, error) {
	var chests []models.Vaults

	if amount <= 0 {
		return 0, errors.New(fmt.Sprintf("Ammount chest is invalid: %d", amount))
	}

	for x := 0; x < amount; x++ {
		chests = append(chests, models.Vaults{Accid: int(accounts.Id), Items: "-1, -1, -1, -1, -1, -1, -1, -1"})
	}

	return database.GetDBEngine().Insert(chests)
}

func (service *AccountService) ContainsAndGetSkin(accounts *models.Accounts, skintype int) (bool, []int, error) {
	skins, err := utils.FromCommaSpaceSeparated(accounts.Ownedskins)
	var contains = false

	for _, v := range skins {
		if v == skintype {
			contains = true
		}
	}

	return contains, skins, err
}

func (service *AccountService) AccountExists(uuid string) (bool, error) {
	return database.GetDBEngine().Where("uuid = ?", uuid).Exist(&models.Accounts{})
}

func (service *AccountService) Register(email string, password string) (*models.Accounts, error) {
	if ex, _ := service.AccountExists(email); ex {
		return nil, errors.New("email already exits")
	}

	session := database.GetDBEngine().NewSession()

	session.Begin()

	defer session.Close()

	index := int(utils.Hashcode(email)) % len(randomNames)

	account := &models.Accounts{
		Uuid:          email,
		Password:      utils.HashStringSHA1(password),
		ObjectId:      "",
		Authtoken:     utils.RandomString(128),
		Name:          randomNames[index],
		Regtime:       time.Now(),
		Rank:          0,
		Vaultcount:    1,
		Maxcharslot:   2,
		Isageverified: 1,
	}

	rows, err := database.GetDBEngine().Insert(account)

	if err != nil || rows < 1 {
		session.Rollback()
		return nil, err
	}

	stats := &models.Stats{
		Accid:        int(account.Id),
		Fame:         1000,
		Totalfame:    1000,
		Credits:      20000,
		Totalcredits: 20000,
	}

	vault := &models.Vaults{
		Accid: int(account.Id),
		Items: "-1, -1, -1, -1, -1, -1, -1, -1",
	}

	rows, err = session.Insert(stats, vault)

	if err != nil || rows < 2 {
		session.Rollback()
		return nil, err
	}

	session.Commit()

	return account, nil
}
func (service *AccountService) NameExists(name string) (bool, error) {
	return database.GetDBEngine().Where("name = ?", name).Exist(&models.Accounts{})
}

func GetAccountService() *AccountService {
	return &instance
}
