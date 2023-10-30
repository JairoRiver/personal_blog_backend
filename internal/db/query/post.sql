-- name: CreatePost :one
INSERT INTO posts (
  category_id
 ,title
 ,subtitle
 ,content
 ,publicated
) VALUES (
  $1,$2,$3,$4,$5
) RETURNING *;

-- name: GetPostByIdPublic :one
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.content
      ,po.publicated
      ,po.category_id
      ,po.created_at
      ,po.updated_at
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name || "|" || ta.id) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id 
WHERE po.id = $1
  AND po.publicated IS TRUE
GROUP BY 1,2,3,4,5,6,7,8
LIMIT 1;

-- name: GetPostByIdPrivate :one
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.content
      ,po.publicated
      ,po.category_id
      ,po.created_at
      ,po.updated_at
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id 
WHERE po.id = $1
GROUP BY 1,2,3,4,5,6,7,8
LIMIT 1;

-- name: GetPostByCategoryPublic :many
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.content
      ,po.publicated
      ,po.category_id
      ,po.created_at
      ,po.updated_at
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id 
WHERE ca.id = $1
  AND po.publicated IS TRUE
GROUP BY 1,2,3,4,5,6,7,8;

-- name: GetPostByCategoryPrivate :many
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.content
      ,po.publicated
      ,po.category_id
      ,po.created_at
      ,po.updated_at
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id 
WHERE ca.id = $1
GROUP BY 1,2,3,4,5,6,7,8;

-- name: GetPostByTagPublic :many
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.content
      ,po.publicated
      ,po.category_id
      ,po.created_at
      ,po.updated_at
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id 
WHERE ta.id = $1
  AND po.publicated IS TRUE
GROUP BY 1,2,3,4,5,6,7,8;

-- name: GetPostByTagPrivate :many
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.content
      ,po.publicated
      ,po.category_id
      ,po.created_at
      ,po.updated_at
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id 
WHERE ta.id = $1
GROUP BY 1,2,3,4,5,6,7,8;

-- name: ListPostsPublic :many
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.created_at
      ,po.category_id
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id
WHERE po.publicated IS TRUE
GROUP BY 1,2,3,4,5,6;

-- name: ListPostsPrivate :many
SELECT po.id
      ,po.title
      ,po.subtitle
      ,po.created_at
      ,po.category_id
      ,ca.name AS category_name
      ,ARRAY_AGG(ta.name) AS tags
FROM posts AS po
JOIN categories AS ca ON po.category_id = ca.id
LEFT JOIN posts_tags AS pt ON pt.post_id = po.id
LEFT JOIN tags AS ta on pt.tag_id = ta.id
GROUP BY 1,2,3,4,5,6;

-- name: UpdatePost :one
UPDATE posts
SET
  title = COALESCE(sqlc.narg(title), title)
 ,subtitle = COALESCE(sqlc.narg(subtitle), subtitle)
 ,content = COALESCE(sqlc.narg(content), content)
 ,publicated = COALESCE(sqlc.narg(publicated), publicated)
 ,category_id = COALESCE(sqlc.narg(category_id), category_id)
 ,updated_at = NOW()
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;