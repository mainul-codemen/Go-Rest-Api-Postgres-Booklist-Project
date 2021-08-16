# Go-Rest-Api-Postgres-Booklist-Project

### DockerFile

docker container ls

docker image ls

docker build -t go-rest-api-booklist

docker run -8080:8000 go-rest-api-booklist

docker run --name postgres12 -p 5432:5432 -e POSTGRES_PASSWORD=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

```yml
FROM golang:latest

LABEL maintainer="Mainul Hasan"

WORKDIR /app

COPY go.mod

COPY go.sum

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build

CMD["/.Go-Rest-Api-Postgres-Booklist-Project"]
```

## All in one file :

```go
package main

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

)

type Book struct {
    ID int `json:id`
    Title string `json:title`
    Author string `json:author`
    Year string `json:year`
}

var books []Book

func main() {
r := mux.NewRouter()

    books = append(books,

    	Book{ID: 1, Title: "Go lang book", Author: "google", Year: "2021"},

    	Book{ID: 2, Title: "Java program", Author: "Ordinanry writer", Year: "2021"},

    	Book{ID: 3, Title: "C programming", Author: "Denish Ritche", Year: "1990"},
    )

    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books", addBook).Methods("POST")
    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", removeBook).Methods("DELETE")


    http.ListenAndServe(":8000", r)

}

func getBooks(w http.ResponseWriter, r \*http.Request) {
    json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r \*http.Request) {
    params := mux.Vars(r)
    // convert stirng to int
    i, _ := strconv.Atoi(params["id"])
    for _, book := range books {

            if book.ID == i {
                json.NewEncoder(w).Encode(&book)
            }
        }

}

func addBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    // decode data from browser
    _ = json.NewDecoder(r.Body).Decode(&book)

    books = append(books, book)
    json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r \*http.Request) {
    var book Book

    json.NewDecoder(r.Body).Decode(&book)
    for i,item := range books{
    	if item.ID == book.ID{
    		books[i] = book
    	}
    }

    json.NewEncoder(w).Encode(books)

}

func removeBook(w http.ResponseWriter, r \*http.Request) {
    params := mux.Vars(r)
    id,_ := strconv.Atoi(params["id"])
    for i, item := range books{
    if item.ID == id {
        books = append(books[:i],books[i+1:]... )
    }
}

    json.NewEncoder(w).Encode(books)

}

```
