package db

import (
	"fmt"

	"github.com/codixir/books-generic/config"
	"github.com/codixir/books-generic/internal/model"

	"database/sql"

	_ "github.com/lib/pq"
)

const dbPrefix = "BOOKS_"

type (
	DB struct {
		db *sql.DB
	}

	BooksDB interface {
		Ping() error
		GetBooks(pagination model.Pagination) ([]model.Book, error)
		GetBook(id string) (model.Book, error)
		CreateBook(book model.Book) (string, error)
		// UpdateBook(book Book) (int64, error)
		// DeleteBook(id int64) (int64, error)
	}
)

func NewDB(db *sql.DB) *DB {
	psqlDB := DB{
		db: db,
	}

	return &psqlDB
}

func CreateDBConnection() (*sql.DB, error) {
	Global := config.SetConfig(dbPrefix)

	source := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		Global.DBHost,
		Global.DBPort,
		Global.Database,
		Global.DBUser,
		Global.DBPassword)

	d, err := sql.Open("postgres", source)

	if err != nil {
		return nil, err
	}

	return d, nil
}

func (p *DB) Ping() error {
	return p.db.Ping()
}

func (p *DB) GetBooks(pg model.Pagination) ([]model.Book, error) {
	var book model.Book
	var books []model.Book

	stmt, err := p.db.Prepare("SELECT * FROM books offset $1 limit $2")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(pg.Offset, pg.Limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&book.ID, &book.Title, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (p *DB) GetBook(id string) (model.Book, error) {
	var book model.Book

	stmt, err := p.db.Prepare("SELECT * FROM books WHERE id = $1")
	if err != nil {
		return book, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&book.ID, &book.Title, &book.CreatedAt, &book.UpdatedAt)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (p *DB) CreateBook(book model.Book) (string, error) {
	var id string

	stmt, err := p.db.Prepare("INSERT INTO books(id, title, created_at, updated_at) values($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return "", err
	}

	err = stmt.QueryRow(book.ID, book.Title, book.CreatedAt, book.UpdatedAt).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
