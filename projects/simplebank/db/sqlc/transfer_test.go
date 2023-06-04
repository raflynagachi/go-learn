package db_test

import (
	"context"
	"testing"
	"time"

	db "github.com/raflynagachi/simplebank/db/sqlc"
	"github.com/raflynagachi/simplebank/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, acc1, acc2 db.Account) db.Transfer {
	arg := db.CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	assert.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	assert.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	assert.Equal(t, arg.Amount, transfer.Amount)
	assert.NotZero(t, transfer.ID)
	assert.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	createRandomTransfer(t, acc1, acc2)
}

func TestGetTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	trf := createRandomTransfer(t, acc1, acc2)

	trf2, err := testQueries.GetTransfer(context.Background(), trf.ID)
	require.NoError(t, err)
	require.NotEmpty(t, trf2)
	assert.Equal(t, trf.ID, trf.ID)
	assert.Equal(t, trf.FromAccountID, trf.FromAccountID)
	assert.Equal(t, trf.ToAccountID, trf.ToAccountID)
	assert.Equal(t, trf.Amount, trf.Amount)
	assert.WithinDuration(t, trf.CreatedAt, trf.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, acc1, acc2)
	}
	arg := db.ListTransfersParams{
		ToAccountID:   acc2.ID,
		FromAccountID: acc1.ID,
		Limit:         3,
		Offset:        2,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	assert.Len(t, transfers, 3)

	for _, trf := range transfers {
		require.NotEmpty(t, trf)
		assert.Equal(t, arg.FromAccountID, trf.FromAccountID)
		assert.Equal(t, arg.ToAccountID, trf.ToAccountID)
	}
}
