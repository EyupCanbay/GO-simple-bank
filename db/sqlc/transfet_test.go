package db

import (
	"context"
	"fmt"
	"simple_bank/utils"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	to_account := createRandomAccount(t)
	from_account := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: pgtype.Int8{Int64: from_account.ID, Valid: true},
		ToAccountID:   pgtype.Int8{Int64: to_account.ID, Valid: true},
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	fmt.Println(transfer)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, to_account.ID, transfer.ToAccountID.Int64)
	require.Equal(t, from_account.ID, transfer.FromAccountID.Int64)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)

	fetch_transfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, fetch_transfer)
	require.Equal(t, transfer.ID, fetch_transfer.ID)
	require.Equal(t, transfer.FromAccountID, fetch_transfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, fetch_transfer.ToAccountID)
	require.Equal(t, transfer.Amount, fetch_transfer.Amount)
	require.WithinDuration(t, transfer.CreatedAt.Time, fetch_transfer.CreatedAt.Time, time.Second)

}

func TestUpdateTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: transfer.Amount,
	}

	updated_transfer, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updated_transfer)
	require.Equal(t, updated_transfer.ID, transfer.ID)
	require.Equal(t, updated_transfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, updated_transfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, updated_transfer.Amount, transfer.Amount)
	require.WithinDuration(t, updated_transfer.CreatedAt.Time, transfer.CreatedAt.Time, time.Second)

}

func TestDeleteTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	fetch_transfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, err.Error())
	require.Empty(t, fetch_transfer)

}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	list_transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, list_transfers, 5)

	for _, transfer := range list_transfers {
		require.NotEmpty(t, transfer)
	}

}
