package main

import (
	"context"
	"go-clean-architecture/internal/database"
	"go-clean-architecture/internal/delivery/rest"
	uRep "go-clean-architecture/internal/repository/user"
	uUsecase "go-clean-architecture/internal/usecase/user"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

const dbURL = "postgresql://postgres:password@localhost:5432/be_test"

func main(){	

	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dbURL)


	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	store := database.NewStore(conn)

	r := gin.Default()

	userRepo := uRep.NewRepo(&store)
	userUseCase := uUsecase.GetUseCase(userRepo)

	handler := rest.NewHandler(userUseCase)
	rest.LoadRoutes(r, handler)

	r.Run(":8080")
}