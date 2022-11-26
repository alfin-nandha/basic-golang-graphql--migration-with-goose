package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"basicGraphql/graph/generated"
	"basicGraphql/graph/model"
	"context"
	"fmt"
)

// AddAuthor is the resolver for the add_author field.
func (r *mutationResolver) AddAuthor(ctx context.Context, input model.AuthorInput) (*model.PostStatus, error) {
	author, err := r.AUTHORSERV.AddAuthor(input)

	des := "success add author"
	response := model.PostStatus{
		Iserror:     false,
		Description: &des,
		ID:          author.AuthorID,
	}

	if err != nil {
		des = "fail add author"
		response = model.PostStatus{
			Iserror:     true,
			Description: &des,
		}
	}

	return &response, err
}

// UpdateAuthor is the resolver for the update_author field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, input *model.UpdateAuthorInput) (*model.PutStatus, error) {
	panic(fmt.Errorf("not implemented: UpdateAuthor - update_author"))
}

// DeleteAuthor is the resolver for the delete_author field.
func (r *mutationResolver) DeleteAuthor(ctx context.Context, authorID string) (*model.DeleteStatus, error) {
	err := r.AUTHORSERV.DeleteAuthor(authorID)

	des := "success delete author"
	response := model.DeleteStatus{
		Iserror:     false,
		Description: &des,
	}

	if err != nil {
		des = "fail delete author"
		response = model.DeleteStatus{
			Iserror:     true,
			Description: &des,
		}
	}
	return &response, err
}

// AddGenre is the resolver for the add_genre field.
func (r *mutationResolver) AddGenre(ctx context.Context, input model.GenreInput) (*model.PostStatus, error) {
	result, err := r.GENRESERV.AddGenre(input)

	des := "success add genre"
	response := model.PostStatus{
		Iserror:     false,
		Description: &des,
		ID:          result.GenreID,
	}
	if err != nil {
		des := "fail add genre"
		response = model.PostStatus{
			Iserror:     true,
			Description: &des,
		}
	}
	return &response, err
}

// UpdateGenre is the resolver for the update_genre field.
func (r *mutationResolver) UpdateGenre(ctx context.Context, input *model.UpdateGenreInput) (*model.PutStatus, error) {
	panic(fmt.Errorf("not implemented: UpdateGenre - update_genre"))
}

// DeleteGenre is the resolver for the delete_genre field.
func (r *mutationResolver) DeleteGenre(ctx context.Context, genreID string) (*model.DeleteStatus, error) {
	err := r.GENRESERV.DeleteGenre(genreID)
	des := "success delete genre"
	response := model.DeleteStatus{
		Iserror:     false,
		Description: &des,
	}

	if err != nil {
		des = "fail delete genre"
		response = model.DeleteStatus{
			Iserror:     true,
			Description: &des,
		}
	}
	return &response, err
}

// AddBook is the resolver for the add_book field.
func (r *mutationResolver) AddBook(ctx context.Context, input model.BookInput) (*model.PostStatus, error) {
	panic(fmt.Errorf("not implemented: AddBook - add_book"))
}

// UpdateBook is the resolver for the update_book field.
func (r *mutationResolver) UpdateBook(ctx context.Context, input *model.UpdateBookInput) (*model.PutStatus, error) {
	panic(fmt.Errorf("not implemented: UpdateBook - update_book"))
}

// DeleteBook is the resolver for the delete_book field.
func (r *mutationResolver) DeleteBook(ctx context.Context, bookID string) (*model.DeleteStatus, error) {
	panic(fmt.Errorf("not implemented: DeleteBook - delete_book"))
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.GetBook, error) {
	panic(fmt.Errorf("not implemented: Book - book"))
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented: Books - books"))
}

// Genre is the resolver for the genre field.
func (r *queryResolver) Genre(ctx context.Context, genreID string) (*model.GetGenre, error) {
	panic(fmt.Errorf("not implemented: Genre - genre"))
}

// Genres is the resolver for the genres field.
func (r *queryResolver) Genres(ctx context.Context) ([]*model.Genre, error) {
	return r.GENRESERV.GetAllGenre()
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, authorID string) (*model.GetAuthor, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return r.AUTHORSERV.GetAllAuthor()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
