package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/victorrenop/raven/internal/adapters"
)

type ConfigService interface {
	GetConfigByVersion(ctx context.Context, config.ConfigName string, config.ConfigEnv string,  config.ConfigVersion string) (domain.Config, error),
	GetConfig(ctx context.Context, config.ConfigName string, config.ConfigEnv string) (domain.Config, error),
}

type ConfigService struct {
	configRepo domain.ConfigRepository
}

func (ConfigService *ConfigService) GetConfigByVersion(ctx context.Context, configName string, configEnv string, configVersion string) (domain.Config, error) {
	config, err:= ConfigService.configRepo.GetWithVersion(ctx, configName, configEnv, configVersion)
	if err != nil {
		return config, fmt.Errorf("error getting config: %w", err)
	}
	config.ConfigName = configName
	config.ConfigEnv = configEnv
	config.ConfigVersion = configVersion
	config.ConfigData = ConfigData
	config.ConfigState = ConfigState

	return config, nil
}

func (ConfigService *ConfigService) GetConfig(ctx context.Context, configName string, configEnv string) (domain.Config, error) {
	config, err:= ConfigService.configRepo.GetLatest(ctx, configName, configEnv, configVersion)
	if err != nil {
		return config, fmt.Errorf("error getting config: %w", err)
	}
	config.ConfigName = configName
	config.ConfigEnv = configEnv
	config.ConfigVersion = configVersion
	config.ConfigData = ConfigData
	config.ConfigState = ConfigState

	return config, nil
}


