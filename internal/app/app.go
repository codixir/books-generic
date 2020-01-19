package app

import (
	"fmt"
	"net/http"

	utils "github.com/codixir/books-generic/utils"
	"github.com/gorilla/mux"
)

type (
	Application struct {
		*mux.Router
	}
)

func NewApplication() *Application {
	applicationRouter := mux.NewRouter().StrictSlash(true)

	application := Application{
		Router: applicationRouter,
	}

	applicationRouter.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	applicationRouter.HandleFunc("/app", ApplicationHandler).Methods("GET")

	return &application
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(fmt.Errorf("Not Found"), http.StatusNotFound, w)
}

func ApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey this is the application router"))
}
