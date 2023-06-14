package db_test

import (
	"context"
	"math"
	"testing"

	db "github.com/raflynagachi/simplebank/db/sqlc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := db.NewStore(conn)

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := float64(10)

	errs := make(chan error)
	results := make(chan db.TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), db.TransferTxParams{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		assert.Equal(t, acc1.ID, transfer.FromAccountID)
		assert.Equal(t, acc2.ID, transfer.ToAccountID)
		assert.Equal(t, amount, transfer.Amount)
		assert.NotZero(t, transfer.ID)
		assert.NotZero(t, transfer.CreatedAt)

		// check transfer is exist
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		assert.Equal(t, acc1.ID, fromEntry.AccountID)
		assert.Equal(t, -amount, fromEntry.Amount)
		assert.NotZero(t, fromEntry.ID)
		assert.NotZero(t, fromEntry.CreatedAt)

		// check from entry is exist
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// check entries
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		assert.Equal(t, acc2.ID, toEntry.AccountID)
		assert.Equal(t, amount, toEntry.Amount)
		assert.NotZero(t, toEntry.ID)
		assert.NotZero(t, toEntry.CreatedAt)

		// check to entry is exist
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAcc := result.FromAccount
		require.NotEmpty(t, fromAcc)
		assert.Equal(t, acc1.ID, fromAcc.ID)

		toAcc := result.ToAccount
		require.NotEmpty(t, toAcc)
		assert.Equal(t, acc2.ID, toAcc.ID)

		// check account's balance
		diff1 := acc1.Balance - fromAcc.Balance
		diff2 := toAcc.Balance - acc2.Balance
		assert.Equal(t, diff1, diff2)
		assert.True(t, diff1 > 0)
		assert.True(t, math.Mod(diff1, amount) == 0)

		k := int(diff1 / amount)
		assert.True(t, k >= 1 && k <= n)
		assert.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated database
	updatedAcc1, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	updatedAcc2, err := testQueries.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)

	assert.Equal(t, acc1.Balance-float64(n)*amount, updatedAcc1.Balance)
	assert.Equal(t, acc2.Balance+float64(n)*amount, updatedAcc2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := db.NewStore(conn)

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 10
	amount := float64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := acc1.ID
		toAccountID := acc2.ID
		if i%2 == 0 {
			fromAccountID = acc2.ID
			toAccountID = acc1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), db.TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated database
	updatedAcc1, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	updatedAcc2, err := testQueries.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)

	assert.Equal(t, acc1.Balance, updatedAcc1.Balance)
	assert.Equal(t, acc2.Balance, updatedAcc2.Balance)
}
