package main

import (
	"context"
	"database/sql"
	"handling-migration/pkg/db"
	"handling-migration/pkg/db/sqlgen"
	"handling-migration/pkg/repository"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// connect to an in-memory database
	sqlDB := db.Sqlite{}
	dbConn, err := sqlDB.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	defer func(dbConn *sql.DB) {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbConn)

	queries := sqlgen.New(dbConn)
	ctx := context.Background()

	authorRepository := repository.NewAuthorRepository(ctx, queries)
	bookRepository := repository.NewBookRepository(ctx, queries)

	// Insert an author
	newAuthor := sqlgen.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio: sql.NullString{
			String: "Brian Wilson Kernighan is a Canadian computer scientist. He is a professor at the Department of Computer Science at Princeton University, USA.",
			Valid:  true,
		},
	}

	insertedAuthor, err := authorRepository.Insert(newAuthor)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("inserted author: %v\n", insertedAuthor)

	// List all authors
	authors, err := authorRepository.FindAll()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("authors: %v\n", authors)

	// Find author with books by id
	authorWithBooks, err := authorRepository.FindWithBooksById(1)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("author with books: %v\n", authorWithBooks)

	// List all books
	books, err := bookRepository.FindAll()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("authors: %v\n", books)

	// Author counts
	authorCount, err := authorRepository.Count()
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	log.Printf("author count: %d\n", authorCount)
}
