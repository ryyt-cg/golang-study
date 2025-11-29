-- name: ListBooks :many
SELECT * FROM book ORDER BY title;

-- name: GetBook :one
SELECT * FROM book
WHERE id = ? LIMIT 1;

-- name: CreateBook :one
INSERT INTO book (
    title, published, author_id
) VALUES (?, ?, ?)
    RETURNING *;

-- name: UpdateBook :one
UPDATE book
set title = ?,
    published = ?,
    author_id = ?
WHERE id = ?
RETURNING *;