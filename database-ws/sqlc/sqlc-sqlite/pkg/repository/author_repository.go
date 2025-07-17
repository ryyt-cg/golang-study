package repository

import (
	"context"
	"log"
	"sqlc-sqlite/pkg/db/sqlgen"
)

type AuthorRepositorier interface {
	FindAll() ([]sqlgen.Author, error)
	FindById(id int) (*sqlgen.Author, error)
	FindWithBooksById(id int) ([]sqlgen.Author, error)
	Insert(author sqlgen.Author) (*sqlgen.Author, error)
	Update(author sqlgen.Author) (*sqlgen.Author, error)
}

// AuthorRepository searches owner from the database
type AuthorRepository struct {
	ctx     context.Context
	queries *sqlgen.Queries
}

func NewAuthorRepository(ctx context.Context, queries *sqlgen.Queries) *AuthorRepository {
	return &AuthorRepository{
		ctx:     ctx,
		queries: queries,
	}
}

// FindAll
// find all authors
func (repo *AuthorRepository) FindAll() ([]sqlgen.Author, error) {
	log.Printf("search all authors\n")

	fetchedAuthors, err := repo.queries.ListAuthors(repo.ctx)
	if err != nil {
		log.Printf("fails to find all authors, err: %v\n", err.Error())
		return nil, err
	}

	return fetchedAuthors, nil
}

// FindById
// find the author by id
func (repo *AuthorRepository) FindById(id int64) (*sqlgen.Author, error) {
	log.Printf("search owner by id: %v\n", id)

	fetchedAuthor, err := repo.queries.GetAuthor(repo.ctx, id)
	if err != nil {
		log.Printf("fails to find owner by id: %v, err: %v", id, err.Error())
		return nil, err
	}

	return &fetchedAuthor, nil
}

// FindWithBooksById
// find the author by id with books
func (repo *AuthorRepository) FindWithBooksById(id int64) ([]sqlgen.GetAuthorWithBooksRow, error) {
	log.Printf("search author by id: %v\n", id)

	fetchedAuthors, err := repo.queries.GetAuthorWithBooks(repo.ctx, id)
	if err != nil {
		log.Printf("fails to find owner by id: %v, err: %v", id, err.Error())
		return nil, err
	}

	return fetchedAuthors, nil
}

// Insert
// insert an author
func (repo *AuthorRepository) Insert(author sqlgen.CreateAuthorParams) (*sqlgen.Author, error) {
	log.Printf("insert an author\n")
	newAuthor, err := repo.queries.CreateAuthor(repo.ctx, author)
	if err != nil {
		log.Printf("fails to insert an author, err: %v\n", err.Error())
		return nil, err
	}

	return &newAuthor, nil
}

// Update an author
func (repo *AuthorRepository) Update(author sqlgen.UpdateAuthorParams) (*sqlgen.Author, error) {
	log.Printf("update an author\n")
	updatedAuthor, err := repo.queries.UpdateAuthor(repo.ctx, author)
	if err != nil {
		log.Printf("fails to update an author, err: %v\n", err.Error())
		return nil, err
	}

	return &updatedAuthor, nil
}

// Delete an author
func (repo *AuthorRepository) Delete(id int64) error {
	log.Printf("delete an author\n")
	err := repo.queries.DeleteAuthor(repo.ctx, id)
	if err != nil {
		log.Printf("fails to delete an author, err: %v\n", err.Error())
		return err
	}

	return nil
}
