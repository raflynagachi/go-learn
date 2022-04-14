package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return fmt.Sprintf("Hello %s, my name is %s", name, myPage.Name)
}

func TemplateFunc(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("Function").Parse(`{{.SayHello "gordon"}}`))
	t.ExecuteTemplate(w, "Function", MyPage{
		Name: "RaflyNagachi",
	})
}

func TestTemplateFunc(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateFunc(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateFuncGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("Function")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
		"lower": func(value string) string {
			return strings.ToLower(value)
		},
	})

	t = template.Must(t.Parse(`{{.SayHello .Name | upper}}`))
	t.ExecuteTemplate(w, "Function", MyPage{
		Name: "RaflyNagachi",
	})
}

func TestTemplateFuncGlobal(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)

	TemplateFuncGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
