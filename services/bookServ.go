package services

import (
	"basicGraphql/graph/model"
	"basicGraphql/repositories"
)

type BookService interface {
	GetAllBook() (Books []*model.Book)
}

type BookRepo struct {
	BookData repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &BookRepo{
		BookData: bookRepo,
	}
}

func (r *BookRepo) GetAllBook() []*model.Book {
	print("abc")
	r.BookData.GetAllData()
	return []*model.Book{}
}
