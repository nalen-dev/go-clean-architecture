package repository

import (
	"context"
	"go-clean-architecture/internal/model"
)

type UserRepository interface {
	CheckUserCredentials(ctx context.Context, user *model.UserCredentials) (error)
}


type JobRepository interface {

}