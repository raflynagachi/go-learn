package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raflynagachi/go-store/controllers"
)

func (s *Server) InitializeRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", controllers.Home).Methods(http.MethodGet)
}
