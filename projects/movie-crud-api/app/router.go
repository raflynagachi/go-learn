package app

import (
	"movie-crud-api/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(c controller.MovieController) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", c.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/movies/{id}", c.FindById).Methods(http.MethodGet)
	router.HandleFunc("/movies", c.Create).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", c.Update).Methods(http.MethodPut)
	router.HandleFunc("/movies/{id}", c.Delete).Methods(http.MethodDelete)
	return router
}
