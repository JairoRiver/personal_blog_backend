-- name: CreateUser :one
INSERT INTO users (
    username,
    role_id,
    email,
    password
) VALUES (
             $1,
             $2,
             $3,
             $4
         ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id
      ,username
      ,email
      ,role_id
      ,created_at
      ,updated_at
FROM users
ORDER BY created_at
LIMIT $1
    OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET username = $2
  ,role_id = $3
  ,email = $4
  ,updated_at = NOW()
WHERE id = $1
RETURNING id
         ,username
         ,email
         ,role_id
         ,updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;