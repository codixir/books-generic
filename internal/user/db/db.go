package db

import (
	"fmt"

	"github.com/codixir/books-generic/config"
	"github.com/codixir/books-generic/internal/model"

	"database/sql"

	_ "github.com/lib/pq"
)

type (
	DB struct {
		db *sql.DB
	}

)

func NewDB(db *sql.DB) *DB {
	psqlDB := DB{
		db: db,
	}

	return &psqlDB
}

func CreateDBConnection(dbName string) (*sql.DB, error) {
	if dbName == "books" {
		dbPrefix = "BOOKS_"
	}

	if dbName == "users" {
		dbPrefix = "USERS_"
	}

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
}
