package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
	"widgets/graph"
	"widgets/graph/generated"
	"widgets/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const (
	defaultPort  = "8080"
	numWidgets   = 100000
	minPrice     = 0
	maxPrice     = 100
	minInventory = 0
	maxInventory = 100
	minSize      = 1
	maxSize      = 20
)

func main() {
	dbUrl := fmt.Sprintf("postgresql://%s:%s@db:5432/%s", "postgres", "testtest", "widgets")
	conn, err := pgx.Connect(context.Background(), dbUrl)
	for err != nil {
		log.Println("Failed to connect to DB, reconnecting...")
		time.Sleep(100 * time.Millisecond)
		conn, err = pgx.Connect(context.Background(), dbUrl)
	}
	defer conn.Close(context.Background())

	err = createWidgets(conn)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DbConn: conn}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func addWidget(conn *pgx.Conn, color model.WidgetColor, id, price, inventory, size int) error {
	if _, err := conn.Exec(context.Background(), "insert into widgets (id, color, size) values ($1, $2, $3)", id, string(color), size); err != nil {
		return errors.Wrap(err, "Failed to insert widget")
	}

	if _, err := conn.Exec(context.Background(), "insert into prices (widget_id, price) values ($1, $2)", id, price); err != nil {
		return errors.Wrap(err, "Failed to insert price")
	}

	if _, err := conn.Exec(context.Background(), "insert into inventory (widget_id, inventory) values ($1, $2)", id, inventory); err != nil {
		return errors.Wrap(err, "Failed to insert inventory")
	}

	return nil
}

func createWidgets(conn *pgx.Conn) error {
	log.Println("Checking for data")

	var rowCount int
	err := conn.QueryRow(context.Background(), "select id from widgets limit 1").Scan(&rowCount)
	if err == pgx.ErrNoRows {
		log.Println("Creating widgets...")

		priceRange := maxPrice - minPrice
		inventoryRange := maxInventory - minInventory
		sizeRange := maxSize - minSize

		for i := 0; i < numWidgets; i++ {
			colorIdx := rand.Intn(len(model.AllWidgetColor))
			color := model.AllWidgetColor[colorIdx]

			price := int(math.Floor(float64(priceRange)*rand.Float64()) + minPrice)
			inventory := int(math.Floor(float64(inventoryRange)*rand.Float64()) + minInventory)
			size := int(math.Floor(float64(sizeRange)*rand.Float64()) + minSize)

			if err := addWidget(conn, color, i, price, inventory, size); err != nil {
				return err
			}
		}
		log.Println("Done creating widgets")
	} else if err != nil {
		return errors.Wrap(err, "Failed to get count")
	} else {
		log.Println("Widgets already in DB")
	}

	return nil
}
