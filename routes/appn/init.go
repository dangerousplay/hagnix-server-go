package appn

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
)

func handleInit(ctx iris.Context) {
	ctx.ContentType("application/xml")
	ctx.WriteString(config.GetFilesConfig().Init)
}
