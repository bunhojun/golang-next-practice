package migrations

import (
	"database/sql"
	"log"
)

// Migration represents a single database migration
type Migration struct {
	Name string
	Fn   func(*sql.DB) error
}

// RunAll executes all migrations in order
func RunAll(db *sql.DB) error {
	migrations := []Migration{
		{
			Name: "genesis",
			Fn:   Genesis,
		},
		// Add new migrations here, they'll run in order
		// {
		//     Name: "add_todo_priority",
		//     Fn:   AddTodoPriority,
		// },
	}

	for _, m := range migrations {
		log.Printf("Running migration: %s", m.Name)
		if err := m.Fn(db); err != nil {
			return err
		}
		log.Printf("Completed migration: %s", m.Name)
	}

	return nil
}