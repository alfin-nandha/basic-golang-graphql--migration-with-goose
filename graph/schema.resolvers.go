package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"basicGraphql/graph/generated"
	"basicGraphql/graph/model"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
)

// AddBook is the resolver for the add_book field.
func (r *mutationResolver) AddBook(ctx context.Context, input model.BookInput) (*model.PostStatus, error) {
	log.Println("Adding Book to BookStore")
	a := input.Authors
	authorList := []*model.Author{}
	for _, auth := range a {
		authorList = append(authorList, &model.Author{
			Name: auth.Name,
			Mail: auth.Mail,
		})
	}
	bookID, err := nextBookID(r.BOOKSTORE)
	if err != nil {
		fmt.Println(err)
		des := "Internal error"
		return &model.PostStatus{
			Iserror:     true,
			Description: &des,
		}, nil
	}
	gBook := model.Book{
		Title:   input.Title,
		BookID:  &bookID,
		Genre:   input.Genre,
		Authors: authorList,
	}
	r.BOOKSTORE = append(r.BOOKSTORE, &gBook)

	log.Println("Successfully added the book.")
	des := "Successfully added"
	return &model.PostStatus{
		Iserror:     false,
		Description: &des,
		BookID:      &bookID,
	}, err
}

// AddAuthor is the resolver for the add_author field.
func (r *mutationResolver) AddAuthor(ctx context.Context, input model.AuthorInput) (*model.PostStatus, error) {
	resp := model.PostStatus{
		Iserror:     false,
		Description: new(string),
		BookID:      new(string),
	}
	return &resp, nil
}

// UpdateBook is the resolver for the update_book field.
func (r *mutationResolver) UpdateBook(ctx context.Context, input *model.UpdateInput) (*model.PutStatus, error) {
	bookID := input.BookID
	log.Printf("updating the Book (BookID : %s) from BOOKSTORE\n", bookID)
	for ind, book := range r.BOOKSTORE {
		if book.BookID == &bookID {
			r.BOOKSTORE = append(r.BOOKSTORE[ind:], r.BOOKSTORE[:ind+1]...)
			uAuthorList := []*model.Author{}
			uTitle := ""
			var uGenre *model.BookGenre

			if input.Authors != nil {
				for _, a := range input.Authors {
					uAuthorList = append(uAuthorList, &model.Author{
						Name: *a.Name,
						Mail: a.Mail,
					})
				}
			} else {
				uAuthorList = book.Authors
			}
			if input.Genre != nil {
				uGenre = input.Genre
			}
			if input.Title != nil {
				uTitle = *input.Title
			}

			uBook := model.Book{
				Title:   uTitle,
				BookID:  &bookID,
				Genre:   uGenre,
				Authors: uAuthorList,
			}
			r.BOOKSTORE = append(r.BOOKSTORE, &uBook)

			des := "update success"
			log.Println(des)
			return &model.PutStatus{
				Iserror:     false,
				Description: &des,
			}, nil
		}
	}
	des := "failed to find book"
	return &model.PutStatus{
		Iserror:     true,
		Description: &des,
	}, nil
}

// DeleteBook is the resolver for the delete_book field.
func (r *mutationResolver) DeleteBook(ctx context.Context, bookID string) (*model.DeleteStatus, error) {
	log.Printf("Deleting the book (BookID : %s) from the BOOKSTORE\n", bookID)
	for ind, book := range r.BOOKSTORE {
		if book.BookID == &bookID {
			r.BOOKSTORE = append(r.BOOKSTORE[ind:], r.BOOKSTORE[:ind+1]...)
			des := "deletion success"
			log.Println(des)
			return &model.DeleteStatus{
				Iserror:     true,
				Description: &des,
			}, errors.New(des)
		}
	}
	des := "failed to find book"
	return &model.DeleteStatus{
		Iserror:     true,
		Description: &des,
	}, nil
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, bookID string) (*model.GetBookResult, error) {
	log.Println("Getting Book ", bookID, " from BOOKSTORE")
	for _, book := range r.BOOKSTORE {
		if *book.BookID == bookID {
			return &model.GetBookResult{Isexists: true, Book: book}, nil
		}
	}
	return &model.GetBookResult{Isexists: false}, errors.New("bookid not found")
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	log.Println("Getting all Books from BOOKSTORE")
	result := r.BOOKSERV.GetAllBook()

	return result, nil
}

// Getbooks is the resolver for the getbooks field.
func (r *queryResolver) Getbooks(ctx context.Context, getgenre *model.BookGenre) ([]*model.Book, error) {
	log.Println("Getting books which are ", getgenre, " from BOOKSTORE")
	var returnList = []*model.Book{}
	for _, book := range r.BOOKSTORE {
		if *book.Genre == *getgenre {
			returnList = append(returnList, book)
		}
	}
	return returnList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func nextBookID(bookstore []*model.Book) (string, error) {
	fmt.Println("Finding next Book ID")
	if len(bookstore) < 1 {
		fmt.Println("Next BOOKID: 1")
		return "1", nil
	}
	max := *bookstore[0].BookID
	for _, book := range bookstore {
		cBookID, err := strconv.Atoi(*book.BookID)
		if err != nil {
			return "0", err
		}
		maxVar, err := strconv.Atoi(max)
		if err != nil {
			return "0", err
		}
		if cBookID > maxVar {
			max = *book.BookID
		}
	}
	maxInt, _ := strconv.Atoi(max)
	max = strconv.Itoa(maxInt + 1)
	fmt.Println("Next BOOKID : ", max)
	return max, nil
}
