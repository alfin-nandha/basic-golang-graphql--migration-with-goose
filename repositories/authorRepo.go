package repositories

import (
	"gorm.io/gorm"
)

type AuthorRepository interface {
	InsertAuthor(Author) (Author, error)
	SelectAllAuthor() ([]Author, error)
}

type AuthorRepo struct {
	DB *gorm.DB
}

func NewAuthorRepo(DBConn *gorm.DB) AuthorRepository {
	return &AuthorRepo{
		DB: DBConn,
	}
}

func (db *AuthorRepo) InsertAuthor(author Author) (Author, error) {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&author).Error
	})

	return author, err
}

func (db *AuthorRepo) SelectAllAuthor() ([]Author, error) {
	authorsModel := []Author{}

	err := db.DB.Find(&authorsModel).Error

	return authorsModel, err
}
