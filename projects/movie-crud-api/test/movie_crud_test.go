package test

import (
	"encoding/json"
	"io"
	"movie-crud-api/app"
	"movie-crud-api/controller"
	"movie-crud-api/dummy"
	model "movie-crud-api/model/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMovie(t *testing.T) {
	var movies = dummy.GenereateDummy()
	var movieController = controller.NewMovieControllerImpl(movies)
	var router = app.NewRouter(movieController)

	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(`{"title":"Yellow Cars Movie","director":{"firstname":"juju","lastname":"larry"}}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/movies", requestBody)
	request.Header.Add("content-type", "application/json")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var movie model.Movie
	json.Unmarshal(body, &movie)

	assert.Equal(t, "Yellow Cars Movie", movie.Title)
	assert.Equal(t, "larry", movie.Director.Lastname)
}

func TestUpdateMovie(t *testing.T) {
	var movies = dummy.GenereateDummy()
	var movieController = controller.NewMovieControllerImpl(movies)
	var router = app.NewRouter(movieController)

	recorder := httptest.NewRecorder()
	requestBody := strings.NewReader(`{"title":"Bunny Girl","director":{"firstname":"Rigan","lastname":"Nagachi"}}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/movies/00001", requestBody)
	request.Header.Add("content-type", "application/json")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var movie model.Movie
	json.Unmarshal(body, &movie)

	assert.Equal(t, "Bunny Girl", movie.Title)
	assert.Equal(t, "Rigan", movie.Director.Firstname)
	assert.Equal(t, "Nagachi", movie.Director.Lastname)
}

func TestDeleteMovie(t *testing.T) {
	var movies = dummy.GenereateDummy()
	var movieController = controller.NewMovieControllerImpl(movies)
	var router = app.NewRouter(movieController)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/movies/00001", nil)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var movie model.Movie
	json.Unmarshal(body, &movie)

	assert.Equal(t, "Red Dead Flowers", movie.Title)
	assert.Equal(t, "Rafly", movie.Director.Firstname)
	assert.Equal(t, "Nagachi", movie.Director.Lastname)
}

func TestFindAllMovie(t *testing.T) {
	var movies = dummy.GenereateDummy()
	var movieController = controller.NewMovieControllerImpl(movies)
	var router = app.NewRouter(movieController)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/movies", nil)

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var moviesNew []model.Movie
	json.Unmarshal(body, &moviesNew)

	assert.Equal(t, "Red Dead Flowers", moviesNew[0].Title)
	assert.Equal(t, "Rafly", moviesNew[0].Director.Firstname)
	assert.Equal(t, "Nagachi", moviesNew[0].Director.Lastname)

	assert.Equal(t, "Blue Dragon Town", moviesNew[1].Title)
	assert.Equal(t, "Hackie", moviesNew[1].Director.Firstname)
	assert.Equal(t, "Nanashi", moviesNew[1].Director.Lastname)
}

func TestFindByIdMovie(t *testing.T) {
	var movies = dummy.GenereateDummy()
	var movieController = controller.NewMovieControllerImpl(movies)
	var router = app.NewRouter(movieController)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/movies/54322", nil)
	request.Header.Add("content-type", "application/json")

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var movie model.Movie
	json.Unmarshal(body, &movie)

	assert.Equal(t, "Blue Dragon Town", movie.Title)
	assert.Equal(t, "Hackie", movie.Director.Firstname)
	assert.Equal(t, "Nanashi", movie.Director.Lastname)
}
