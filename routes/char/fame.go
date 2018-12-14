package char

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/messages"
	"hagnix-server-go1/routes/utils"
	"hagnix-server-go1/service"
	"io"
)

func handleFame(ctx iris.Context) {
	accountId := ctx.PostValue("accountId")
	charId := ctx.PostValue("charId")

	if len(accountId) < 1 || len(charId) < 1 {
		ctx.XML(messages.BadRequest)
		return
	}

	var death models.Death

	success, err := database.GetDBEngine().Where("accId = ? AND charId = ?", accountId, charId).Get(&death)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if !success {
		ctx.XML(messages.BadRequest)
		return
	}

	//TODO implement Fame

	acxml, acc, err := service.GetAccountService().VerifyGenerateAccountXMLbyId(accountId)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	if acxml == nil || acc == nil {
		return
	}

	fameBytes, err := base64.StdEncoding.DecodeString(death.Famestats)

	if utils.DefaultErrorHandler(ctx, err) {
		return
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
			break
		case 1:
			break
		case 2:
			break
		case 3:
			break
		case 4:
			break
		case 5:
			break
		case 6:
			break
		case 7:
			break
		case 8:
			break
		case 9:
			break
		case 10:
			break

		}

		i, _ = buffer.ReadByte()

		if i < 0 {
			break
		}
	}
}
