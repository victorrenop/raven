package adapters

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/victorrenop/raven/internal/domain"
	"log"
)

type ConfigSQLiteRepository struct {
	client sql.DB
}

func NewConfigSQLiteRepository(client sql.DB) (*ConfigSQLiteRepository, error) {
	return &ConfigSQLiteRepository{
		client: client,
	}, nil
}

func (configSQLiteRepository *ConfigSQLiteRepository) GetLatest(ctx context.Context, projectName string, env string) (domain.Config, error) {

	stmt, err := configSQLiteRepository.client.Prepare("SELECT * FROM tablename WHERE project_name = ? AND env = ? ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	config := domain.Config{}
	err = stmt.QueryRow(projectName, env).Scan(&config.ConfigVersion, &config.ConfigProjectName, &config.ConfigEnv, &config.ConfigCreatedAt, &config.ConfigState, &config.ConfigData)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}

func (configSQLiteRepository *ConfigSQLiteRepository) GetWithVersion(ctx context.Context, projectName string, env string, version int) (domain.Config, error) {

	stmt, err := configSQLiteRepository.client.Prepare("SELECT * FROM tablename WHERE project_name = ? AND env = ? AND version = ? LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	config := domain.Config{}
	err = stmt.QueryRow(projectName, env, version).Scan(&config.ConfigVersion, &config.ConfigProjectName, &config.ConfigEnv, &config.ConfigCreatedAt, &config.ConfigState, &config.ConfigData)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}
