package services

import (
	"basicGraphql/graph/model"
	"basicGraphql/repositories"

	"github.com/jinzhu/copier"
)

type GenreService interface {
	AddGenre(model.GenreInput) (model.Genre, error)
	GetAllGenre() ([]*model.Genre, error)
}

type GenreRepo struct {
	genreData repositories.GenreRepository
}

func NewGenreService(data repositories.GenreRepository) GenreService {
	return &GenreRepo{
		genreData: data,
	}
}

func (r *GenreRepo) AddGenre(input model.GenreInput) (model.Genre, error) {
	genreModelRepo := repositories.Genre{}
	copier.Copy(&genreModelRepo, input)

	result, err := r.genreData.InsertGenre(genreModelRepo)

	resp := model.Genre{}
	copier.Copy(&resp, result)

	return resp, err
}

func (r *GenreRepo) GetAllGenre() ([]*model.Genre, error) {

	genreRepo, err := r.genreData.SelectAllGenre()

	genreGraph := []*model.Genre{}
	copier.Copy(&genreGraph, genreRepo)

	return genreGraph, err
}
