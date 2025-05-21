package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"zsxyww.com/scheduler/config"
)

// use pgx to connect
func PGSQL() {
	pgx, err := pgxpool.New(context.Background(), config.Default.DB.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	version := ""
	if err := pgx.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		fmt.Printf("Query failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to:", version)
	PGX = pgx
}
