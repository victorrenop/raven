package domain

import (
	"context"
)

type ConfigRepository interface {
	GetLatest(ctx context.Context, projectName string, env string) (Config, error)
	GetWithVersion(ctx context.Context, projectName string, env string, configVersion int) (Config, error)
	Write(ctx context.Context, projectName string, env string, configData map[string]interface{}) error
}
