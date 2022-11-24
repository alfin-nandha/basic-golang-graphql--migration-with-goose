package services

import "basicGraphql/repositories"

type AuthorService interface {
	AddAuthor() error
}

type AuthorRepo struct {
	AuthorData repositories.AuthorRepository
}

func NewAuthorService(authorRepo repositories.AuthorRepository) AuthorService {
	return &AuthorRepo{
		AuthorData: authorRepo,
	}
}

func (r *AuthorRepo) AddAuthor() error {
	return nil
}
