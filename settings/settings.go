package settings

import (
	_ "embed"
	"time"

	yaml "gopkg.in/yaml.v2"
)

//go:embed settings.yaml
var settingsContent []byte

type DatabaseSettings struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Settings struct {
	Port          string           `yaml:"port"`
	Address       string           `yaml:"address"`
	TokenKey      string           `yaml:"token_key"`
	TokenDuration time.Duration    `yaml:"token_duration"`
	DB            DatabaseSettings `yaml:"database"`
}

func LoadConfig() (*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingsContent, &s)
	if err != nil {
		return &Settings{}, err
	}

	return &s, nil
}
