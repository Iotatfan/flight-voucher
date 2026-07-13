package config

import (
	"github.com/glebarez/sqlite" // Pure Go driver, no CGO required!
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("vouchers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}
