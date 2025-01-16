package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// DB is the database connection object, which will helpe us to connect to the database

var (
	db *gorm.DB
)

//creat connent Function helps to connect to the database

func Connect() {
	d, err := gorm.Open("sqlite3", "simplerest.db")
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		panic(err)
	}
	db = d
	fmt.Println("Connected to the database successfully")
}

func GetDB() *gorm.DB {
	return db
}
