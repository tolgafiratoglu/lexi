package database

import (
    "fmt"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dsn string) error {
    m, err := migrate.New(
        "file://internal/infrastructure/database/migrations",
        dsn,
    )
    if err != nil {
        return fmt.Errorf("migration init error: %w", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("migration failed: %w", err)
    }

    fmt.Println("✅ Migrations applied successfully")
    return nil
}
