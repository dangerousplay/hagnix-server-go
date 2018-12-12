package utils

import (
	"github.com/go-xorm/xorm"
	"github.com/ivahaev/go-logger"
	"github.com/kataras/iris"
	"hagnix-server-go1/routes/messages"
)

func DefaultErrorHandler(ctx iris.Context, err error) bool {
	if err != nil {
		ctx.StatusCode(500)
		ctx.XML(messages.Error{RawXml: "Internal Server error: " + err.Error()})
		logger.Warn(err)
		return true
	}
	return false
}

func HandleSessionRowsUpdated(ctx iris.Context, session *xorm.Session, err error, rows int64, rowsMax int64) bool {
	if DefaultErrorHandler(ctx, err) {
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
