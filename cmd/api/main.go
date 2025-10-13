package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
	"github.com/tolgafiratoglu/lexi/internal/infrastructure/database"
)

func main() {
	db, dsn, err := database.Connect()
    if err != nil {
        log.Fatalf("❌ DB connection error: %v", err)
    }
    defer db.Close()

    if err := database.RunMigrations(dsn); err != nil {
        log.Fatalf("❌ Migration error: %v", err)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "service is up and running")
    })

    fmt.Printf("API service is up and running on port %s\n", port)

    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}
