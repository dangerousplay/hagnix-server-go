package service

import (
	"errors"
	"fmt"
	"github.com/ivahaev/go-logger"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/modelxml"
	"hagnix-server-go1/routes/utils"
	"k8s.io/apimachinery/pkg/util/rand"
	"strconv"
	"strings"
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

	err := session.Begin()

	if err != nil {
		return nil, err
	}

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
		Lastseen:      time.Now(),
	}

	rows, err := database.GetDBEngine().InsertOne(account)

	if err != nil || rows < 1 {
		session.Rollback()

		if err != nil {
			logger.Warn(err)
		}

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

func (service *AccountService) Lock(uuid string) (int64, error) {
	return database.GetDBEngine().Where("uuid = ?", uuid).Update(&models.Accounts{Accountinuse: 1})
}

func (service *AccountService) Unlock(uuid string) (int64, error) {
	return database.GetDBEngine().Where("uuid = ?", uuid).Update(&models.Accounts{Accountinuse: 0})
}

func (service *AccountService) NextCharId(account *models.Accounts) (int, error) {
	maps, err := database.GetDBEngine().SQL("SELECT IFNULL(MAX(charId), 0) + 1 FROM characters WHERE accId = ?", account.Id).Query()

	if err != nil {
		return 0, err
	}

	number, err := strconv.Atoi(string(maps[0]["IFNULL(MAX(charId), 0) + 1"]))

	if err != nil {
		return 0, err
	}

	return number, err

}

func (service *AccountService) VerifyGenerateAccountXML(uuid string, password string) (*modelxml.AccountXML, *models.Accounts, error) {
	account, err := service.Verify(uuid, password)

	if err != nil || account == nil {
		return nil, nil, err
	}

	xmlt, err := generateAccountXML(account)

	return xmlt, account, err
}

func generateAccountXML(account *models.Accounts) (*modelxml.AccountXML, error) {
	stats := &models.Stats{}

	success, err := database.GetDBEngine().Where("accId = ?", account.Id).Get(stats)

	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errors.New("stats for account not found")
	}

	var verifiedEmail = false

	if account.Verified == 1 {
		verifiedEmail = true
	}

	var admin = false

	if account.Rank > 2 {
		admin = true
	}

	var classes []models.Classstats

	err = database.GetDBEngine().Where("accId = ?", account.Id).Find(&classes)

	if err != nil {
		return nil, err
	}

	var dailyQuest models.Dailyquests

	success, err = database.GetDBEngine().Where("accId = ?", account.Id).Get(&dailyQuest)

	if err != nil {
		return nil, err
	}

	var imageIndex int

	if dailyQuest.Tier <= 0 {
		imageIndex = 0
	} else {
		imageIndex = dailyQuest.Tier - 1
	}

	var goalIndex = imageIndex

	var nameChose = false

	if account.Namechosen != 0 {
		nameChose = true
	}

	goals, err := utils.FromCommaSpaceSeparated(dailyQuest.Goals)

	if err != nil {
		return nil, err
	}

	var guild models.Guilds

	success, err = database.GetDBEngine().Where("id = ?", account.Guild).Get(&guild)

	if err != nil {
		return nil, err
	}

	var vaults []models.Vaults

	err = database.GetDBEngine().Where("accId = ?", account.Id).Find(&vaults)

	if err != nil {
		return nil, err
	}

	classesXML := modelxml.ToClassStatsXML(classes)

	xmlt := modelxml.AccountXML{
		Id:                      account.Id,
		Name:                    account.Name,
		Namechosen:              nameChose,
		Admin:                   admin,
		Verified:                verifiedEmail,
		Credits:                 stats.Credits,
		FortuneTokens:           stats.Fortunetokens,
		NextCharSlotPrice:       NextCharSlotPrice(account),
		BeginnerPackageTimeLeft: 0,
		PetYardType:             account.Petyardtype,
		ArenaTickets:            0,
		IsAgeVerified:           account.Isageverified,

		Stats: modelxml.StatsXML{
			ClassStats:   classesXML,
			TotalFame:    stats.Totalfame,
			Fame:         stats.Fame,
			BestCharFame: getMaxCharFame(classes),
		},

		DailyQuest: modelxml.DailyQuestXML{
			Description: modelxml.DailyQuestDescriptions[dailyQuest.Tier-1],
			Tier:        dailyQuest.Tier,
			Image:       modelxml.ImageTiers[imageIndex],
			Goal:        goals[goalIndex],
		},

		Guild: modelxml.GuildXML{
			Id:   account.Guild,
			Rank: account.Guildrank,
			Fame: account.Guildfame,
			Name: guild.Name,
		},

		Vault: *modelxml.ToVaultXML(vaults),
		Gifts: modelxml.GiftsXML{
			Gifts: account.Gifts,
		},
	}
	return &xmlt, err
}

func (service *AccountService) VerifyGenerateAccountXMLbyId(uuid string) (*modelxml.AccountXML, *models.Accounts, error) {
	var account models.Accounts

	_, err := database.GetDBEngine().Where("id = ?", uuid).Get(&account)

	if err != nil {
		return nil, nil, err
	}

	xmlt, err := generateAccountXML(&account)

	return xmlt, &account, err
}

func (service *AccountService) VerifyGenerateAccountXMLbyUuid(uuid string) (*modelxml.AccountXML, *models.Accounts, error) {
	var account models.Accounts

	_, err := database.GetDBEngine().Where("uuid = ?", uuid).Get(&account)

	if err != nil {
		return nil, nil, err
	}

	xmlt, err := generateAccountXML(&account)

	return xmlt, &account, err
}

