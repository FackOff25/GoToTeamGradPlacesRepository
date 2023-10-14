package configReader

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	ServerPort string `toml:"serverPort"`
}

func NewConfig(configFilePath string) (*Config, error) {
	config := &Config{}

	if _, err := toml.DecodeFile(configFilePath, config); err != nil {
		return nil, err
	}

	return config, nil
}
