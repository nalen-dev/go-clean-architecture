// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package database

import (
	"context"
)

const findUserByEmail = `-- name: FindUserByEmail :one
SELECT id, first_name, last_name, email, password, is_verified, profile_picture, created_at, updated_at FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) FindUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, findUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Password,
		&i.IsVerified,
		&i.ProfilePicture,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
