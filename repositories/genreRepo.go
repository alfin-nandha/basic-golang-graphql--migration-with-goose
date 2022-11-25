package repositories

import "gorm.io/gorm"

type GenreRepository interface {
	InsertGenre(Genre) (Genre, error)
	SelectAllGenre() ([]Genre, error)
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
		return tx.Create(&genre).Error
	})

	return genre, err
}

func (db *GenreData) SelectAllGenre() ([]Genre, error) {
	genreModels := []Genre{}

	err := db.DB.Find(&genreModels).Error

	return genreModels, err
}
