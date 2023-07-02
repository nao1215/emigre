// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package schema

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO
    ` + "`" + `user` + "`" + ` (` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + `)
VALUES
    (?, ?, ?, ?)
`

type CreateUserParams struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
}

const deleteUser = `-- name: DeleteUser :execrows
DELETE FROM
    ` + "`" + `user` + "`" + `
WHERE
    ` + "`" + `id` + "`" + ` = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteUser, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getUser = `-- name: GetUser :one
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + `
FROM
    ` + "`" + `user` + "`" + `
WHERE
    ` + "`" + `id` + "`" + ` = ?
LIMIT
    1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `created_at` + "`" + `, ` + "`" + `updated_at` + "`" + `
FROM
    ` + "`" + `user` + "`" + `
ORDER BY
    ` + "`" + `created_at` + "`" + ` ASC
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserName = `-- name: UpdateUserName :execresult
UPDATE
    ` + "`" + `user` + "`" + `
SET
    ` + "`" + `name` + "`" + ` = ?,
    ` + "`" + `updated_at` + "`" + ` = ?
WHERE
    ` + "`" + `id` + "`" + ` = ?
`

type UpdateUserNameParams struct {
	Name      string
	UpdatedAt time.Time
	ID        string
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserName, arg.Name, arg.UpdatedAt, arg.ID)
}
