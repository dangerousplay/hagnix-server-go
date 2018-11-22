package routes

import (
	"encoding/xml"
	"fmt"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
)

type FameListElement struct {
	XMLName   xml.Name `xml:"FameListElem"`
	AccountId int      `xml:"accountId,attr"`
	CharId    int      `xml:"charId,attr"`
	Name      string   `xml:"name"`
	CharType  int      `xml:"ObjectType"`
	Text1     int      `xml:"Tex1"`
	Text2     int      `xml:"Tex2"`
	Skin      int      `xml:"Texture"`
	Items     string   `xml:"Equipment"`
	TotalFame int      `xml:"TotalFame"`
}

type FameList struct {
	XMLName  xml.Name `xml:"FameList"`
	Timespan string   `xml:"timespan,attr"`
	List     []FameListElement
}

func handleFameList(ctx iris.Context) {
	timespan := ctx.URLParam("timespan")
	accountId := ctx.URLParam("accountId")
	charId := ctx.URLParam("charId")

	var where string

	switch timespan {
	case "week":
		where = "(time >= DATE_SUB(NOW(), INTERVAL 1 WEEK))"
		break
	case "month":
		where = "(time >= DATE_SUB(NOW(), INTERVAL 1 MONTH))"
		break
	case "all":
		where = "TRUE"
		break
	default:
		ctx.StatusCode(400)
		ctx.XML(messages.Error{RawXml: "Invalid fame list"})
		return
	}

	query := database.GetDBEngine().Where(where)

	if len(accountId) > 0 && len(charId) > 0 {
		query = query.Or("(accId = ? AND chrId = ?)", accountId, charId)
	}

	fameElements := []FameListElement{}

	err := query.Desc("totalFame").Limit(20).Iterate(&models.Death{}, func(idx int, bean interface{}) error {
		death, _ := bean.(*models.Death)
		fameElements = append(fameElements, FameListElement{
			AccountId: death.AccountId,
			CharId:    death.CharacterId,
			Name:      death.Name,
			CharType:  death.CharType,
			Text1:     death.Tex1,
			Text2:     death.Tex2,
			Skin:      death.Skin,
			Items:     death.Items,
			TotalFame: death.TotalFame,
		})
		return nil
	})

	if err != nil {
		ctx.StatusCode(500)
		ctx.WriteString("Internal server error")
		fmt.Println(err)
	}

	ctx.XML(FameList{
		List:     fameElements,
		Timespan: timespan,
	})
}
