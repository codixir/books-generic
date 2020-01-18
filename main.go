package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/codixir/books-generic/internal/app"
	booksApi "github.com/codixir/books-generic/internal/book/api"
	"github.com/codixir/books-generic/internal/book/db"
	usersApi "github.com/codixir/books-generic/internal/user/api"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := db.CreateDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	db := db.NewDB(conn)

	application := app.NewApplication()

	booksApi.NewApp(db, application)
	usersApi.NewApp(db, application)

	server := http.Server{
		Addr:    ":9000",
		Handler: application.Router,
	}

	log.Println("listening on port ", server.Addr)
	loggedRouter := handlers.LoggingHandler(os.Stdout, application.Router)
	log.Fatal(http.ListenAndServe(""+server.Addr, loggedRouter))
}
