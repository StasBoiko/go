Test work

This test job contains a simple Golang service which serves the graphql server

Please make sure you have VPN disabled as this may affect the execution of the commands described below (docker-compose execution)

The easiest way to launch this application is to follow these steps:

1) docker-compose build
2) docker-compose up -d
3) docker exec -it test-work_app_1 bash /app/apply_db_schema.sh

Go to:
http://localhost:8080/

To add an author, you can use (mutation create author):

mutation {
 createAuthor (
  data: {name: "Vasya"}
 ) {
  id
  name
 }
}

To add a book, you can use (mutation to create a book):

mutation {
 createBook (
  data: {title: "Golang for dummies", authorIDs: "1"}
 ) {
  id
  title
 }
}

To get all authors you can use (request to get all authors):

query {
 authors {
   id
   name
  books {
   id
   title
  }
 }
}

For getting all books you can use (request to get all books):

query {
 books {
   id
   title
  authors {
   id
   name
  }
 }
}

You can update the author (mutation for updating the author by ID):

mutation {
 updateAuthor (id: 1,
 data: {name: "Kolya"}
 ) {
  id
  name
 }
}

You can update the book (mutation for updating the book):

mutation {
 updateBook (id: 1,
 data: {title: "Golang for advanced dummies", authorIDs: 1}
 ) {
  id
  title
 }
}