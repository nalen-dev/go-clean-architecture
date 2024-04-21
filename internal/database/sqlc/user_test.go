package database

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func TestFindUserByEmail(t *testing.T) { 

	user, err := testQueries.FindUserByEmail(context.Background(), "user@mail.com")

	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())	
	require.Empty(t, user)
}