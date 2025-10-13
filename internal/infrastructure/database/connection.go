package database

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, string, error) {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    name := os.Getenv("DB_NAME")
    sslmode := os.Getenv("DB_SSLMODE")

    if sslmode == "" {
        sslmode = "disable"
    }

    // PostgreSQL DSN
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        user, pass, host, port, name, sslmode)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, "", fmt.Errorf("failed to open DB: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, "", fmt.Errorf("failed to connect DB: %w", err)
    }

    fmt.Println("✅ Connected to PostgreSQL")
    return db, dsn, nil
}
