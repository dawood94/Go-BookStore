package controllers

import (
	"GO-BOOKSTORE/pkg/config"
	"GO-BOOKSTORE/pkg/models"
	"GO-BOOKSTORE/pkg/utils"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// var NewBook models.Book
var db = config.GetDB()

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks) //converting to json
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json") //setting the header response
	w.WriteHeader(http.StatusOK)                       // 200 everything is ok
	w.Write(res)                                       //sending the response

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                       // getting the value from request route
	bookId := vars["bookId"]                  //getting the value of id
	ID, err := strconv.ParseInt(bookId, 0, 0) //converting the string to int
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	bookDetails, db := models.GetBookById(ID) //getting the book by id
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}   //creating a book instance
	utils.ParseBody(r, CreateBook) // because we are sending the data in the body of the request we need to parse it
	b := CreateBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book := models.DeleteBook(ID)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{} // this is the requset from the user and they have Name , Author and Publication
	utils.ParseBody(r, updateBook)  // parsing the body of the request
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
	}
	bookDetails, db := models.GetBookById(ID) // we find the book by id that we want to update
	if db.Error != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
	}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name //updating the name of the book by given new name from the user

	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author //updating the author of the book by given new author from the user
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication //updating the publication of the book by given new publication from the user
	}

	db.Save(&bookDetails) //saving the updated book
	res, err := json.Marshal(bookDetails)
	if err != nil {
		http.Error(w, "Error while marshalling", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
