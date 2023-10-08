-- name: CreateTag :one
INSERT INTO tags (
  name,
  image_url
) VALUES (
  $1,$2
) RETURNING *;

-- name: GetTag :one
SELECT id
      ,name
      ,image_url
      ,created_at
      ,updated_at
FROM tags
WHERE id = $1 
LIMIT 1;

-- name: GetTagByName :one
SELECT id
      ,name
      ,image_url
      ,created_at
      ,updated_at
FROM tags
WHERE name = $1 
LIMIT 1;

-- name: ListTags :many
SELECT id
      ,name
      ,image_url
      ,created_at
      ,updated_at
FROM tags;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1;