package database

import (
	"database/sql"
	"fmt"
	"sort"
)

type Migration struct {
	Version     int
	Description string
	Up          func(*sql.DB) error
	Down        func(*sql.DB) error
}

var migration []Migration

func RegisterMigration(m Migration) {
	migration = append(migration, m)
}

func RunMigration(db *sql.DB) error {
	sort.Slice(migration, func(i, j int) bool {
		return migration[i].Version < migration[j].Version
	})

	for _, m := range migration {
		fmt.Printf("Running migration %d: %s\n", m.Version, m.Description)
		if err := m.Up(db); err != nil {
			return fmt.Errorf("Error running migration %d: %v", m.Version, err)
		}
	}
	return nil
}
