package handler

import (
	"gorm.io/gorm"
	"zsxyww.com/scheduler/database"
)

// Unit Operations
type uoPrototype struct {
	c *gorm.DB
}

func init() {
	uo := uoPrototype{c: db.Main}
	_ = uo
}
