package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Template from string
func SimpleTemplate(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleTemplate(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	SimpleTemplate(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// Template from files
func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/simple.go.html"))
	t.ExecuteTemplate(w, "simple.go.html", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	SimpleHTMLFile(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("template/*.go.html"))
	t.ExecuteTemplate(w, "simple.go.html", "Hello HTML Tempate")
}

func TestTemplateDirectory(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	TemplateDirectory(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

//go:embed template/*.go.html
var templates embed.FS

func TemplateFS(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templates, "template/*.go.html"))
	t.ExecuteTemplate(w, "simple.go.html", "Hello HTML Tempate")
}

func TestTemplateFS(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	TemplateFS(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// Template data map
func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/name.go.html"))
	t.ExecuteTemplate(w, "name.go.html", map[string]interface{}{
		"Title": "Template data map",
		"Name":  map[string]string{"Firstname": "Rafly"},
	})
}

func TestTemplateDataMap(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

// Template data struct
type Name struct {
	Firstname string
}

type Page struct {
	Title string
	Name  Name
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/name.go.html"))
	t.ExecuteTemplate(w, "name.go.html", Page{
		Title: "Template data struct",
		Name: Name{
			Firstname: "Rafly",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
