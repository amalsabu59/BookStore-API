package config

import (
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
d, err := gorm.Open("mysql","amal:amal@123/bookstore?charset=utf8&parseTime=true&loc=Local")
if err != nil {
	panic(err)
}
db = d
}

func GetDB() *gorm.DB {
	return db
}