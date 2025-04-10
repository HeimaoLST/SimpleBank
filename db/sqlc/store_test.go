package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTX(t *testing.T) {
	
	account1 := CreateAccount(t)
	account2 := CreateAccount(t)

	amount := int64(10)

	arg := TransferTxParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        amount,
	}

	result, err := store.TransferTX(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// check transfer
	transfer := result.Transfer
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

}
