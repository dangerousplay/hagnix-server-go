package service

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/modelxml"
	"math"
)

var fame = &FameService{}

const ANCESTOR_DESCRIPTION = "FameBonus.AncestorDescription"
const ANCESTOR = "FameBonus.Ancestor"

type FameService struct{}

type PCStats struct {
	Shots                  int
	ShotsThatDamage        int
	SpecialAbilityUses     int
	TilesUncovered         int
	Teleports              int
	PotionsDrunk           int
	MonsterKills           int
	MonsterAssists         int
	GodKills               int
	GodAssists             int
	CubeKills              int
	OryxKills              int
	QuestsCompleted        int
	PirateCavesCompleted   int
	UndeadLairsCompleted   int
	AbyssOfDemonsCompleted int
	SnakePitsCompleted     int
	SpiderDensCompleted    int
	SpriteWorldsCompleted  int
	LevelUpAssists         int
	MinutesActive          int
	TombsCompleted         int
	TrenchesCompleted      int
	JunglesCompleted       int
	ManorsCompleted        int
}

func (service *FameService) GetDeathFame(account *models.Accounts, character *models.Characters) (*modelxml.DeathXML, error) {
	if account == nil || character == nil {
		return nil, errors.New(
			fmt.Sprintf("account or character are nil, account: %p character: %p", account, character))
	}

	var fame = &modelxml.DeathXML{
		BaseFame: character.Fame,
	}

	var bonus = 0

	if character.Charid < 2 {
		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: ANCESTOR_DESCRIPTION,
			Id:          ANCESTOR,
			Value:       fmt.Sprintf("%d", caculateAncestor(character.Fame, &bonus)),
		})
	}
	return fame, nil

	//TODO PCStats Review on C#, make gRPC call...
}

func caculateAncestor(baseFame int, bonus *int) int {
	//(Math.Floor(((baseFame + Math.Floor(bonus)) * 0.1) + 20)
	//Math.Floor(bonus) + ((baseFame + Math.Floor(bonus)) * 0.1) + 20;
	fbaseFame := float64(baseFame)
	fbonus := float64(*bonus)
	result := int(math.Floor(((fbaseFame + math.Floor(fbonus)) * 0.1) + 20))

	*bonus = int(fbonus + (((fbaseFame + fbonus) * 0.1) + 20))

	return result
}

func GetFameService() *FameService {
	return fame
}
