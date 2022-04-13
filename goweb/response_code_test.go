package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest) //400
		fmt.Fprint(w, "name is empty")
	} else {
		fmt.Fprintf(w, "name is %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	ResponseCode(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)
	fmt.Println(result.StatusCode)
	fmt.Println(result.Status)
	fmt.Println(string(body))
}
