package gqlgen

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/StasBoiko/test-work/postgres"
)

// Resolver connects individual resolvers with the datalayer.
type Resolver struct {
	Repository postgres.Repository
}

func (r *authorResolver) Books(ctx context.Context, obj *postgres.Author) ([]postgres.Book, error) {
	return r.Repository.ListBooksByAuthorID(ctx, obj.ID)
}

func (r *bookResolver) Authors(ctx context.Context, obj *postgres.Book) ([]postgres.Author, error) {
	return r.Repository.ListAuthorsByBookID(ctx, obj.ID)
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, data AuthorInput) (*postgres.Author, error) {
	author, err := r.Repository.CreateAuthor(ctx, data.Name)
	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, id int64, data AuthorInput) (*postgres.Author, error) {
	author, err := r.Repository.UpdateAuthor(ctx, postgres.UpdateAuthorParams{
		ID:   id,
		Name: data.Name,
	})
	if err != nil {
		println(err)
		return nil, err
	}
	return &author, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, data BookInput) (*postgres.Book, error) {
	book, err := r.Repository.CreateBook(ctx, data.Title, data.AuthorIDs)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, id int64, data BookInput) (*postgres.Book, error) {
	book, err := r.Repository.UpdateBook(ctx, postgres.UpdateBookParams{
		ID:    id,
		Title: data.Title,
	}, data.AuthorIDs)

	if err != nil {
		println(err)
		return nil, err
	}
	return book, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]postgres.Author, error) {
	return r.Repository.ListAuthors(ctx)
}

func (r *queryResolver) Books(ctx context.Context) ([]postgres.Book, error) {
	return r.Repository.ListBooks(ctx)
}

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Book returns BookResolver implementation.
func (r *Resolver) Book() BookResolver { return &bookResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
