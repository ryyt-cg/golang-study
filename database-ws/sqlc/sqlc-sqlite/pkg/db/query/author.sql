-- name: GetAuthor :one
SELECT * FROM author
WHERE id = ? LIMIT 1;

-- name: GetAuthorWithBooks :many
SELECT sqlc.embed(author), sqlc.embed(book) FROM author
LEFT JOIN book ON author.id = book.author_id
WHERE author.id = ?;

-- name: ListAuthors :many
SELECT * FROM author
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO author (
    name, bio
) VALUES (?, ?)
    RETURNING *;

-- name: UpdateAuthor :one
UPDATE author
set name = ?,
    bio = ?
WHERE id = ?
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM author
WHERE id = ?;
