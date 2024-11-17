package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Load() {
	// where to read config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	readconfig()

	check()

	debugprint()

	if len(os.Args) != 1 {
		handleArguments()
	}

}

func check() {
	// 暂时只支持SQLite
	if DB.Type != "SQLite" {
		fmt.Println("sorry,we support SQLite only so far(At config/config.go : check())")
		os.Exit(1)
	}

}
func debugprint() {

	fmt.Printf("ListenPort=%v\n", ListenPort)
	fmt.Printf("database type:%s\n", DB.Type)
	fmt.Printf("database path : %s\n", DB.Path)
	fmt.Printf("database port:%d\n", DB.Port)
	fmt.Printf("database user:%s\n", DB.User)
	fmt.Printf("database passowrd:%s\n", DB.Password)
	fmt.Printf("database name:%s\n", DB.Name)
	fmt.Printf("session:%s\n", Session)
	fmt.Printf("semester:%d\n", Semester)
	fmt.Printf("start time:%s\n", StartTime)
	fmt.Printf("week:%d\n", Week)
	fmt.Printf("File=%v\n", File)

}

func handleArguments() {
	if len(os.Args) > 2 {
		fmt.Println("Please enter only 1 argument")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "newsemester":
		if DB.Type == "SQLite" {
			sqliteNewSemester()
		}
	default:
		panic("invalid argument")
	}

}

func readconfig() {
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}
	ListenPort = viper.GetInt("ListenPort")
	DB.Type = viper.GetString("DB.Type")
	DB.Path = viper.GetString("DB.Path")
	DB.Port = viper.GetInt("DB.Port")
	DB.User = viper.GetString("DB.User")
	DB.Password = viper.GetString("DB.Password")
	DB.Name = viper.GetString("DB.Name")
	Session = viper.GetString("Session")
	Semester = viper.GetInt("Semester")
	StartTime = viper.GetString("StartTime")
	Week = viper.GetInt("Week")
	File = viper.GetString("File")
}
