package goweb

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/if.go.html"))
	t.ExecuteTemplate(w, "if.go.html", Page{
		Title: "Selamat Pagi",
		Name: Name{
			Firstname: "Rafly",
		},
	})
}

func TestTemplateActionIf(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateActionIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateActionIfComp(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/comparator.go.html"))
	t.ExecuteTemplate(w, "comparator.go.html", map[string]interface{}{
		"Score": 70,
	})
}

func TestTemplateActionIfComp(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateActionIfComp(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/range.go.html"))
	t.ExecuteTemplate(w, "range.go.html", map[string]interface{}{
		"score": map[string]int{
			"Fisika":     70,
			"Kimia":      65,
			"Matematika": 89,
			"Biologi":    78,
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateActionRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/with.go.html"))
	t.ExecuteTemplate(w, "with.go.html", map[string]interface{}{
		"Name": map[string]string{
			"firstname": "Rafly",
			"lastname":  "Nagachi",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateActionWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
