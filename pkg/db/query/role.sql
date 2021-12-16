-- name: CreateRole :one
INSERT INTO roles (
    name
) VALUES (
             $1
         ) RETURNING *;

-- name: GetRole :one
SELECT * FROM roles
WHERE id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles
ORDER BY created_at;

-- name: UpdateRole :one
UPDATE roles
SET name = $2
  ,updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1;