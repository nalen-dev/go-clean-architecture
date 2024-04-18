package user

import (
	"context"
	"go-clean-architecture/internal/model"
)

type Usecase interface{
	Login(ctx context.Context, request model.UserCredentials) (model.UserSession, error)
}
