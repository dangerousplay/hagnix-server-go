package routes

import "github.com/kataras/iris"

func HandleCrossDomain(ctx iris.Context){
	ctx.ContentType("text/*")
	ctx.WriteString("<cross-domain-policy><allow-access-from domain=\"*\"/></cross-domain-policy>")
}