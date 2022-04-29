package dummy

import model "movie-crud-api/model/domain"

func GenereateDummy() []model.Movie {
	var movies []model.Movie
	movies = append(movies, model.Movie{
		ID:    "00001",
		Title: "Red Dead Flowers",
		Director: &model.Director{
			Firstname: "Rafly",
			Lastname:  "Nagachi",
		},
	})
	movies = append(movies, model.Movie{
		ID:    "54322",
		Title: "Blue Dragon Town",
		Director: &model.Director{
			Firstname: "Hackie",
			Lastname:  "Nanashi",
		},
	})
	return movies
}
