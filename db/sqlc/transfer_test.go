package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/takahiro-impara/simplebank/util"
)

func createRandomTransfer(t *testing.T, account1 Account, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, account1, account2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.Equal(t, transfer.CreatedAt, transfer2.CreatedAt)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	count := 10
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < count; i++ {
		createRandomTransfer(t, account1, account2)
	}

	arg := ListTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)

	for _, entry := range transfers {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, account1, account2)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.NotEqual(t, transfer.Amount, arg.Amount)
	require.Equal(t, transfer.CreatedAt, transfer2.CreatedAt)
}

func TestDeleteTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	transfer := createRandomTransfer(t, account1, account2)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err2 := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}
