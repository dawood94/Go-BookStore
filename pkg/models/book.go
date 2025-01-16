package models

import (
	"GO-BOOKSTORE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model // t"his will add ID, CreatedAt, UpdatedAt, DeletedAt fields to the model

	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect() //connecting to the database Verbinfung herstellen
	db = config.GetDB()
	db.AutoMigrate(&Book{}) //creating the table
}
