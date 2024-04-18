package user

import (
	"context"
	"go-clean-architecture/internal/model"
	"go-clean-architecture/internal/repository"
)


type userUsecase struct{
	userRepo repository.UserRepository
}

func GetUseCase(userRepo repository.UserRepository) Usecase{
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Login(ctx context.Context, request model.UserCredentials) (model.UserSession, error){

	if err := u.userRepo.CheckUserCredentials(ctx, &request); err != nil {
		return  model.UserSession{}, err
	}

	//todo: add new session

	return model.UserSession{}, nil

}