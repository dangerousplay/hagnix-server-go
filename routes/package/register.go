package _package

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/database"
	"hagnix-server-go1/database/models"
	"hagnix-server-go1/routes/modelxml"
	"hagnix-server-go1/routes/utils"
	"time"
)

func RegisterRoutes(app *iris.Application) {
	papp := app.Party("package")
	papp.Post("/getPackages", handleGetPackage)
}

func handleGetPackage(ctx iris.Context) {
	packages := []models.Packages{}
	err := database.GetDBEngine().Where("endDate >= ?", time.Now().Unix()).Find(&packages)

	if utils.DefaultErrorHandler(ctx, err) {
		return
	}

	packagesXML := []modelxml.Package{}

	for _, v := range packages {

		packagesXML = append(packagesXML, modelxml.Package{
			Id:          v.Id,
			Name:        v.Name,
			Price:       v.Price,
			Quantity:    v.Quantity,
			MaxPurchase: v.Maxpurchase,
			Weight:      v.Weight,
			BgURL:       v.Bgurl,
			EndDate:     v.Enddate.Format("05/12/2006 13:17:22 GMT-0000"),
		})
	}

	xmlPackage := &modelxml.PackageResponse{
		Packages: modelxml.PackageWrapper{Packages: packagesXML},
	}

	ctx.XML(xmlPackage)
}
