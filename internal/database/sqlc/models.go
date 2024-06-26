// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Job struct {
	ID          int64            `json:"id"`
	UserID      pgtype.Int8      `json:"user_id"`
	Title       string           `json:"title"`
	Location    string           `json:"location"`
	Description pgtype.Text      `json:"description"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

type User struct {
	ID             int64            `json:"id"`
	FirstName      string           `json:"first_name"`
	LastName       pgtype.Text      `json:"last_name"`
	Email          string           `json:"email"`
	Password       string           `json:"password"`
	IsVerified     pgtype.Bool      `json:"is_verified"`
	ProfilePicture pgtype.Text      `json:"profile_picture"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
}
