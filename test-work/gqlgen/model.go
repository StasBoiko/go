// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

type AuthorInput struct {
	Name string `json:"name"`
}

type BookInput struct {
	Title     string  `json:"title"`
	AuthorIDs []int64 `json:"authorIDs"`
}