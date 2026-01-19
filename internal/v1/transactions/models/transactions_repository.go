package models

import (
	"context"
	"financialcontrol/internal/models"
	"financialcontrol/internal/models/errors"

	"github.com/google/uuid"
)

type TransactionsRepository interface {
	Create(context context.Context, transaction CreateTransaction) (ShortTransaction, []errors.ApiError)
	Read(context context.Context, params models.PaginatedParams) ([]Transaction, int64, []errors.ApiError)
	ReadInToDates(context context.Context, params models.PaginatedParamsWithDateRange) ([]Transaction, int64, []errors.ApiError)
	ReadById(context context.Context, id uuid.UUID) (Transaction, []errors.ApiError)
	Update(context context.Context, transaction Transaction) (ShortTransaction, []errors.ApiError)
	Delete(context context.Context, id uuid.UUID) []errors.ApiError
	Pay(context context.Context, id uuid.UUID, paid bool) []errors.ApiError
	GetCreditcardTotalAmount(ctx context.Context, model TransactionsCreditCardTotal) (float64, error)
}
