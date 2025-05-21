package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connect to database
	dsn := "host=localhost user=postgres password=postgres dbname=develop port=5435 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Clear existing data
	if _, err := db.Exec("DELETE FROM todos"); err != nil {
		log.Fatalf("failed to clear todos: %v", err)
	}
	if _, err := db.Exec("DELETE FROM users"); err != nil {
		log.Fatalf("failed to clear users: %v", err)
	}

	// Reset sequences
	if _, err := db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1"); err != nil {
		log.Fatalf("failed to reset users sequence: %v", err)
	}
	if _, err := db.Exec("ALTER SEQUENCE todos_id_seq RESTART WITH 1"); err != nil {
		log.Fatalf("failed to reset todos sequence: %v", err)
	}

	// Insert users
	users := []string{"Alice", "Bob", "Charlie"}
	for _, name := range users {
		if _, err := db.Exec("INSERT INTO users (name) VALUES ($1)", name); err != nil {
			log.Fatalf("failed to insert user %s: %v", name, err)
		}
	}

	// Insert todos
	todos := []struct {
		text   string
		done   bool
		userID int
	}{
		{"Buy milk", false, 1},
		{"Learn GraphQL", false, 1},
		{"Build a todo app", true, 2},
		{"Write tests", false, 3},
	}

	for _, todo := range todos {
		if _, err := db.Exec(
			"INSERT INTO todos (text, done, user_id) VALUES ($1, $2, $3)",
			todo.text, todo.done, todo.userID,
		); err != nil {
			log.Fatalf("failed to insert todo %s: %v", todo.text, err)
		}
	}

	log.Println("Seeding completed successfully!")
}