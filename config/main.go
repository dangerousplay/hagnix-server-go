package config

import (
	"github.com/InVisionApp/go-logger"
	"os"
	"path/filepath"
)

var logger = log.NewSimple()
var config = &ROTMGConfig{}
var filesConfig = &FilesConfig{}

func Init() {
	variable := os.Getenv("SETTINGS_VARIABLE")

	logger.Info("Loading server configuration...")

	if len(variable) > 0 {
		err := config.LoadFromContent(variable)

		if err != nil {
			panic(err)
		}
	} else {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			panic(err)
		}

		err = config.LoadFromFile(dir + "/server.json")

		if err != nil {
			panic(err)
		}
	}

	initWatcher()

	config.Loaded = true
}
