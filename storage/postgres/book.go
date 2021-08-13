package postgres

import (
	"Go-Rest-Api-Postgres-Booklist-Project/storage"
	"database/sql"
	"fmt"
	"log"

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

func (s *Storage) GetDataById(id int64) (storage.Book, error) {
	jason := storage.Book{}
	err := s.db.Get(&jason, selectByIdQuery, id)
	return jason, err
}

const update = `
	UPDATE 
		book 
	SET 

		title=:title,
		author=:author,
		year=:year

	WHERE 
		id =: id`

func (s *Storage) UpdateBook(b storage.Book) (sql.Result,error){
	res, err := s.db.NamedExec(update, b)
	log.Println("----------------------------------> ",res)
	if err != nil {
		log.Println("Unable to update data")
	}

	fmt.Println("------>", res)
	return res,err
}

func (s *Storage) DeleteDataById(id int) (int64, error) {
	fmt.Println("what is the id ===  ??? === ",id)
	result, err := s.db.Exec("delete from book where id = $1", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}