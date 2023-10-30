-- name: CreatePostTag :one
INSERT INTO posts_tags (
  post_id
 ,tag_id
) VALUES (
  $1,$2
) RETURNING *;

-- name: DeletePostTag :exec
DELETE FROM posts_tags
WHERE id = $1;