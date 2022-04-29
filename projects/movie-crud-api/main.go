package main

import (
	"fmt"
	"log"
	"movie-crud-api/app"
	"movie-crud-api/controller"
	"movie-crud-api/dummy"
	"net/http"
)

func main() {
	var movies = dummy.GenereateDummy()

	movieController := controller.NewMovieControllerImpl(movies)
	router := app.NewRouter(movieController)

	fmt.Printf("Starting server at port 8080")
	server := http.ListenAndServe("localhost:8080", router)
	log.Fatal(server)
}
