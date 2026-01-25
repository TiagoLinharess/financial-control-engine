package repositories

import (
	c "context"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	"financialcontrol/internal/repositories/dtos"
	pgs "financialcontrol/internal/store/pgstore"
	u "financialcontrol/internal/utils"
	tm "financialcontrol/internal/v1/transactions/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r Repository) CreateTransaction(context c.Context, transaction tm.CreateTransaction) (tm.ShortTransaction, []e.ApiError) {
	value := u.Float64ToNumeric(transaction.Value)

	param := pgs.CreateTransactionParams{
		UserID:                    transaction.UserID,
		Name:                      transaction.Name,
		Date:                      pgtype.Timestamptz{Time: transaction.Date, Valid: true},
		Value:                     value,
		Paid:                      transaction.Paid,
		CategoryID:                transaction.CategoryID,
		CreditCardID:              u.UUIDToPgTypeUUID(transaction.CreditcardID),
		MonthlyTransactionsID:     u.UUIDToPgTypeUUID(transaction.MonthlyTransactionsID),
		AnnualTransactionsID:      u.UUIDToPgTypeUUID(transaction.AnnualTransactionsID),
		InstallmentTransactionsID: u.UUIDToPgTypeUUID(transaction.InstallmentTransactionsID),
	}

	createdTransaction, err := r.store.CreateTransaction(context, param)

	if err != nil {
		return tm.ShortTransaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return tm.ShortTransaction{
		ID:        createdTransaction.ID,
		Name:      createdTransaction.Name,
		Date:      createdTransaction.Date.Time,
		Value:     u.NumericToFloat64(createdTransaction.Value),
		Paid:      createdTransaction.Paid,
		CreatedAt: createdTransaction.CreatedAt.Time,
		UpdatedAt: createdTransaction.UpdatedAt.Time,
	}, nil
}

func (r Repository) ReadTransactions(context c.Context, params m.PaginatedParams) ([]tm.Transaction, int64, []e.ApiError) {
	args := pgs.ListTransactionsByUserIDPaginatedParams{
		UserID: params.UserID,
		Limit:  params.Limit,
		Offset: params.Offset,
	}

	transactions, err := r.store.ListTransactionsByUserIDPaginated(context, args)

	if err != nil {
		return []tm.Transaction{}, 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(transactions) == 0 {
		return []tm.Transaction{}, 0, nil
	}

	var transactionModels []tm.Transaction
	count := transactions[len(transactions)-1].TotalCount

	for _, transaction := range transactions {
		transactionModels = append(transactionModels, dtos.StoreTransactionToTransaction(dtos.StoreTransactionPaginatedToStoreTransaction(transaction)))
	}

	return transactionModels, count, nil
}

func (r Repository) ReadTransactionsInToDates(context c.Context, params m.PaginatedParamsWithDateRange) ([]tm.Transaction, int64, []e.ApiError) {
	args := pgs.ListTransactionsByUserAndDateParams{
		UserID: params.UserID,
		Limit:  params.Limit,
		Offset: params.Offset,
		Date:   pgtype.Timestamptz{Time: params.StartDate, Valid: true},
		Date_2: pgtype.Timestamptz{Time: params.EndDate, Valid: true},
	}

	transactions, err := r.store.ListTransactionsByUserAndDate(context, args)

	if err != nil {
		return []tm.Transaction{}, 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(transactions) == 0 {
		return []tm.Transaction{}, 0, nil
	}

	var transactionModels []tm.Transaction
	count := transactions[len(transactions)-1].TotalCount

	for _, transaction := range transactions {
		transactionModels = append(transactionModels, dtos.StoreTransactionToTransaction(dtos.StoreTransactionListToStoreTransaction(transaction)))
	}

	return transactionModels, count, nil
}

func (r Repository) ReadTransactionById(context c.Context, id uuid.UUID) (tm.Transaction, []e.ApiError) {
	transaction, err := r.store.GetTransactionByID(context, id)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return dtos.StoreTransactionToTransaction(transaction), nil
}

func (r Repository) UpdateTransaction(context c.Context, transaction tm.Transaction) (tm.ShortTransaction, []e.ApiError) {
	value := u.Float64ToNumeric(transaction.Value)

	var creditCardID pgtype.UUID
	if transaction.Creditcard != nil {
		creditCardID = u.UUIDToPgTypeUUID(&transaction.Creditcard.ID)
	}

	params := pgs.UpdateTransactionParams{
		ID:                        transaction.ID,
		Name:                      transaction.Name,
		Date:                      pgtype.Timestamptz{Time: transaction.Date, Valid: true},
		Value:                     value,
		Paid:                      transaction.Paid,
		CategoryID:                transaction.Category.ID,
		CreditCardID:              creditCardID,
		MonthlyTransactionsID:     u.UUIDToPgTypeUUID(transaction.MonthlyTransaction),
		AnnualTransactionsID:      u.UUIDToPgTypeUUID(transaction.AnnualTransaction),
		InstallmentTransactionsID: u.UUIDToPgTypeUUID(transaction.InstallmentTransaction),
	}

	transactionUpdated, err := r.store.UpdateTransaction(context, params)

	if err != nil {
		return tm.ShortTransaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return tm.ShortTransaction{
		ID:        transactionUpdated.ID,
		Name:      transactionUpdated.Name,
		Date:      transactionUpdated.Date.Time,
		Value:     u.NumericToFloat64(transactionUpdated.Value),
		Paid:      transactionUpdated.Paid,
		CreatedAt: transactionUpdated.CreatedAt.Time,
		UpdatedAt: transactionUpdated.UpdatedAt.Time,
	}, nil
}

func (r Repository) DeleteTransaction(context c.Context, id uuid.UUID) []e.ApiError {
	err := r.store.DeleteTransaction(context, id)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}

func (r Repository) PayTransaction(context c.Context, id uuid.UUID, paid bool) []e.ApiError {
	params := pgs.PayTransactionParams{
		ID:   id,
		Paid: paid,
	}

	err := r.store.PayTransaction(context, params)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}

func (r Repository) GetCreditcardTotalAmount(ctx c.Context, model tm.TransactionsCreditCardTotal) (float64, error) {
	normalizedDate := time.Date(model.Date.Year(), model.Date.Month(), model.Date.Day(), 0, 0, 0, 0, model.Date.Location())
	dateOnly := pgtype.Date{Time: normalizedDate, Valid: true}

	params := pgs.GetTotalFromCreditTransactionsByUserAndMonthParams{
		Column1:      dateOnly,
		CreditCardID: u.UUIDToPgTypeUUID(&model.CreditcardID),
		UserID:       model.UserID,
	}

	total, err := r.store.GetTotalFromCreditTransactionsByUserAndMonth(ctx, params)

	if err != nil {
		return 0, err
	}

	if !total.Valid {
		return 0, nil
	}

	floatValue := u.NumericToFloat64(total)

	return floatValue, nil
}
