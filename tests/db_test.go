package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	// pseudo-code, some implementation excluded:
	//
	// 1. create test.db if it does not exist
	// 2. run our DDL statements to create the required tables if they do not exist
	// 3. run our tests
	// 4. truncate the test db tables

	db, err := sql.Open("sqlite3", "file:../test.db?cache=shared")
	// db, err := sql.Open("sqlite3", "file:..")
	if err != nil {
		return -1, fmt.Errorf("could not connect to database: %w", err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS config_data; CREATE TABLE IF NOT EXISTS `config_data` (`version` INTEGER PRIMARY KEY AUTOINCREMENT, `project_name` TEXT, `env` TEXT, `created_at` DATETIME, `state` TEXT, `data` TEXT)")
	if err != nil {
		return -1, fmt.Errorf("could not create test table: %w", err)
	}

	_, err = db.Exec(
		`INSERT INTO config_data(project_name, env, created_at, state, data)
		VALUES('test_project', 'dev', '2023-01-01T00:00:00.000Z', 'active', '{"some_config": "some_value", "another_config": 123}'),
		('test_project', 'dev', '2023-01-02T00:00:00.000Z', 'active', '{"some_config": "some_value", "another_config": 456}');
	`)
	if err != nil {
		return -1, fmt.Errorf("could not create test table: %w", err)
	}

	// truncates all test data after the tests are run
	defer func() {
		_, _ = db.Exec("DELETE FROM config_data")

		db.Close()
	}()

	return m.Run(), nil
}
