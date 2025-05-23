// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: get_user_by_id.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getUserById = `-- name: GetUserById :one

SELECT id, name, created_at, updated_at 
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
