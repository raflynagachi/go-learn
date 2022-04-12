package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contetType := r.Header.Get("content-type")

	fmt.Fprint(w, contetType)
}

func TestRequestHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	request.Header.Add("content-type", "application/json")

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "RaflyNagachi")
	fmt.Fprint(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	request.Header.Add("content-type", "application/json")

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	xPoweredBy := recorder.Header().Get("x-powered-by")
	fmt.Println(xPoweredBy)
}
