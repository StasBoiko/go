package gqlgen

import (
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/StasBoiko/test-work/postgres"
	"github.com/StasBoiko/test-work/postgres/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mutation Resolver

func TestMutationResolverCreateAuthor(t *testing.T) {
	ctxVe := context.TODO()
	var author postgres.Author
	tests := []struct {
		name         string
		wantAuthor   postgres.Author
		wantErr      bool
		c            context.Context
		mockBehavior func(r *mock.MockRepository, author interface{})
	}{
		{
			name:       "ok",
			wantAuthor: author, wantErr: false, c: ctxVe,
			mockBehavior: func(r *mock.MockRepository, author interface{}) {
				r.EXPECT().CreateAuthor(gomock.Any(), "Vasya").Return(author, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock.NewMockRepository(c)
			tt.mockBehavior(repo, author)
			r := &mutationResolver{&Resolver{Repository: repo}}
			gotAuthors, err := r.CreateAuthor(ctxVe, AuthorInput{Name: "Vasya"})

			if tt.wantErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.EqualValues(t, gotAuthors, &tt.wantAuthor)
		})
	}
}

func TestMutationResolverCreateBook(t *testing.T) {
	ctxVe := context.TODO()
	var book postgres.Book
	autorIDs := []int64{1, 2}
	bookTitle := "Test book title"
	tests := []struct {
		name         string
		wantBook     postgres.Book
		wantErr      bool
		c            context.Context
		mockBehavior func(r *mock.MockRepository, title, authorIDs interface{})
	}{
		{
			name:     "ok",
			wantBook: book, wantErr: false, c: ctxVe,
			mockBehavior: func(r *mock.MockRepository, title, authorIDs interface{}) {
				r.EXPECT().CreateBook(gomock.Any(), bookTitle, autorIDs).Return(&book, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock.NewMockRepository(c)
			tt.mockBehavior(repo, bookTitle, autorIDs)
			r := &mutationResolver{&Resolver{Repository: repo}}
			gotAuthors, err := r.CreateBook(ctxVe, BookInput{Title: bookTitle, AuthorIDs: autorIDs})

			if tt.wantErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.EqualValues(t, gotAuthors, &tt.wantBook)
		})
	}
}

//FIXME
//The following two test functions (TestMutationResolverUpdateAuthor and TestMutationResolverUpdateBook) throw an error:
//there are no expected calls of the method "UpdateAuthor"/"UpdateBook" for that receiver
//This error is displayed in cases when the code does not see the generated methods (methods inside mocks)
//But postgres/mock/postgres_mock.go file contains the "UpdateAuthor"/"UpdateBook" methods
//For some reason the code doesn't see it. Need to investigate
func TestMutationResolverUpdateAuthor(t *testing.T) {
	ctxVe := context.TODO()
	c := gomock.NewController(t)
	defer c.Finish()
	repo := mock.NewMockRepository(c)
	r := &mutationResolver{&Resolver{Repository: repo}}
	_, _ = r.Mutation().UpdateAuthor(context.Background(), 1, AuthorInput{
		Name: "test name",
	})
	updateAuthor := postgres.UpdateAuthorParams{ID: 1, Name: "test name"}
	result, err := repo.UpdateAuthor(ctxVe, updateAuthor)
	ok(t, err)
	equals(t, result, AuthorInput{
		Name: "test name",
	})
}

func TestMutationResolverUpdateBook(t *testing.T) {
	ctxVe := context.TODO()
	autorIDs := []int64{1}
	c := gomock.NewController(t)
	defer c.Finish()
	repo := mock.NewMockRepository(c)
	r := &mutationResolver{&Resolver{Repository: repo}}
	_, _ = r.Mutation().UpdateBook(context.Background(), 1, BookInput{
		Title: "test title",
	})
	updateBook := postgres.UpdateBookParams{ID: 1, Title: "test title"}
	result, err := repo.UpdateBook(ctxVe, updateBook, autorIDs)
	ok(t, err)
	equals(t, result, BookInput{
		Title: "test title",
	})
}

// Query Resolver

func TestQueryResolverAuthors(t *testing.T) {
	ctxVe := context.TODO()
	var authors []postgres.Author
	tests := []struct {
		name         string
		wantAuthor   []postgres.Author
		wantErr      bool
		c            context.Context
		mockBehavior func(r *mock.MockRepository, author interface{})
	}{
		{
			name:       "ok",
			wantAuthor: authors, wantErr: false, c: ctxVe,
			mockBehavior: func(r *mock.MockRepository, author interface{}) {
				r.EXPECT().ListAuthors(ctxVe).Return(authors, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock.NewMockRepository(c)
			tt.mockBehavior(repo, authors)
			r := &queryResolver{&Resolver{Repository: repo}}
			gotAuthors, err := r.Query().Authors(ctxVe)

			if tt.wantErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.EqualValues(t, gotAuthors, tt.wantAuthor)
		})
	}
}

func TestQueryResolverBooks(t *testing.T) {
	ctxVe := context.TODO()
	var books []postgres.Book
	tests := []struct {
		name         string
		wantBook     []postgres.Book
		wantErr      bool
		c            context.Context
		mockBehavior func(r *mock.MockRepository, book interface{})
	}{
		{
			name:     "ok",
			wantBook: books, wantErr: false, c: ctxVe,
			mockBehavior: func(r *mock.MockRepository, book interface{}) {
				r.EXPECT().ListBooks(ctxVe).Return(books, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := mock.NewMockRepository(c)
			tt.mockBehavior(repo, books)
			r := &queryResolver{&Resolver{Repository: repo}}
			gotBooks, err := r.Query().Books(ctxVe)

			if tt.wantErr {
				require.NotNil(t, err)
				return
			}
			require.Nil(t, err)
			assert.EqualValues(t, gotBooks, tt.wantBook)
		})
	}
}

//// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

//// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
