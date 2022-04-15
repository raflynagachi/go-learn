package httprouterlearn

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func httprouterlearnFirst() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Router")
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
