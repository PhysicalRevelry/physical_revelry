package main

import (
    "log"
    "net/http"
    "os"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/rs/cors"

    "github.com/yourusername/my-app/backend/graph"
    "github.com/yourusername/my-app/backend/internal/config"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize configuration
    cfg := config.New()

    // Initialize services
    resolver := &graph.Resolver{
        Config: cfg,
    }

    // GraphQL server
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

    // Router setup
    router := mux.NewRouter()

    // GraphQL endpoints
    router.Handle("/graphql", srv)
    router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

    // CORS setup
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
        AllowedMethods: []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders: []string{"*"},
    })

    handler := c.Handler(router)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server running on http://localhost:%s", port)
    log.Printf("GraphQL playground: http://localhost:%s/playground", port)
    log.Fatal(http.ListenAndServe(":"+port, handler))
}