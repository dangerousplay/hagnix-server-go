package service

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/modelxml"
	"hagnix-server-go1/routes/utils"
	"io"
	"math"
)

var fame = &FameService{}

const (
	ANCESTOR_DESCRIPTION = "FameBonus.AncestorDescription"
	ANCESTOR             = "FameBonus.Ancestor"

	PACIFIST             = "FameBonus.Pacifist"
	PACIFIST_DESCRIPTION = "FameBonus.PacifistDescription"

	THIRSTY             = "FameBonus.Thirsty"
	THIRSTY_DESCRIPTION = "FameBonus.ThirstyDescription"

	MUNDANE             = "FameBonus.Mundane"
	MUNDANE_DESCRIPTION = "FameBonus.MundaneDescription"

	BOOTS_ON_GROUND             = "FameBonus.BootsOnGround"
	BOOTS_ON_GROUND_DESCRIPTION = "FameBonus.BootsOnGroundDescription"

	TUNNEL_RAT             = "FameBonus.TunnelRat"
	TUNNEL_RAT_DESCRIPTION = "FameBonus.TunnelRatDescription"

	GOD_ENEMY              = "FameBonus.GodEnemy"
	GOD_ENEMNY_DESCRIPTION = "FameBonus.GodEnemyDescription"

	GOD_SLAYER             = "FameBonus.GodSlayer"
	GOD_SLAYER_DESCRIPTION = "FameBonus.GodSlayerDescription"

	ORYX_SLAYER             = "FameBonus.OryxSlayer"
	ORYX_SLAYER_DESCRIPTION = "FameBonus.OryxSlayer"

	ACCURATE             = "FameBonus.Accurate"
	ACCURATE_DESCRIPTION = "FameBonus.AccurateDescription"

	SHARPSHOOTER             = "FameBonus.SharpShooter"
	SHARPSHOOTER_DESCRIPTION = "FameBonus.SharpShooterDescription"

	SNIPER             = "FameBonus.Sniper"
	SNIPER_DESCRIPTION = "FameBonus.SniperDescription"

	EXPLORER             = "FameBonus.Explorer"
	EXPLORER_DESCRIPTION = "FameBonus.ExplorerDescription"

	CARTOGRAPHER             = "FameBonus.Cartographer"
	CARTOGRAPHER_DESCRIPTION = "FameBonus.CartographerDescription"

	TEAMPLAYER             = "FameBonus.TeamPlayer"
	TEAMPLAYER_DESCRIPTION = "FameBonus.TeamPlayerDescription"

	LEADEROFMAN             = "FameBonus.LeaderOfMen"
	LEADEROFMAN_DESCRIPTION = "FameBonus.LeaderOfMenDescription"

	DOER_OF_DEEDS             = "FameBonus.DoerOfDeeds"
	DOER_OF_DEEDS_DESCRIPTION = "FameBonus.DoerOfDeedsDescription"

	CUBEFRIEND             = "FameBonus.CubeFriend"
	CUBEFRIEND_DESCRIPTION = "FameBonus.CubeFriendDescription"

	WELLEQUIPPED             = "FameBonus.WellEquipped"
	WELLEQUIPPED_DESCRIPTION = "FameBonus.WellEquipped"

	FIRST_BORN             = "FameBonus.FirstBorn"
	FIRST_BORN_DESCRIPTION = "FameBonus.FirstBornDescription"
)

type FameService struct{}

type PCStats struct {
	Shots                  int32
	ShotsThatDamage        int32
	SpecialAbilityUses     int32
	TilesUncovered         int32
	Teleports              int32
	PotionsDrunk           int32
	MonsterKills           int32
	MonsterAssists         int32
	GodKills               int32
	GodAssists             int32
	CubeKills              int32
	OryxKills              int32
	QuestsCompleted        int32
	PirateCavesCompleted   int32
	UndeadLairsCompleted   int32
	AbyssOfDemonsCompleted int32
	SnakePitsCompleted     int32
	SpiderDensCompleted    int32
	SpriteWorldsCompleted  int32
	LevelUpAssists         int32
	MinutesActive          int32
	TombsCompleted         int32
	TrenchesCompleted      int32
	JunglesCompleted       int32
	ManorsCompleted        int32
}

