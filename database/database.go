package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"zsxyww.com/scheduler/config"
	"zsxyww.com/scheduler/model"
)

var err error

func Connect() {
	if config.DB.Type == "SQLite" {
		connectSQLite()
	} else {
		fmt.Println("sorry,we support SQLite only so far,check **DB.Type** entry in the config file")
		os.Exit(1)
	}
	Main.AutoMigrate(&model.Member{})
}

func connectSQLite() {
	Main, err = gorm.Open(sqlite.Open(config.DB.Path), &gorm.Config{})
	if err != nil {
		fmt.Printf("error in connecting to SQLite:")
		fmt.Println(err)
		os.Exit(1)

	}
}
