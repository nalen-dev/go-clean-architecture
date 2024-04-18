package database

import (
	db "go-clean-architecture/internal/database/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLStore struct{
	ConnPool 	*pgxpool.Pool
	Queries		*db.Queries		
}

func NewStore(connPool *pgxpool.Pool) SQLStore{
	return SQLStore { 
		ConnPool: connPool,
		Queries: db.New(connPool),
	}
}