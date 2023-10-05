-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT u.id
      ,u.username
      ,u.email
      ,u.password 
      ,u.created_at
      ,u.updated_at
FROM users as u
WHERE u.id = $1 
LIMIT 1;

-- name: GetUserByUsername :one
SELECT u.id
      ,u.username
      ,u.password
FROM users as u
WHERE u.username = $1 
LIMIT 1;

-- name: ListUsers :many
SELECT u.id
      ,u.username
      ,u.email
      ,u.created_at
FROM users as u;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
  password = COALESCE(sqlc.narg(password), password),
  updated_at = NOW(),
  username = COALESCE(sqlc.narg(username), username),
  email = COALESCE(sqlc.narg(email), email)
WHERE
  id = sqlc.arg(id)
RETURNING *;