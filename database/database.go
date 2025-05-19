package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/model"
)

var err error

func Connect() {
	switch config.DB.Type {
	case "SQLite":
		connectSQLite()
	case "PostgreSQL":
		connectPGSQL()
	default:
		panic("DBType error")
	}
	Main.AutoMigrate(&model.Member{}, &model.Tweak{})
}

func connectSQLite() {
	Main, err = gorm.Open(sqlite.Open(config.DB.Path), &gorm.Config{})
	if err != nil {
		fmt.Printf("error in connecting to SQLite:")
		fmt.Println(err)
		os.Exit(1)

	}
}

func connectPGSQL() {
	Main, err = gorm.Open(postgres.Open(config.DB.Path))
	if err != nil {
		panic(err)
	}
}
