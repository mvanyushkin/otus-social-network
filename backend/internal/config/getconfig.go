package config

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"os"
	"path"
)

func GetConfig(configFilePath *string) (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configFileAbsPath := path.Join(wd, *configFilePath)
	_, err = os.Stat(configFileAbsPath)
	if configFilePath == nil || os.IsNotExist(err) {
		defaultConfigPath := path.Join(wd, "local_config.json")
		configFilePath = &defaultConfigPath
	}

	loader := confita.NewLoader(
		env.NewBackend(),
		file.NewBackend(configFileAbsPath),
	)

	cfg := Config{}
	err = loader.Load(context.Background(), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
