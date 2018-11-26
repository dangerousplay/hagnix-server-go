package utils

import (
	"github.com/InVisionApp/go-logger"
	"github.com/go-xorm/xorm"
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

func HandleSessionRowsUpdated(ctx iris.Context, session *xorm.Session, err error, log log.Logger, rows int64, rowsMax int64) bool {
	if DefaultErrorHandler(ctx, err, log) {
		session.Rollback()
		return true
	}

	if rows != rowsMax {
		ctx.XML(messages.DefaultError)
		session.Rollback()
		return true
	}

	return false
}
