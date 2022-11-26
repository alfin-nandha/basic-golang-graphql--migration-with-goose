package repositories

import (
	"errors"

	"gorm.io/gorm"
)

type GenreRepository interface {
	InsertGenre(Genre) (Genre, error)
	SelectAllGenre() ([]Genre, error)
	SelectGenre(id string, name string) (Genre, error)
	DeleteGenre(string) error
}

type GenreData struct {
	DB *gorm.DB
}

func NewGenreRepository(db *gorm.DB) GenreRepository {
	return &GenreData{
		DB: db,
	}
}

func (db *GenreData) InsertGenre(genre Genre) (Genre, error) {

	err := db.DB.Transaction(func(tx *gorm.DB) error {
		selectGenre, _ := db.SelectGenre("", genre.Name)
		if selectGenre.Name != "" {
			return errors.New("genre's name already exist")
		}

		return tx.Create(&genre).Error
	})

	return genre, err
}

func (db *GenreData) SelectAllGenre() ([]Genre, error) {
	genreModels := []Genre{}

	err := db.DB.Find(&genreModels).Error

	return genreModels, err
}

func (db *GenreData) SelectGenre(id, name string) (genreModel Genre, err error) {
	if id != "" {
		err = db.DB.Where("id = ?", id).First(&genreModel).Error
	} else {
		err = db.DB.Where("name = ?", name).First(&genreModel).Error
	}

	return genreModel, err
}

func (db *GenreData) DeleteGenre(id string) (err error) {

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		selectGenre, _ := db.SelectGenre(id, "")
		if selectGenre.Name == "" {
			return errors.New("genre not found")
		}
		return db.DB.Delete(&selectGenre).Error
	})
	return err
}
