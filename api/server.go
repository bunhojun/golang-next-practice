package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/bunhojun/golang-next-practice/api/graph"
	"github.com/bunhojun/golang-next-practice/api/migrations"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Open database
	dsn := "host=localhost user=postgres password=postgres dbname=develop port=5435 sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := migrations.RunAll(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	resolver := &graph.Resolver{DB: db}
    // Create the GraphQL handler

    srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	srv.Use(extension.Introspection{})

    // Add CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })

    // Create a new mux
    mux := http.NewServeMux()
    
    // Add the playground handler
    mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
    
    // Add the GraphQL handler with CORS
    mux.Handle("/query", c.Handler(srv))

    // Configure the server
    srv.AddTransport(transport.POST{})
    srv.AddTransport(transport.GET{})
    srv.AddTransport(transport.Options{})

    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, mux))
}
