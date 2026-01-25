package models

import (
	"context"
	"financialcontrol/internal/categories"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	cr "financialcontrol/internal/v1/creditcards/models"
	tm "financialcontrol/internal/v1/transactions/models"

	"github.com/google/uuid"
)

type TransactionsRepository interface {
	ReadCategoryByID(context context.Context, categoryID uuid.UUID) (categories.Category, []e.ApiError)
	ReadCreditCardByID(context context.Context, creditCardId uuid.UUID) (cr.CreditCard, []e.ApiError)
	CreateTransaction(context context.Context, transaction tm.CreateTransaction) (tm.ShortTransaction, []e.ApiError)
	ReadTransactions(context context.Context, params m.PaginatedParams) ([]tm.Transaction, int64, []e.ApiError)
	ReadTransactionsInToDates(context context.Context, params m.PaginatedParamsWithDateRange) ([]tm.Transaction, int64, []e.ApiError)
	ReadTransactionById(context context.Context, id uuid.UUID) (tm.Transaction, []e.ApiError)
	UpdateTransaction(context context.Context, transaction tm.Transaction) (tm.ShortTransaction, []e.ApiError)
	DeleteTransaction(context context.Context, id uuid.UUID) []e.ApiError
	PayTransaction(context context.Context, id uuid.UUID, paid bool) []e.ApiError
	GetCreditcardTotalAmount(ctx context.Context, model tm.TransactionsCreditCardTotal) (float64, error)
}
