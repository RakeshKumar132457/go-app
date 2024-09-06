package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-app/internal/config"
)

func SetupDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	if err = RunMigration(db); err != nil {
		return nil, fmt.Errorf("Error running migrations: %v", err)
	}

	return db, nil
}
