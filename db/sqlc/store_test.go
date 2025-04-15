package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTX(t *testing.T) {

	store := NewStore(testDB)

	account1 := CreateAccount(t)
	account2 := CreateAccount(t)

	amount := int64(10)

	results := make(chan TransferTxResult)

	errs := make(chan error)
	arg := TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	}
	fmt.Println(">>before:", account1.Balance, account2.Balance)
	n := 5
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTX(context.Background(), arg)

			errs <- err
			results <- result
		}()
	}
	//check
	for i := 0; i < n; i++ {
		err := <-errs

		result := <-results

		require.NoError(t, err)
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer
		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
		require.Equal(t, arg.Amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)

		require.NoError(t, err)

		//check Entry

		fromEntry := result.FromEntry

		require.Equal(t, arg.FromAccountID, fromEntry.AccountID)
		require.Equal(t, -arg.Amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry

		require.Equal(t, arg.ToAccountID, toEntry.AccountID)
		require.Equal(t, arg.Amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//check account
		fmt.Println(">>tx:", result.FromAccount.Balance, result.ToAccount.Balance)

		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		//check balance

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)

		require.True(t, diff1 > 0)
		require.True(t, diff2%amount == 0)

		k := int(diff1 / amount)

		require.True(t, k >= 1 && k <= n)

		require.NotContains(t, existed, k)
		existed[k] = true

	}

	//check final balance

	updatedAccount1, err := store.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">>after:", updatedAccount1.Balance, updatedAccount2.Balance)

	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

}
