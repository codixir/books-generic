package main

import (
	"log"
	"net/http"
	"os"

	app "github.com/codixir/books-generic/internal/app"
	booksApi "github.com/codixir/books-generic/internal/book/api"
	db1 "github.com/codixir/books-generic/internal/book/db"
	usersApi "github.com/codixir/books-generic/internal/user/api"
	db2 "github.com/codixir/books-generic/internal/user/db"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn1, err := db1.CreateBooksDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	booksDb := db1.NewDB(conn1)

	conn2, err := db2.CreateUsersDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	usersDb := db2.NewDB(conn2)

	application := app.NewApplication()

	booksApi.NewApp(booksDb, application)
	usersApi.NewApp(usersDb, application)

	server := http.Server{
		Addr:    ":9000",
		Handler: application.Router,
	}

	log.Println("listening on port ", server.Addr)
	loggedRouter := handlers.LoggingHandler(os.Stdout, application.Router)
	log.Fatal(http.ListenAndServe(""+server.Addr, loggedRouter))
}
