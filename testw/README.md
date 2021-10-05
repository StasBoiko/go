# Test work

This test job contains a simple Golang service which serves the graphql server.

## Prerequisite

1. Docker installed.

2. docker-compose installed.

3. Please make sure you have VPN disabled as this may affect the docker image building process.

## Local Deployment

In order to deploy the application locally, do the following:

1. Build the image

```sh
docker-compose build
```

2. Start the stack

```sh
docker-compose up -d
```

3. Apply the database schema (should be run only once)

```sh
docker exec -it test-work-app bash /app/apply_db_schema.sh
```

4. Check the application at http://localhost:8080/.

## Application Usage

To add an author, you can use (mutation create author):

```json
mutation { createAuthor ( data: {name: "Vasya"} ) { id name } }
```

To add a book, you can use (mutation to create a book):

```json
mutation { createBook ( data: {title: "Golang for dummies", authorIDs: "1"} ) { id title } }
```

To get all authors you can use (request to get all authors):

```json
query { authors { id name books { id title } } }
```

For getting all books you can use (request to get all books):

```json
query { books { id title authors { id name } } }
```

You can update the author (mutation for updating the author by ID):

```json
mutation { updateAuthor (id: 1, data: {name: "Kolya"} ) { id name } }
```

You can update the book (mutation for updating the book):

```json
mutation { updateBook (id: 1, data: {title: "Golang for advanced dummies", authorIDs: 1} ) { id title } }
```