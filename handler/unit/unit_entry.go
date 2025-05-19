package uo

import (
	"gorm.io/gorm"
	"zsxyww.com/scheduler/database"
)

// Unit Operations
type uoPrototype struct {
	c *gorm.DB
}

func init() {
	Uo := &uoPrototype{c: db.Main}
	_ = Uo
}
