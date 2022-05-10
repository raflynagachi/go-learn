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

func (s *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Printf("Welcome to %s\n", appConfig.AppName)

	s.InitializeDB(dbConfig)
	s.MigrateDB()
	s.InitializeRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Printf("Listening to port%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	appConfig.setupEnv()
	dbConfig.setupEnv()

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
