package migrations

import (
	"database/sql"
	"log"
)

// Genesis runs the initial database migration
func Genesis(db *sql.DB) error {
	schema := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			text TEXT NOT NULL,
			done BOOLEAN NOT NULL DEFAULT false,
			user_id INTEGER NOT NULL REFERENCES users(id)
		);
	`
	if _, err := db.Exec(schema); err != nil {
		return err
	}
	log.Println("Genesis migration completed successfully")
	return nil
}