package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "user=username password=password dbname=database_name host=hostname port=port sslmode=mode"
"
	d, err := gorm.Open("postgres", dsn)
if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
