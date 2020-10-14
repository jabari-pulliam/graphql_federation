package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"
	"prices/graph"
	"prices/graph/generated"
	"prices/middleware"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
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

	r := gin.Default()
	r.Use(middleware.WidgetLoaderMiddleware(conn))
	r.POST("/query", ProvideGraphqlHandler(conn))

	server := http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ProvideGraphqlHandler(conn *pgx.Conn) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the playground.go file
	// Resolver is in the resolver.go file
	config := generated.Config{Resolvers: &graph.Resolver{DbConn: conn}}

	server := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
	}
}