func (service *FameService) GetDeathFame(account *models.Accounts, character *models.Characters, death *models.Death) (*modelxml.DeathXML, error) {
	if account == nil || character == nil {
		return nil, errors.New(
			fmt.Sprintf("account or character are nil, account: %p character: %p", account, character))
	}

	var fame = &modelxml.DeathXML{
		BaseFame:  character.Fame,
		CreatedOn: fmt.Sprintf("%d", character.Createtime.Unix()),
		KilledBy:  death.Killer,
	}

	var stats models.Classstats

	_, err := database.GetDBEngine().Where("accId = ?", account.Id).Get(&stats)

	if err != nil {
		return nil, err
	}

	var bonus float64 = 0

	if character.Charid < 2 {
		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: ANCESTOR_DESCRIPTION,
			Id:          ANCESTOR,
			Value:       fmt.Sprintf("%d", caculateAncestor(float64(character.Fame), &bonus)),
		})
	}

	pc, err := getPCStats(character.Famestats)

	fame.Shots = pc.Shots
	fame.ShotsThatDamage = pc.ShotsThatDamage
	fame.SpecialAbilityUses = pc.SpecialAbilityUses
	fame.TilesUncovered = pc.TilesUncovered
	fame.Teleports = pc.Teleports
	fame.PotionsDrunk = pc.PotionsDrunk
	fame.MonsterKills = pc.MonsterKills
	fame.MonsterAssists = pc.MonsterAssists
	fame.GodKills = pc.GodKills
	fame.GodAssists = pc.GodAssists
	fame.CubeKills = pc.CubeKills
	fame.OryxKills = pc.OryxKills
	fame.QuestsCompleted = pc.QuestsCompleted
	fame.PirateCavesCompleted = pc.PirateCavesCompleted
	fame.UndeadLairsCompleted = pc.UndeadLairsCompleted
	fame.AbyssOfDemonsCompleted = pc.AbyssOfDemonsCompleted
	fame.SnakePitsCompleted = pc.SnakePitsCompleted
	fame.SpiderDensCompleted = pc.SpiderDensCompleted
	fame.SpriteWorldsCompleted = pc.SpriteWorldsCompleted
	fame.LevelUpAssists = pc.LevelUpAssists
	fame.MinutesActive = pc.MinutesActive
	fame.TombsCompleted = pc.TombsCompleted
	fame.TrenchesCompleted = pc.TrenchesCompleted
	fame.JunglesCompleted = pc.JunglesCompleted
	fame.ManorsCompleted = pc.ManorsCompleted

	if err != nil {
		return nil, err
	}

	baseFame := float64(character.Fame)

	if pc.ShotsThatDamage == 0 {
		self := (baseFame + math.Floor(bonus)) * 0.25

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: PACIFIST,
			Id:          PACIFIST_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if pc.PotionsDrunk == 0 {

		self := (baseFame + math.Floor(bonus)) * 0.25

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: THIRSTY,
			Id:          THIRSTY_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if pc.SpecialAbilityUses == 0 {
		self := math.Floor(bonus) + (baseFame+math.Floor(bonus))*0.25

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: MUNDANE,
			Id:          MUNDANE_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if pc.Teleports == 0 {
		self := math.Floor(bonus) + (baseFame+math.Floor(bonus))*0.25

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: BOOTS_ON_GROUND,
			Id:          BOOTS_ON_GROUND_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if pc.PirateCavesCompleted > 0 &&
		pc.UndeadLairsCompleted > 0 &&
		pc.AbyssOfDemonsCompleted > 0 &&
		pc.SnakePitsCompleted > 0 &&
		pc.SpiderDensCompleted > 0 &&
		pc.SpriteWorldsCompleted > 0 &&
		pc.TombsCompleted > 0 &&
		pc.TrenchesCompleted > 0 &&
		pc.JunglesCompleted > 0 &&
		pc.ManorsCompleted > 0 {
		self := math.Floor(bonus) + (baseFame+math.Floor(bonus))*0.25

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: TUNNEL_RAT,
			Id:          TUNNEL_RAT_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + (baseFame+math.Floor(bonus))*0.1
	}

	if float64(pc.GodKills/pc.GodKills+pc.MonsterKills) > 0.1 {

		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: GOD_ENEMY,
			Id:          GOD_ENEMNY_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if float64(pc.GodKills/pc.GodKills+pc.MonsterKills) > 0.5 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: GOD_SLAYER,
			Id:          GOD_SLAYER_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if pc.OryxKills > 0 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: ORYX_SLAYER,
			Id:          ORYX_SLAYER_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if float64(pc.ShotsThatDamage/pc.Shots) > 0.25 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: ACCURATE,
			Id:          ACCURATE_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	if float64(pc.ShotsThatDamage/pc.Shots) > 0.5 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: SHARPSHOOTER,
			Id:          SHARPSHOOTER_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}
	if float64(pc.ShotsThatDamage/pc.Shots) > 0.75 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: SNIPER,
			Id:          SNIPER_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}
	if pc.TilesUncovered > 1000000 {

		self := (baseFame + math.Floor(bonus)) * 0.05

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: EXPLORER,
			Id:          EXPLORER_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}
	if pc.TilesUncovered > 4000000 {
		self := (baseFame + math.Floor(bonus)) * 0.05

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: CARTOGRAPHER,
			Id:          CARTOGRAPHER_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}
	if pc.LevelUpAssists > 100 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: TEAMPLAYER,
			Id:          TEAMPLAYER_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}
	if pc.LevelUpAssists > 1000 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: LEADEROFMAN,
			Id:          LEADEROFMAN_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}
	if pc.QuestsCompleted > 1000 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: DOER_OF_DEEDS,
			Id:          DOER_OF_DEEDS_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}
	if pc.CubeKills == 0 {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: CUBEFRIEND,
			Id:          CUBEFRIEND_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	var eq = 0.0

	items, err := utils.FromCommaSpaceSeparated(character.Items)

	if err != nil {
		return nil, err
	}

	for i := 0; i < 4; i++ {
		if items[i] == -1 {
			continue
		}

		var b = modelxml.GetBonusItem(fmt.Sprintf("0x%x", items[i]))
		if b > 0 {
			eq += baseFame + math.Floor(bonus)*float64(b)/100.0
		}

	}

	fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
		Description: WELLEQUIPPED,
		Id:          WELLEQUIPPED_DESCRIPTION,
		Value:       fmt.Sprintf("%f", math.Floor(eq)),
	})

	bonus = math.Floor(bonus) + math.Floor(eq)

	if (baseFame + math.Floor(bonus)) > float64(stats.Bestfame) {
		self := (baseFame + math.Floor(bonus)) * 0.1

		fame.Bonus = append(fame.Bonus, modelxml.BonusXML{
			Description: FIRST_BORN,
			Id:          FIRST_BORN_DESCRIPTION,
			Value:       fmt.Sprintf("%f", math.Floor(self)),
		})

		bonus = math.Floor(bonus) + self
	}

	charXML, err := GetAccountService().GetCharById(account.Id, character.Id)

	if err != nil {
		return nil, err
	}

	fame.TotalFame = int(bonus)
	fame.Character = *charXML

	return fame, nil
}

