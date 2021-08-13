package handler

import (
	"Go-Rest-Api-Postgres-Booklist-Project/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// type BookForm struct{
// 	ID     int
// 	Title  string
// 	Author string
// 	Year   string
// }


func (s *Server)getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get All Book Called....")
	bk,err := s.store.GetBook()
	if err != nil {
		log.Println("Unable to find data")
	}
	json.NewEncoder(w).Encode(bk)
	
}

func (s *Server)getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Book By Id Called....")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	bk,err := s.store.GetDataById(int64(i))
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
	json.NewEncoder(w).Encode(bk)
}

// Add book to the database
func (s *Server)addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Create Book called....")
	var book storage.Book
	
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil{
		log.Fatalln("Decoding Erorr")
	}

	bookId,err := s.store.CreateBook(book)
	if err != nil {
		log.Fatal("Unable to Store data.",err)
	}

	json.NewEncoder(w).Encode(bookId)	
}

func (s *Server)updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update Book Called....")
	var book storage.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil{
		log.Fatalln("Decoding Error",err)
	}

	res,_ := s.store.UpdateBook(book)
	rowup ,err := res.RowsAffected()
	
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(rowup)
}

func (s *Server)removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("remove Book Called....")
	params := mux.Vars(r)
	i, _ := strconv.Atoi(params["id"])
	log.Println("Remove books of id = ",i)
	rowsAffected,err := s.store.DeleteDataById(i)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}else{
		fmt.Println("Rows Affected:",rowsAffected)
	}

	json.NewEncoder(w).Encode(rowsAffected)
}


func (s *Server)error(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("Data is not present")

}
