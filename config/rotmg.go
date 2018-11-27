package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

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

type FilesConfig struct {
	Init      string
	Languages map[string]string
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

func checkInitialized() {
	if config == nil || !config.Loaded {
		Init()
	}
}

func GetFilesConfig() *FilesConfig {
	checkInitialized()

	return filesConfig
}

func GetConfig() *ROTMGConfig {
	checkInitialized()

	return config
}
