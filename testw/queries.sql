-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name)
VALUES ($1)
RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors
SET name = $2
WHERE id = $1
RETURNING *;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY title;

-- name: CreateBook :one
INSERT INTO books (title)
VALUES ($1)
RETURNING *;

-- name: UpdateBook :one
UPDATE books
SET title = $2
WHERE id = $1
RETURNING *;

-- name: SetBookAuthor :exec
INSERT INTO book_authors (book_id, author_id)
VALUES ($1, $2);

-- name: UnsetBookAuthors :exec
DELETE FROM book_authors
WHERE book_id = $1;

-- name: ListBooksByAuthorID :many
SELECT books.* FROM books, book_authors
WHERE books.id = book_authors.book_id AND book_authors.author_id = $1;

-- name: ListAuthorsByBookID :many
SELECT authors.* FROM authors, book_authors
WHERE authors.id = book_authors.author_id AND book_authors.book_id = $1;