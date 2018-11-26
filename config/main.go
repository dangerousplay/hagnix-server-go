package config

import (
	"encoding/json"
	"github.com/InVisionApp/go-logger"
	"io/ioutil"
	"os"
	"path/filepath"
)

var logger = log.NewSimple()
var config = &ROTMGConfig{}

type Servers struct {
	Name     string
	Address  string
	Location string
}

type ServerConfig struct {
	VerifyEmail  bool   `json:"VerifyEmail"`
	ServerDomain string `json:"ServerDomain"`
	Servers      []Servers
}

type ROTMGConfig struct {
	ServerConfig *ServerConfig
	Loaded       bool
}

func (cfg *ROTMGConfig) LoadFromVariable(variable string) error {
	values := os.Getenv(variable)

	config := &ServerConfig{}

	err := json.Unmarshal([]byte(values), config)

	cfg.ServerConfig = config

	return err
}

func (cfg *ROTMGConfig) LoadFromContent(content string) error {
	config := &ServerConfig{}

	err := json.Unmarshal([]byte(content), config)

	cfg.ServerConfig = config

	return err
}

func (cfg *ROTMGConfig) LoadFromFile(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	config := &ServerConfig{}

	err = json.Unmarshal(bytes, config)

	cfg.ServerConfig = config

	return err
}

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

	config.Loaded = true
}

func GetConfig() *ROTMGConfig {
	if config == nil || !config.Loaded {
		Init()
	}

	return config
}
