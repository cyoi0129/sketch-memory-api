package main

import (
	"log"
	"net/http"
	"os"
	"sketch/graph"

	"sketch/db"

	"sketch/loader"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := db.ConnectDatabase()
	// loaderの初期化
	ldrs := loader.NewLoaders(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://sketch-memory-app.vercel.app", "http://localhost:3000", "http://localhost:8080", "http://localhost:5173"},
		AllowCredentials: true,
	})
	// http.Handle("/query", c.Handler(srv))
	http.Handle("/query", c.Handler(loader.Middleware(ldrs, srv)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
