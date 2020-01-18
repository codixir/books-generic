package api

import (
	"net/http"

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
			Pattern:     "/users",
			HandlerFunc: app.GetUsers,
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

func (app *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("abcd----books"))
}
