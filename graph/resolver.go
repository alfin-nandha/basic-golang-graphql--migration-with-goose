package graph

import (
	"basicGraphql/graph/model"
	"basicGraphql/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	BOOKSTORE  []*model.Book
	BOOKSERV   services.BookService
	AUTHORSERV services.AuthorService
}
