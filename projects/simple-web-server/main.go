package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed static
var resources embed.FS

func main() {
	router := httprouter.New()
	dir, _ := fs.Sub(resources, "static")
	router.ServeFiles("/static/*filepath", http.FS(dir))

	router.GET("/hello", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})
	router.POST("/form", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Parse error %v", err)
		}
		name := r.FormValue("name")
		addr := r.FormValue("address")
		fmt.Fprintf(w, "Hello %s from %s", name, addr)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
