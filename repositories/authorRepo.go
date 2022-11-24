package repositories

import "gorm.io/gorm"

type AuthorRepository interface {
	InsertAuthor() error
}

type AuthorRepo struct {
	DB *gorm.DB
}

func NewAuthorRepo(DBConn *gorm.DB) AuthorRepository {
	return &AuthorRepo{
		DB: DBConn,
	}
}

func (db *AuthorRepo) InsertAuthor() error {
	return nil
}