func (service *AccountService) GetAvailableClasses(accounts *models.Accounts) ([]modelxml.ClassAvailabilityXML, error) {
	var classes []models.Unlockedclasses
	err := database.GetDBEngine().Cols("class", "available").Where("accId = ?", accounts.Id).Find(&classes)

	if err != nil {
		return nil, err
	}

	if len(classes) < 1 {
		session := database.GetDBEngine().NewSession()
		err = session.Begin()

		defer session.Close()

		if err != nil {
			session.Rollback()
			return nil, err
		}

		for _, v := range modelxml.Classes {
			session.Insert(models.Unlockedclasses{Class: v.Class, Available: v.Restricted})
		}
		err = session.Commit()

		if err != nil {
			session.Rollback()
			return nil, err
		}

		return modelxml.Classes, err
	} else {
		var xmls []modelxml.ClassAvailabilityXML
		for _, v := range classes {
			xmls = append(xmls, modelxml.ClassAvailabilityXML{Class: v.Class, Restricted: v.Available})
		}
		return xmls, nil
	}
}

func (service *AccountService) GenerateAccountXML(uuid string, password string) (*modelxml.AccountXML, error) {
	account, _, err := service.VerifyGenerateAccountXML(uuid, password)
	return account, err
}

func (service *AccountService) GetCharById(accountId int64, charId int) (*modelxml.CharXML, error) {
	var chars models.Characters
	success, err := database.GetDBEngine().Where("id = ?", charId).Get(&chars)

	if err != nil {
		return nil, err
	}

	if !success {
		return nil, errors.New("no character found by id")
	}
	return &toCharXML(accountId, chars)[0], err
}

func (service *AccountService) GetCharsXML(account *models.Accounts) ([]modelxml.CharXML, error) {
	var chars []models.Characters
	err := database.GetDBEngine().Where("accId = ? AND dead = FALSE", account.Id).Find(&chars)

	if err != nil {
		return nil, err
	}

	var charsXML = toCharXML(account.Id, chars...)

	return charsXML, err
}

func (service *AccountService) GetRandomName() string {
	rands := rand.IntnRange(0, len(randomNames))
	return randomNames[rands]
}

func toCharXML(accountId int64, chars ...models.Characters) []modelxml.CharXML {
	var charsXML []modelxml.CharXML

	for _, v := range chars {
		stats, err := utils.FromCommaSpaceSeparated(v.Stats)

		if err != nil {
			logger.Warn("Can't load char: " + err.Error())
			continue
		}

		var pet models.Pets

		success, err := database.GetDBEngine().Where("accId = ? AND petId = ?", accountId, v.Petid).Get(&pet)

		if err != nil {
			logger.Warn("Can't load pet: " + err.Error())
			continue
		}

		var dead = false

		if v.Dead != 0 {
			dead = true
		}

		charXML := modelxml.CharXML{
			Id:         v.Charid,
			ObjectType: v.Chartype,
			//CharacterId: rdr.GetInt32("charId"),
			Level:            v.Level,
			Exp:              v.Exp,
			CurrentFame:      v.Fame,
			Equipment:        v.Items,
			MaxHitPoints:     stats[0],
			HitPoints:        v.Hp,
			MaxMagicPoints:   stats[1],
			MagicPoints:      v.Mp,
			Attack:           stats[2],
			Defense:          stats[3],
			Speed:            stats[4],
			Dexterity:        stats[7],
			HpRegen:          stats[5],
			MpRegen:          stats[6],
			HealthStackCount: v.Hppotions,
			MagicStackCount:  v.Mppotions,
			HasBackpack:      v.Hasbackpack,
			Tex1:             v.Tex1,
			Tex2:             v.Tex2,
			Dead:             dead,
			PCStats:          v.Famestats,
			Skin:             v.Skin,
		}

		if success {
			charXML.Pet = &modelxml.PetItemXML{
				SkinName:        pet.Skinname,
				Type:            pet.Objtype,
				InstanceId:      pet.Petid,
				Skin:            pet.Skin,
				Rarity:          pet.Rarity,
				MaxAbilityPower: pet.Maxlevel,
				Abilities:       modelxml.AbilityWrapper{Abilities: toAbilitiesXML(pet)},
			}
		}

		charsXML = append(charsXML, charXML)
	}
	return charsXML
}

func toAbilitiesXML(abilities models.Pets) []modelxml.AbilityItemXML {
	lenght := len(strings.Split(abilities.Levels, ","))

	var abilityXML []modelxml.AbilityItemXML

	xps, _ := utils.FromCommaSpaceSeparated(abilities.Xp)
	levels, _ := utils.FromCommaSpaceSeparated(abilities.Levels)
	abily, _ := utils.FromCommaSpaceSeparated(abilities.Abilities)

	for i := 0; i < lenght; i++ {
		abilityXML = append(abilityXML, modelxml.AbilityItemXML{
			Points: xps[i],
			Power:  levels[i],
			Type:   abily[i],
		})
	}

	return abilityXML
}

func getMaxCharFame(classes []models.Classstats) int {
	var best int

	for _, v := range classes {
		if v.Bestfame >= best {
			best = v.Bestfame
		}
	}
	return best
}

func NextCharSlotPrice(account *models.Accounts) int {
	return NextCharSlotPriceByChars(account.Maxcharslot)
}

func NextCharSlotPriceByChars(chars int) int {
	var price int

	if chars == 1 {
		price = 600
	} else if chars == 2 {
		price = 800
	} else {
		price = 1000
	}

	return price
}

func GetAccountService() *AccountService {
	return &instance
}
