package db_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	db "github.com/raflynagachi/simplebank/db/sqlc"
	"github.com/raflynagachi/simplebank/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) db.Account {
	arg := db.CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(
		context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.Equal(t, arg.Currency, account.Currency)

	assert.NotZero(t, account.ID)
	assert.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	assert.Equal(t, acc1.Owner, acc2.Owner)
	assert.Equal(t, acc1.Balance, acc2.Balance)
	assert.Equal(t, acc1.Currency, acc2.Currency)
	assert.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomAccount(t)
	}

	arg := db.ListAccountsParams{
		Limit:  3,
		Offset: 2,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 3)

	for _, acc := range accounts {
		require.NotEmpty(t, acc)
	}
}

func TestUpdateAccount(t *testing.T) {
	acc := createRandomAccount(t)
	arg := db.UpdateAccountParams{
		ID:      acc.ID,
		Balance: util.RandomMoney(),
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)

	assert.Equal(t, acc.ID, acc2.ID)
	assert.Equal(t, acc.Owner, acc2.Owner)
	assert.Equal(t, arg.Balance+acc.Balance, acc2.Balance)
	assert.Equal(t, acc.Currency, acc2.Currency)
	assert.WithinDuration(t, acc.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}
