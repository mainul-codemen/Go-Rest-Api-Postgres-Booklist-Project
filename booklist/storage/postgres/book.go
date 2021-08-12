package postgres

import (
	"booklist/storage"
	"log"
	"strconv"
)

const e = `
SELECT 
	book.id, 
	title,
	author,
	year
	   
FROM 
	book
`

func (s *Storage) GetBook() ([]storage.Book, error) {
	b := make([]storage.Book, 0)
	if err := s.db.Select(&b, e); err != nil {
		return nil, err
	}
	return b, nil
}

const createBookQuery = `
	INSERT INTO book(
		title,
		author,
		year
	)
	VALUES(
		:title,
		:author,
		:year
	)
	RETURNING id
	`

func (s *Storage) CreateBook(b storage.Book) (int32, error) {
	stmt, err := s.db.PrepareNamed(createBookQuery)
	if err != nil {
		return 0, err
	}
	var id int32
	if err := stmt.Get(&id, b); err != nil {
		return 0, err
	}
	return id, nil
}

const selectByIdQuery = `
SELECT 
	book.id, 
	title,
	author,
	year
	
FROM book

WHERE book.id = $1
`

func (s *Storage) GetDataById(id string) (storage.Book, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("Unable to parse String to integer")
	}
	jason := storage.Book{}
	err = s.db.Get(&jason, selectByIdQuery, i)
	return jason, err
}
