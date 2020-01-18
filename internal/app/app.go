package app

import (
	"net/http"

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found..."))
}

func ApplicationHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hey this is the application router"))
}
