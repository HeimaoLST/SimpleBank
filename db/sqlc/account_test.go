package db

import (
	"context"
	"github/heimaolst/simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {

	CreateAccount(t)

}

func TestGetAccount(t *testing.T) {
	account := CreateAccount(t)

	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)

	require.NotEmpty(t, account1)

	require.Equal(t, account.ID, account1.ID)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, account.Balance, account1.Balance)
	require.Equal(t, account.Currency, account1.Currency)
	require.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := CreateAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	account1, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, account1)

	require.Equal(t, account.ID, account1.ID)
	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, arg.Balance, account1.Balance)
	require.Equal(t, account.Currency, account1.Currency)
	require.WithinDuration(t, account.CreatedAt, account1.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := CreateAccount(t)

	_, err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account1)
}

func TestListAccounts(t *testing.T) {

	for i := 0; i < 10; i++ {
		CreateAccount(t)

	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
