package migrations

import (
	"database/sql"
	"go-app/internal/database"
)

func init() {
	database.RegisterMigration(database.Migration{
		Version:     1,
		Description: "Create user table",
		Up: func(db *sql.DB) error {
			_, err := db.Exec(`
				CREATE TABLE IF NOT EXISTS users (
					id SERIAL PRIMARY KEY,
					username VARCHAR(50) UNIQUE NOT NULL,
					email VARCHAR(100) UNIQUE NOT NULL,
					password_hash VARCHAR(100) NOT NULL,
					first_name VARCHAR(50),
					last_name VARCHAR(50),
					date_of_birth DATE,
					created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
					last_login TIMESTAMP WITH TIME ZONE
				)
			`)
			return err
		},
		Down: func(db *sql.DB) error {
			_, err := db.Exec("DROP TABLE IF EXISTS users")
			return err
		},
	})
}
