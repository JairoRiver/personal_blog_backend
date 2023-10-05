-- name: CreateCategory :one
INSERT INTO categories (
  name
) VALUES (
  $1
) RETURNING *;

-- name: GetCategory :one
SELECT ca.id
      ,ca.name
      ,ca.created_at
      ,ca.updated_at
FROM categories as ca
WHERE ca.id = $1 
LIMIT 1;

-- name: ListCategories :many
SELECT ca.id
      ,ca.name
      ,ca.created_at
      ,ca.updated_at
FROM categories as ca;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: UpdateCategory :one
UPDATE categories
SET
  name = COALESCE(sqlc.narg(name), name),
  updated_at = NOW()
WHERE
  id = sqlc.arg(id)
RETURNING *;