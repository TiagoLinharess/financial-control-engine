package store

import (
	"context"
	"financialcontrol/internal/store/pgstore"

	"github.com/google/uuid"
)

type MonthlyTransactionsStore interface {
	CreateMonthlyTransaction(ctx context.Context, arg pgstore.CreateMonthlyTransactionParams) (pgstore.MonthlyTransaction, error)
	ListMonthlyTransactionsByUserIDPaginated(ctx context.Context, arg pgstore.ListMonthlyTransactionsByUserIDPaginatedParams) ([]pgstore.ListMonthlyTransactionsByUserIDPaginatedRow, error)
	GetMonthlyTransactionByID(ctx context.Context, id uuid.UUID) (pgstore.GetMonthlyTransactionByIDRow, error)
	UpdateMonthlyTransaction(ctx context.Context, arg pgstore.UpdateMonthlyTransactionParams) (pgstore.MonthlyTransaction, error)
	DeleteMonthlyTransaction(ctx context.Context, id uuid.UUID) error
}
