package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"

	app "github.com/codixir/books-generic/internal/app"
	"github.com/codixir/books-generic/internal/book/db"
	"github.com/codixir/books-generic/internal/model"

	"github.com/gorilla/mux"
)

type (
	App struct {
		*mux.Router
		db db.BooksDB
	}
)

func NewApp(db db.BooksDB, application *app.Application) *App {
	app := App{
		Router: application.Router,
		db:     db,
	}

	app.Router.HandleFunc("/health", app.HealthCheck).Methods("GET")

	for _, r := range app.GetRoutes() {
		app.Router.HandleFunc(r.Pattern, r.HandlerFunc).Methods(r.Method)
	}

	return &app
}

func (app *App) GetRoutes() []model.Route {
	return []model.Route{
		{
			Method:      "GET",
			Pattern:     "/books",
			HandlerFunc: app.GetBooks,
		},
		{
			Method:      "GET",
			Pattern:     "/books/{id}",
			HandlerFunc: app.GetBook,
		},
		{
			Method:      "POST",
			Pattern:     "/books",
			HandlerFunc: app.CreateBook,
		},
	}
}

func (app *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := app.db.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError) //500
		return
	}

	w.WriteHeader(http.StatusOK) //200
}

func (app *App) GetBooks(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var limit int
	var offset int

	if params["limit"] != nil {
		limit, _ = strconv.Atoi(params["limit"][0])
	} else {
		limit = 1000000
	}

	if params["offset"] != nil {
		offset, _ = strconv.Atoi(params["offset"][0])
	}

	pg := model.Pagination{
		Limit:  limit,
		Offset: offset,
	}

	books, err := app.db.GetBooks(pg)
	if err != nil {
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}

	res := model.PaginatedResponse{
		Data:       books,
		Pagination: &pg,
	}

	respondWithSuccess(res, w)
}

func (app *App) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	book, err := app.db.GetBook(id)
	if err != nil {
		respondWithError(err, http.StatusNotFound, w)
		return
	}

	res := model.ResponseSingleBody{
		Data: &model.Item{
			Value: book,
		},
	}

	respondWithSuccess(res, w)
}

func (app *App) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	now := time.Now()

	json.NewDecoder(r.Body).Decode(&book)

	if book.Title == "" {
		err := errors.New("failed to create book")
		respondWithError(err, http.StatusBadRequest, w)
		return
	}

	book.ID = uuid.New().String()
	book.CreatedAt = now
	book.UpdatedAt = now

	result, err := app.db.CreateBook(book)

	if err != nil {
		respondWithError(err, http.StatusInternalServerError, w)
		return
	}

	res := model.ResponseSingleBody{
		Data: &model.Item{
			Value: result,
		},
	}

	w.WriteHeader(http.StatusCreated)
	respondWithSuccess(res, w)
}

func respondWithError(err error, statusCode int, w http.ResponseWriter) string {
	body := model.ErrorBody{
		Error:  err.Error(),
		Status: statusCode,
	}

	content, _ := json.Marshal(body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(content)
	return string(content)
}

func respondWithSuccess(response interface{}, w http.ResponseWriter) string {
	content, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(content)

	return string(content)
}
