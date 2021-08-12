package handler

import (
	"booklist/storage/postgres"

	"github.com/gorilla/mux"
)


type (
	Server struct{
		store *postgres.Storage
	}
)

func NewServer(st *postgres.Storage)(*mux.Router,error){
	s := &Server{
		store: st,
	}

	r := mux.NewRouter()


	r.HandleFunc("/books", s.getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", s.getBook).Methods("GET")
	r.HandleFunc("/books", s.addBook).Methods("POST")
	r.HandleFunc("/books", s.getBooks).Methods("GET")
	r.HandleFunc("/books", s.updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", s.removeBook).Methods("DELETE")

	return r,nil
}