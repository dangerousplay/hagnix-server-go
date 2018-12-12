package config

import (
	"github.com/ivahaev/go-logger"
	"os"
	"path/filepath"
)

var config = &ROTMGConfig{}
var filesConfig = &FilesConfig{}
var executablePath string

func Init() {
	variable := os.Getenv("SETTINGS_VARIABLE")

	logger.Info("Loading server configuration...")

	local, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	executablePath = local

	if len(variable) > 0 {
		err := config.LoadFromContent(variable)

		if err != nil {
			panic(err)
		}
	} else {

		if err != nil {
			panic(err)
		}

		err = config.LoadFromFile(executablePath + "/app/server.json")

		if err != nil {
			panic(err)
		}
	}

	initWatcher()

	initGRPC()

	config.Loaded = true
}
