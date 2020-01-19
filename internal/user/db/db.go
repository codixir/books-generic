package db

import (
	"fmt"

	"github.com/codixir/books-generic/config"

	"database/sql"

	_ "github.com/lib/pq"
)

const dbPrefix = "USERS_"

type (
	DB struct {
		db *sql.DB
	}

	UsersDB interface {
		Ping() error
		// GetUser() (model.User, error)
	}
)

func NewDB(db *sql.DB) *DB {
	psqlDB := DB{
		db: db,
	}

	return &psqlDB
}

func CreateUsersDbConnection() (*sql.DB, error) {
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
