package repositories

import (
	"errors"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	InsertAuthor(Author) (Author, error)
	SelectAllAuthor() ([]Author, error)
	SelectAuthor(id, name string) (Author, error)
	DeleteAuthor(id string) error
}

type AuthorRepo struct {
	DB *gorm.DB
}

func NewAuthorRepo(DBConn *gorm.DB) AuthorRepository {
	return &AuthorRepo{
		DB: DBConn,
	}
}

func (db *AuthorRepo) SelectAuthor(id string, name string) (authorModel Author, err error) {

	if id != "" {
		err = db.DB.Where("id = ?", id).First(&authorModel).Error
	} else {
		err = db.DB.Where("name = ?", name).First(&authorModel).Error
	}

	return authorModel, err
}

func (db *AuthorRepo) InsertAuthor(author Author) (Author, error) {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		selectAuthor, _ := db.SelectAuthor("", author.Name)
		if selectAuthor.ID != "" {
			return errors.New("author's name already exist")
		}

		return tx.Create(&author).Error
	})

	return author, err
}

func (db *AuthorRepo) SelectAllAuthor() ([]Author, error) {
	authorsModel := []Author{}

	err := db.DB.Find(&authorsModel).Error

	return authorsModel, err
}

func (db *AuthorRepo) DeleteAuthor(id string) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		selectAuthor, _ := db.SelectAuthor(id, "")
		if selectAuthor.Name == "" {
			return errors.New("author not found")
		}

		return tx.Delete(&selectAuthor).Error
	})
	return err
}
