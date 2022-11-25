package factory

import (
	"basicGraphql/graph"
	Repo "basicGraphql/repositories"
	Serv "basicGraphql/services"

	"gorm.io/gorm"
)

func InitFactory(db *gorm.DB) graph.Resolver {
	bookData := Repo.NewBookRepo(db)
	bookServ := Serv.NewBookService(bookData)

	authorData := Repo.NewAuthorRepo(db)
	authorServ := Serv.NewAuthorService(authorData)

	genreData := Repo.NewGenreRepository(db)
	genreServ := Serv.NewGenreService(genreData)

	return graph.Resolver{
		BOOKSERV:   bookServ,
		AUTHORSERV: authorServ,
		GENRESERV:  genreServ,
	}
}
