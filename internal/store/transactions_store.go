package store

import (
	"context"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type TransactionsStore interface {
	CreateTransaction(ctx context.Context, arg pgstore.CreateTransactionParams) (pgstore.CreateTransactionRow, error)
	ListTransactionsByUserAndDate(ctx context.Context, arg pgstore.ListTransactionsByUserAndDateParams) ([]pgstore.ListTransactionsByUserAndDateRow, error)
	ListTransactionsByUserIDPaginated(ctx context.Context, arg pgstore.ListTransactionsByUserIDPaginatedParams) ([]pgstore.ListTransactionsByUserIDPaginatedRow, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (pgstore.GetTransactionByIDRow, error)
	UpdateTransaction(ctx context.Context, arg pgstore.UpdateTransactionParams) (pgstore.Transaction, error)
	DeleteTransaction(ctx context.Context, id uuid.UUID) error
	PayTransaction(ctx context.Context, arg pgstore.PayTransactionParams) error
}
