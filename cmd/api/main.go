package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
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
