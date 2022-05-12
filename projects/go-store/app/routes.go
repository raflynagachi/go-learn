package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raflynagachi/go-store/controllers"
)

func (s *Server) InitializeRoutes() {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/", controllers.Home).Methods(http.MethodGet)

	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDir))
	s.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods(http.MethodGet)
}
