package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server)getBooks(w http.ResponseWriter, r *http.Request) {
	bk,err := s.store.GetBook()
	if err != nil {
		log.Println("Unable to find data")
	}
	json.NewEncoder(w).Encode(bk)
	
}

func (s *Server)getBook(w http.ResponseWriter, r *http.Request) {

}

func (s *Server)addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Speaker called....")
	

}

func (s *Server)updateBook(w http.ResponseWriter, r *http.Request) {

}

func (s *Server)removeBook(w http.ResponseWriter, r *http.Request) {
	
}
