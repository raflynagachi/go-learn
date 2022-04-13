package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-RN-Name"
	cookie.Value = r.URL.Query().Get("name")

	// all path enabled this cookie
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-RN-Name")
	if err != nil {
		fmt.Fprint(w, "no cookie")
	} else {
		name := cookie.Value
		fmt.Fprint(w, name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/?name=nagachi", nil)

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-RN-Name"
	cookie.Value = "Nagachi"
	request.AddCookie(cookie)

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	GetCookie(recorder, request)
}
