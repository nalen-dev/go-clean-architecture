package database

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dbURL = "postgresql://postgres:password@localhost:5432/be_test"

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}