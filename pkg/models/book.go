package models

import (
	"GO-BOOKSTORE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB //database

type Book struct {
	gorm.Model // this will add ID, CreatedAt, UpdatedAt, DeletedAt fields to the model

	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect() //connecting to the database Verbindung herstellen
	db = config.GetDB()
	db.AutoMigrate(&Book{}) //creating the table . Tabelle wird automatisch erstellt/aktualisiert
}

// CRUD Functions for the Book Model
func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // this function comes from gorm , help to talk to the database and  used to check if a given object is new
	db.Create(&b)   //creating the record
	return b

}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books) //select * from books MySQL query
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
