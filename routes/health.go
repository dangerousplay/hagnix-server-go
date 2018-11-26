package routes

import (
	"github.com/kataras/iris"
	"hagnix-server-go1/config"
	"hagnix-server-go1/database"
)

type HealthState struct {
	State          int      `json:"state"`
	ErrorsMessages []string `json:"errors_messages"`
}

func handleHealth(ctx iris.Context) {
	health := &HealthState{State: 200}

	checkDatabase(health)
	checkConfig(health)

	if len(health.ErrorsMessages) > 0 {
		health.State = 500
	}

	ctx.JSON(health)
}

func checkConfig(health *HealthState) {
	if !config.GetConfig().Loaded {
		health.ErrorsMessages = append(health.ErrorsMessages, "Config not loaded!")
	}
}

func checkDatabase(health *HealthState) {
	err := database.GetDBEngine().Ping()

	if err != nil {
		health.ErrorsMessages = append(health.ErrorsMessages, "Database: "+err.Error())
	}
}