func caculateAncestor(baseFame float64, bonus *float64) int {
	//(Math.Floor(((baseFame + Math.Floor(bonus)) * 0.1) + 20)
	//Math.Floor(bonus) + ((baseFame + Math.Floor(bonus)) * 0.1) + 20;
	result := int(math.Floor(((baseFame + math.Floor(*bonus)) * 0.1) + 20))

	*bonus = *bonus + (((baseFame + *bonus) * 0.1) + 20)

	return result
}

func getPCStats(encoded string) (*PCStats, error) {
	var pc = &PCStats{}

	fameBytes, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(fameBytes)

	i, _ := buffer.ReadByte()
	for {
		var y int32
		err = binary.Read(buffer, binary.BigEndian, &y)

		if err == io.EOF {
			break
		}

		switch i {
		case 0:
			pc.Shots = y
			break
		case 1:
			pc.ShotsThatDamage = y
			break
		case 2:
			pc.SpecialAbilityUses = y
			break
		case 3:
			pc.TilesUncovered = y
			break
		case 4:
			pc.Teleports = y
			break
		case 5:
			pc.PotionsDrunk = y
			break
		case 6:
			pc.MonsterKills = y
			break
		case 7:
			pc.MonsterAssists = y
			break
		case 8:
			pc.GodKills = y
			break
		case 9:
			pc.GodAssists = y
			break
		case 10:
			pc.CubeKills = y
			break
		case 11:
			pc.OryxKills = y
			break
		case 12:
			pc.QuestsCompleted = y
			break
		case 13:
			pc.PirateCavesCompleted = y
			break
		case 14:
			pc.UndeadLairsCompleted = y
			break
		case 15:
			pc.AbyssOfDemonsCompleted = y
			break
		case 16:
			pc.SnakePitsCompleted = y
			break
		case 17:
			pc.SpiderDensCompleted = y
			break
		case 18:
			pc.SpriteWorldsCompleted = y
			break
		case 19:
			pc.LevelUpAssists = y
			break
		case 20:
			pc.MinutesActive = y
			break
		case 21:
			pc.TombsCompleted = y
			break
		case 22:
			pc.TrenchesCompleted = y
			break
		case 23:
			pc.JunglesCompleted = y
			break
		case 24:
			pc.ManorsCompleted = y
			break

		}

		i, _ = buffer.ReadByte()

		if i < 0 {
			break
		}
	}

	return pc, nil
}

func GetFameService() *FameService {
	return fame
}
