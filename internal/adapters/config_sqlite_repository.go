package adapters

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/victorrenop/raven/internal/domain"
	"log"
)
// ConfigSQLiteRepository is a struct representing a SQLite implementation of the ConfigRepository interface
// It has a client field of type sql.DB used to interact with the SQLite database
type ConfigSQLiteRepository struct {
	client sql.DB
}
// NewConfigSQLiteRepository is a function that creates a new instance of ConfigSQLiteRepository with a given client
// It returns a pointer to the created ConfigSQLiteRepository and an error if any occurred
func NewConfigSQLiteRepository(client sql.DB) (*ConfigSQLiteRepository, error) {
	return &ConfigSQLiteRepository{
		client: client,
	}, nil
}
// GetLatest is a method of ConfigSQLiteRepository that retrieves the latest version of a config for a given project and environment from the SQLite database
// It takes a context.Context object, a projectName string, and an env string as parameters
// It returns a domain.Config object and an error if any occurred
func (configSQLiteRepository *ConfigSQLiteRepository) GetLatest(ctx context.Context, projectName string, env string) (domain.Config, error) {
	// prepare a SQL statement to select the latest config version for the given project and environment
	stmt, err := configSQLiteRepository.client.Prepare("SELECT * FROM tablename WHERE project_name = ? AND env = ? ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// create an empty Config struct to hold the retrieved config data
	config := domain.Config{}
	// execute the prepared SQL statement with the given parameters and scan the result into the Config struct
	err = stmt.QueryRow(projectName, env).Scan(&config.ConfigVersion, &config.ConfigProjectName, &config.ConfigEnv, &config.ConfigCreatedAt, &config.ConfigState, &config.ConfigData)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}
// GetWithVersion is a method of ConfigSQLiteRepository that retrieves a specific version of a config for a given project and environment from the SQLite database
// It takes a context.Context object, a projectName string, an env string, and a version int as parameters
// It returns a domain.Config object and an error if any occurred
func (configSQLiteRepository *ConfigSQLiteRepository) GetWithVersion(ctx context.Context, projectName string, env string, version int) (domain.Config, error) {
	// prepare a SQL statement to select a specific config version for the given project, environment, and version
	stmt, err := configSQLiteRepository.client.Prepare("SELECT * FROM tablename WHERE project_name = ? AND env = ? AND version = ? LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	// create an empty Config struct to hold the retrieved config data
	config := domain.Config{}
	// execute the prepared SQL statement with the given parameters and scan the result into the Config struct
	err = stmt.QueryRow(projectName, env, version).Scan(&config.ConfigVersion, &config.ConfigProjectName, &config.ConfigEnv, &config.ConfigCreatedAt, &config.ConfigState, &config.ConfigData)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	return config, nil
}

// Save is a method of ConfigSQLiteRepository that saves a config to the SQLite database
// It takes a context.Context object and a domain.Config object as parameters
// It returns an error if any occurred
func (configSQLiteRepository *ConfigSQLiteRepository) Save(ctx context.Context, config domain.Config) error {
	// prepare a SQL statement to insert a new config into the database
	stmt, err := configSQLiteRepository.client.Prepare("INSERT INTO tablename (project_name, env, version, created_at, state, data) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(config.ConfigProjectName, config.ConfigEnv, config.ConfigVersion, config.ConfigCreatedAt, config.ConfigState, config.ConfigData)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

