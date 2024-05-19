package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/takahiro-impara/simplebank/util"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.WithinDuration(t, account.CreatedAt, entry.CreatedAt, time.Second)

	return entry
}
func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, entry.Amount, entry2.Amount)
	require.Equal(t, entry.CreatedAt, entry2.CreatedAt)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	count := 10
	account := createRandomAccount(t)
	for i := 0; i < count; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntryParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntry(context.Background(), arg)
	require.NoError(t, err)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)

	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry.ID, entry2.ID)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.NotEqual(t, entry.Amount, entry2.Amount)
	require.Equal(t, entry.CreatedAt, entry2.CreatedAt)
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err2 := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}
