package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize() {
	fmt.Println("Welcome to Go-Store")
	s.Router = mux.NewRouter()

	s.InitializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Printf("Listening to port:%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

func Run() {
	var server = Server{}

	server.Initialize()
	server.Run(":8080")
}
