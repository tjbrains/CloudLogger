// Copyright 2024 FlexCDN root@flexcdn.cn. All rights reserved. Official site: https://flexcdn.cn .

package configs

import (
	"github.com/tjbrains/TeaGo/Tea"
	_ "github.com/tjbrains/TeaGo/bootstrap"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port int `yaml:"port"`
}

func NewConfig() *Config {
	return &Config{
		Port: 8010,
	}
}

func DefaultConfig() *Config {
	return NewConfig()
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile(Tea.ConfigFile("config.yaml"))
	if err != nil {
		return nil, err
	}

	var config = NewConfig()
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
