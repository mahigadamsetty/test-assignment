package db

import (
	"database/sql"
	"os"
	"path/filepath"
)

func RunMigrations(db *sql.DB, migrationsPath string) error {
	entries, err := os.ReadDir(migrationsPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(migrationsPath, entry.Name())
		sqlBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if _, err := db.Exec(string(sqlBytes)); err != nil {
			return err
		}
	}

	return nil
}
