-- name: CreateUser :one
INSERT INTO users (id, username, password_hash, email)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByID :one
SELECT id, username, password_hash, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT id, username, password_hash, email, created_at, updated_at
FROM users
WHERE username = $1;

-- name: GetUserByEmail :one
SELECT id, username, password_hash, email, created_at, updated_at
FROM users
WHERE email = $1;

-- name: UpdateUser :exec
UPDATE users
SET username = $2, password_hash = $3, email = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
