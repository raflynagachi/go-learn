package db_test

import (
	"context"
	"testing"
	"time"

	db "github.com/raflynagachi/simplebank/db/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account db.Account) db.Entry {
	arg := db.CreateEntryParams{
		AccountID: account.ID,
		Amount:    0,
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	assert.Equal(t, arg.AccountID, entry.AccountID)
	assert.Equal(t, arg.Amount, entry.Amount)
	assert.NotZero(t, entry.ID)
	assert.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc)
}

func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entry := createRandomEntry(t, acc)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	assert.Equal(t, entry.ID, entry2.ID)
	assert.Equal(t, entry.AccountID, entry2.AccountID)
	assert.Equal(t, entry.Amount, entry2.Amount)
	assert.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntry(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 0; i < 5; i++ {
		createRandomEntry(t, acc)
	}
	arg := db.ListEntriesParams{
		AccountID: acc.ID,
		Limit:     3,
		Offset:    2,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	assert.Len(t, entries, 3)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
