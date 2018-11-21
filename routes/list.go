package routes

import (
	"encoding/xml"
	"fmt"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
)

type FameListElement struct {
	XMLName   xml.Name `xml:"FameListElem"`
	AccountId int      `xml:"accountId,attr"`
	CharId    int      `xml:"charId,attr"`
	Name      string   `xml:"name"`
	CharType  string   `xml:"ObjectType"`
	Text1     string   `xml:"Text1"`
	Text2     string   `xml:"Text2"`
	Skin      string   `xml:"Texture"`
	Items     string   `xml:"Equipment"`
	TotalFame string   `xml:"TotalFame"`
}

type FameList struct {
	XMLName  xml.Name `xml:"FameList"`
	Timespan string   `xml:"timespan,attr"`
	list     []FameListElement
}

func HandleFameList(ctx iris.Context) {
	timespan := ctx.URLParam("timespan")
	accountId := ctx.URLParam("accountId")

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
		ctx.WriteString("<Error>Invalid fame list</Error>")
		return
	}

	query := database.GetDBEngine().Where(timespan)

	if len(accountId) > 0 {
		query = query.Or("(accId=? AND chrId=?)", ctx.URLParam("accountId"), ctx.URLParam("charId"))
	}

	fmt.Printf(where)
}
