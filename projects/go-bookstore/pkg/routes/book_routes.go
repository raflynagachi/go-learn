package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raflynagachi/go-bookstore/pkg/controller"
)

func NewRouter(c controller.BookController) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/book", c.Create).Methods(http.MethodPost)
	router.HandleFunc("/book/{bookId}", c.Update).Methods(http.MethodPut)
	router.HandleFunc("/book/{bookId}", c.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/book/{bookId}", c.FindById).Methods(http.MethodGet)
	router.HandleFunc("/book", c.FindAll).Methods(http.MethodGet)
	return router
}
