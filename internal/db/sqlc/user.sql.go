// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: user.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  email,
  password
) VALUES (
  $1, $2, $3
) RETURNING id, username, email, password, created_at, updated_at
`

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT u.id
      ,u.username
      ,u.email
      ,u.password 
      ,u.created_at
      ,u.updated_at
FROM users as u
WHERE u.id = $1 
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT u.id
      ,u.username
      ,u.password
FROM users as u
WHERE u.username = $1 
LIMIT 1
`

type GetUserByUsernameRow struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT u.id
      ,u.username
      ,u.email
      ,u.created_at
FROM users as u
`

type ListUsersRow struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListUsersRow{}
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  password = COALESCE($1, password),
  updated_at = NOW(),
  username = COALESCE($2, username),
  email = COALESCE($3, email)
WHERE
  id = $4
RETURNING id, username, email, password, created_at, updated_at
`

type UpdateUserParams struct {
	Password pgtype.Text `json:"password"`
	Username pgtype.Text `json:"username"`
	Email    pgtype.Text `json:"email"`
	ID       uuid.UUID   `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Password,
		arg.Username,
		arg.Email,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
