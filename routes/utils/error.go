package utils

import (
	"github.com/InVisionApp/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/messages"
)

func DefaultErrorHandler(ctx iris.Context, err error, log log.Logger) bool {
	if err != nil {
		ctx.StatusCode(500)
		ctx.XML(messages.Error{RawXml: "Internal Server error: " + err.Error()})
		log.Warn(err)
		return true
	}
	return false
}
