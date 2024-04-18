package rest

import "go-clean-architecture/internal/usecase/user"

type handler struct{
	userUsecase user.Usecase
}

func NewHandler(userUsecase user.Usecase) *handler{
	return &handler{
		userUsecase: userUsecase,
	}
}

