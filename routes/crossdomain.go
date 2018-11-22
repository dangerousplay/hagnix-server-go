package routes

import (
	"encoding/xml"
	"github.com/kataras/iris"
)

type AllowAcess struct {
	XMLName xml.Name `xml:"allow-access-from"`
	Domain  string   `xml:"domain,attr"`
}

type CrossDomain struct {
	XMLName xml.Name `xml:"cross-domain-policy"`
	Origin  AllowAcess
}

var cors = CrossDomain{
	Origin: AllowAcess{
		Domain: "*",
	},
}

func handleCrossDomain(ctx iris.Context) {
	ctx.ContentType("text/*")
	ctx.XML(cors)
}
