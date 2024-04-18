package user

import (
	"context"
	"errors"
	"go-clean-architecture/internal/database"
	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/repository"
)


type repo struct {
	DB *database.SQLStore
}

func NewRepo(db *database.SQLStore) repository.UserRepository {
	return &repo{ DB: db }
}

func (ur *repo) CheckUserCredentials(ctx context.Context, userCred *model.UserCredentials) (error){
	
	user, err := ur.DB.Queries.FindUserByEmail(ctx, userCred.Email)
	
	if err != nil {
		return err
	}

	if user.Password != userCred.Password {
		return errors.New("pasword incorect")
	}

	return nil
}