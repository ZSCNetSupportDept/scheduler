package db

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"zsxyww.com/scheduler/config"
)

var err error

func Connect() {
	switch config.Default.DB.Type {
	case "SQLite":
		connectSQLite()
	case "PostgreSQL":
		connectPGSQL()
		PGSQL()
	default:
		panic("DBType error")
	}
}

func connectSQLite() {
	Main, err = gorm.Open(sqlite.Open(config.Default.DB.Path), &gorm.Config{})
	if err != nil {
		fmt.Printf("error in connecting to SQLite:")
		fmt.Println(err)
		os.Exit(1)
	}
}
