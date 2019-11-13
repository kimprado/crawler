package main

import (
	"log"
	"net/http"
	"os"

	"bitbucket.org/challen_ge/smartmei_backend"
	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(smartmei_backend.NewExecutableSchema(smartmei_backend.Config{Resolvers: &smartmei_backend.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
