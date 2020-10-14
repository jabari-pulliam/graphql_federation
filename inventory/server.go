package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"inventory/graph"
	"inventory/graph/generated"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbUrl := fmt.Sprintf("postgresql://%s:%s@db:5432/%s", "postgres", "testtest", "widgets")
	conn, err := pgx.Connect(context.Background(), dbUrl)
	for err != nil {
		log.Println("Failed to connect to DB, reconnecting...")
		time.Sleep(100 * time.Millisecond)
		conn, err = pgx.Connect(context.Background(), dbUrl)
	}
	defer conn.Close(context.Background())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DbConn: conn}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
