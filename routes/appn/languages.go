package appn

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
)

func handleLanguage(ctx iris.Context) {
	language := ctx.URLParam("languageType")

	file := config.GetFilesConfig().Languages[language]

	ctx.ContentType("application/json")
	ctx.WriteString(file)
}
