package tests

import (
	"database/sql"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/go-cmp/cmp"
	_ "github.com/mattn/go-sqlite3"
	"github.com/victorrenop/raven/internal/adapters"
	"github.com/victorrenop/raven/internal/domain"
)

func TestGetLatestSuccessfullyReturnsConfig(t *testing.T) {
	testSqliteClient, _ := sql.Open("sqlite3", "file:../test.db?cache=shared")
	testConfigSqliteRepository, _ := adapters.NewConfigSQLiteRepository(testSqliteClient)

	testProjectName := "test_project"
	testEnv := "dev"

	expectedResult := domain.Config{
		ConfigVersion:     2,
		ConfigProjectName: testProjectName,
		ConfigEnv:         testEnv,
		ConfigCreatedAt:   "2023-01-02T00:00:00Z",
		ConfigState:       "active",
		ConfigData: map[string]interface{}{
			"some_config":    "some_value",
			"another_config": 456.0,
		},
	}

	actualResult, err := testConfigSqliteRepository.GetLatest(nil, testProjectName, testEnv)

	assert.Equal(t, nil, err)
	if !cmp.Equal(actualResult, expectedResult) {
		t.Errorf("got %+v, want %+v", actualResult, expectedResult)
	}
}

func TestGetLatestFailsWithNonExistantProject(t *testing.T) {
	testSqliteClient, _ := sql.Open("sqlite3", "file:../test.db?cache=shared")
	testConfigSqliteRepository, _ := adapters.NewConfigSQLiteRepository(testSqliteClient)

	testProjectName := "non_existant_project_name"
	testEnv := "dev"

	_, err := testConfigSqliteRepository.GetLatest(nil, testProjectName, testEnv)

	assert.NotEqual(t, nil, err)
}

func TestGetLatestFailsWithNonExistantEnv(t *testing.T) {
	testSqliteClient, _ := sql.Open("sqlite3", "file:../test.db?cache=shared")
	testConfigSqliteRepository, _ := adapters.NewConfigSQLiteRepository(testSqliteClient)

	testProjectName := "test_project"
	testEnv := "non_existant_env"

	_, err := testConfigSqliteRepository.GetLatest(nil, testProjectName, testEnv)

	assert.NotEqual(t, nil, err)
}

func TestGetWithVersionSuccessfullyReturnsConfig(t *testing.T) {
	testSqliteClient, _ := sql.Open("sqlite3", "file:../test.db?cache=shared")
	testConfigSqliteRepository, _ := adapters.NewConfigSQLiteRepository(testSqliteClient)

	testProjectName := "test_project"
	testEnv := "dev"

	expectedResult := domain.Config{
		ConfigVersion:     1,
		ConfigProjectName: testProjectName,
		ConfigEnv:         testEnv,
		ConfigCreatedAt:   "2023-01-01T00:00:00Z",
		ConfigState:       "active",
		ConfigData: map[string]interface{}{
			"some_config":    "some_value",
			"another_config": 123.0,
		},
	}

	actualResult, err := testConfigSqliteRepository.GetWithVersion(nil, testProjectName, testEnv, 1)

	assert.Equal(t, nil, err)
	if !cmp.Equal(actualResult, expectedResult) {
		t.Errorf("got %+v, want %+v", actualResult, expectedResult)
	}
}

func TestGetWithVersionFailsWithNonExistantVersion(t *testing.T) {
	testSqliteClient, _ := sql.Open("sqlite3", "file:../test.db?cache=shared")
	testConfigSqliteRepository, _ := adapters.NewConfigSQLiteRepository(testSqliteClient)

	testProjectName := "test_project"
	testEnv := "dev"

	_, err := testConfigSqliteRepository.GetWithVersion(nil, testProjectName, testEnv, 3)

	assert.NotEqual(t, nil, err)
}

func TestSaveSuccessfullySavesConfig(t *testing.T) {
	testSqliteClient, _ := sql.Open("sqlite3", "file:../test.db?cache=shared")
	testConfigSqliteRepository, _ := adapters.NewConfigSQLiteRepository(testSqliteClient)

	testProjectName := "test_project"
	testEnv := "dev"

	expectedResult := domain.Config{
		ConfigVersion:     3,
		ConfigProjectName: testProjectName,
		ConfigEnv:         testEnv,
		ConfigCreatedAt:   "2023-01-03T00:00:00Z",
		ConfigState:       "active",
		ConfigData: map[string]interface{}{
			"some_config":    "some_value",
			"another_config": 999.0,
		},
	}
	err := testConfigSqliteRepository.Save(nil, expectedResult)
	assert.Equal(t, nil, err)
	actualResult, err := testConfigSqliteRepository.GetLatest(nil, testProjectName, testEnv)

	assert.Equal(t, nil, err)
	if !cmp.Equal(actualResult, expectedResult) {
		t.Errorf("got %+v, want %+v", actualResult, expectedResult)
	}
}
