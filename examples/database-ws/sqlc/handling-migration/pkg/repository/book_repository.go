package repository

import (
	"context"
	"handling-migration/pkg/db/sqlgen"
	"log"
)

type BookRepositorier interface {
	FindAll() ([]sqlgen.Book, error)
	FindById(id int) (*sqlgen.Book, error)
	Insert(book sqlgen.Book) (*sqlgen.Book, error)
	Update(book sqlgen.Book) error
}

// BookRepository searches owner from the database
type BookRepository struct {
	ctx     context.Context
	queries *sqlgen.Queries
}

func NewBookRepository(ctx context.Context, queries *sqlgen.Queries) *BookRepository {
	return &BookRepository{
		ctx:     ctx,
		queries: queries,
	}
}

// FindAll
// find all books
func (repo *BookRepository) FindAll() ([]sqlgen.Book, error) {
	log.Printf("search all books\n")

	fetchedBooks, err := repo.queries.ListBooks(repo.ctx)
	if err != nil {
		log.Printf("fails to find all books, err: %v\n", err.Error())
		return nil, err
	}

	return fetchedBooks, nil
}

// FindById
// find the book by id
func (repo *BookRepository) FindById(id int) (*sqlgen.Book, error) {
	log.Printf("search book by id: %v\n", id)

	fetchedBook, err := repo.queries.GetBook(repo.ctx, int64(id))
	if err != nil {
		log.Printf("fails to find book by id: %v, err: %v", id, err.Error())
		return nil, err
	}

	return &fetchedBook, nil
}

// Insert a book
func (repo *BookRepository) Insert(book sqlgen.CreateBookParams) (*sqlgen.Book, error) {
	log.Printf("insert a book\n")
	newBook, err := repo.queries.CreateBook(repo.ctx, book)
	if err != nil {
		log.Printf("fails to insert a book, err: %v\n", err.Error())
		return nil, err
	}

	return &newBook, nil
}

// Update a book
func (repo *BookRepository) Update(book sqlgen.UpdateBookParams) (*sqlgen.Book, error) {
	log.Printf("update a book\n")
	updatedBook, err := repo.queries.UpdateBook(repo.ctx, book)
	if err != nil {
		log.Printf("fails to update a book, err: %v\n", err.Error())
		return nil, err
	}

	return &updatedBook, nil
}
