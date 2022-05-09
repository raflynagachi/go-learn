package app

import (
	"net/http"

	"github.com/raflynagachi/go-store/controllers"
)

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/", controllers.Home).Methods(http.MethodGet)
}
