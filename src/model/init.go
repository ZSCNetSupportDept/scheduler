package model

import (
	"zsxyww.com/scheduler/config"
	db "zsxyww.com/scheduler/database"
)

func init() {
	if config.InitDB == true {
		db.Main.AutoMigrate(&Member{}, &Tweak{})
	}
}
