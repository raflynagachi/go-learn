package controller

import (
	"encoding/json"
	"math/rand"
	model "movie-crud-api/model/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MovieControllerImpl struct {
	Movies []model.Movie
}

func NewMovieControllerImpl(movies []model.Movie) MovieController {
	return &MovieControllerImpl{
		Movies: movies,
	}
}

func (m *MovieControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie model.Movie
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000))
	m.Movies = append(m.Movies, movie)

	encoder := json.NewEncoder(w)
	encoder.Encode(movie)
}

func (m *MovieControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, movie := range m.Movies {
		if movie.ID == params["id"] {
			m.Movies = append(m.Movies[:index], m.Movies[index+1:]...)

			var movieNew model.Movie
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&movieNew)
			movieNew.ID = params["id"]

			m.Movies = append(m.Movies, movie)
			encoder := json.NewEncoder(w)
			encoder.Encode(movieNew)
			return
		}
	}
}

func (m *MovieControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, movie := range m.Movies {
		if movie.ID == params["id"] {
			m.Movies = append(m.Movies[:index], m.Movies[index+1:]...)
			encoder := json.NewEncoder(w)
			encoder.Encode(movie)
			return
		}
	}
}

func (m *MovieControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, movie := range m.Movies {
		if movie.ID == params["id"] {
			encoder := json.NewEncoder(w)
			encoder.Encode(movie)
			return
		}
	}
}

func (m *MovieControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(m.Movies)
}
