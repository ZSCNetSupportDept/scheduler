package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Load() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		os.Exit(1)
	}

	ListenPort = viper.GetInt("ListenPort")
	File = viper.GetString("File")
	DB.Path = viper.GetString("DB.Path")
	//DB.Port = viper.GetString("DB.Port")

	err := check()
	if err != nil {
		fmt.Println("check your config!")
		os.Exit(1)
	}

	debugprint()

}

func check() error {

	return nil

}
func debugprint() {

	fmt.Printf("ListenPort=%v", ListenPort)
	fmt.Printf("File=%v", File)
	fmt.Printf("database path : %s", DB.Path)

}
