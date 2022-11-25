package services

import (
	"basicGraphql/graph/model"
	"basicGraphql/repositories"

	"github.com/jinzhu/copier"
)

type AuthorService interface {
	AddAuthor(model.AuthorInput) (model.Author, error)
	GetAllAuthor() ([]*model.Author, error)
}

type AuthorRepo struct {
	AuthorData repositories.AuthorRepository
}

func NewAuthorService(authorRepo repositories.AuthorRepository) AuthorService {
	return &AuthorRepo{
		AuthorData: authorRepo,
	}
}

func (r *AuthorRepo) AddAuthor(input model.AuthorInput) (result model.Author, err error) {
	authorRepo := repositories.Author{}
	copier.Copy(&authorRepo, input)

	authorRepo, err = r.AuthorData.InsertAuthor(authorRepo)

	copier.Copy(&result, authorRepo)

	return result, err
}

func (r *AuthorRepo) GetAllAuthor() (result []*model.Author, err error) {

	authorsRepo, err := r.AuthorData.SelectAllAuthor()

	copier.Copy(&result, authorsRepo)
	return result, err
}
