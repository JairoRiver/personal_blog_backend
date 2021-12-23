-- name: CreatePost :one
INSERT INTO posts (
                   user_id
                  ,title
                  ,subtitle
                  ,content
) VALUES (
             $1
         ,$2
         ,$3
         ,$4
         ) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT id
      ,Title
      ,Subtitle
      ,Content
      ,Created_at
      ,Updated_at
FROM posts
ORDER BY created_at
LIMIT $1
    OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET title = $2
  ,subtitle = $3
  ,content = $4
  ,updated_at = NOW()
WHERE id = $1
RETURNING title
         ,subtitle
         ,content
         ,updated_at;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;