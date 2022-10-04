package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	Sqlx sqlx.DB
}

func connect(databaseUrl string) (*sqlx.DB, error) {
	database, err := sqlx.Connect("postgres", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = database.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return database, nil
}

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("create postgres driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("create new migrate with db instance: %v", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("running the migration: %v", err)
	}

	return nil
}

func NewDatabase(databaseUrl string) (*sqlx.DB, error) {
	db, err := connect(databaseUrl)
	if err != nil {
		return nil, err
	}

	err = runMigrations(db.DB)
	if err != nil {
		return nil, err
	}

	return db, nil
}
