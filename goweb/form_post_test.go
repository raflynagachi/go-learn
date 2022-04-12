package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")

	// another way without ParseForm explicitly
	firstname2 := r.PostFormValue("firstname")

	fmt.Fprintf(w, "%s %s (%s)", firstname, lastname, firstname2)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=rafly&lastname=nagachi")
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "localhost:8080", requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
