package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"urlshort"
	"urlshort/helper"
	"urlshort/model"

	"github.com/stretchr/testify/assert"
)

func redirectLinkTest(t *testing.T, path string, url string, handleFunc http.HandlerFunc) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080"+path, nil)

	handleFunc(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 302, response.StatusCode)
	assert.Equal(t, url, response.Header.Get("location"))
}

func TestMapRedirect(t *testing.T) {
	pathToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandlerFunc := urlshort.MapHandler(pathToUrls, helper.DefaultMux())
	for path, url := range pathToUrls {
		redirectLinkTest(t, path, url, mapHandlerFunc)
	}
}

func TestYamlRedirect(t *testing.T) {
	pathToUrlsYaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	yamlHandlerFunc, err := urlshort.YAMLHandler([]byte(pathToUrlsYaml), helper.DefaultMux())
	if err != nil {
		panic(err)
	}

	var pathToUrls []model.PathToUrl
	helper.ParseYaml([]byte(pathToUrlsYaml), &pathToUrls)

	for _, pathToUrl := range pathToUrls {
		redirectLinkTest(t, pathToUrl.Path, pathToUrl.URL, yamlHandlerFunc)
	}
}

func TestJsonRedirect(t *testing.T) {
	pathToUrlsJson := `[
		{"path":"/portfolio", "url":"https://raflynagachi.vercel.app"},
		{"path":"/google", "url":"https://google.com"}
	]`
	jsonHandlerFunc, err := urlshort.JSONHandler([]byte(pathToUrlsJson), helper.DefaultMux())
	if err != nil {
		panic(err)
	}

	var pathToUrls []model.PathToUrl
	helper.ParseYaml([]byte(pathToUrlsJson), &pathToUrls)

	for _, pathToUrl := range pathToUrls {
		redirectLinkTest(t, pathToUrl.Path, pathToUrl.URL, jsonHandlerFunc)
	}
}
