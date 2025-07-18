package db

import (
	"context"
	"fmt"
	"os"

	//"gorm.io/driver/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

func connectPGSQL() {
	Main, err = gorm.Open(postgres.Open(config.Default.DB.Path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
