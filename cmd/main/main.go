package main

import (
	"log"
	"net/http"

	"GO-BOOKSTORE/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r)) // Auto-erkannt, dass es auf allen Schnittstellen hören soll , da es nicht mit localhost:9010 funktioniert hat !

}
