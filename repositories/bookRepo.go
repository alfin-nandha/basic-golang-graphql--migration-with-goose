package repositories

import (
	"basicGraphql/graph/model"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllData() []*model.Book
}

type BookRepo struct {
	DB *gorm.DB
}

func NewBookRepo(DBConn *gorm.DB) BookRepository {
	return &BookRepo{
		DB: DBConn,
	}
}

func (db *BookRepo) GetAllData() []*model.Book {

	return []*model.Book{}
}
