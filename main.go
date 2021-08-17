package main

import (
	"Go-Rest-Api-Postgres-Booklist-Project/handler"
	"Go-Rest-Api-Postgres-Booklist-Project/storage/postgres"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Welcome ! Be patient ... it is Connecting")
	dbString := newDBFromConfig()
	store, err := postgres.NewStorage(dbString)
	if err != nil {
		log.Fatal(err)
	}

	r, err := handler.NewServer(store)
	if err != nil {
		log.Fatal("Handler not found.")
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

// func newDBFromConfig() string {
// 	dbParams := " " + "user=postgres"
// 	dbParams += " " + "host=postgres"
// 	dbParams += " " + "port=5432"
// 	dbParams += " " + "dbname=booklist"
// 	dbParams += " " + "password=password"
// 	dbParams += " " + "sslmode=disable"

// 	return dbParams
// }

func newDBFromConfig() string {
	dbParams := " " + "user=postgres"
	dbParams += " " + "host=localhost"
	dbParams += " " + "port=5432"
	dbParams += " " + "dbname=booklist"
	dbParams += " " + "password=password"
	dbParams += " " + "sslmode=disable"

	return dbParams
}
