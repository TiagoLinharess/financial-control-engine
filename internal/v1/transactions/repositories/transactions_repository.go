package repositories

import (
	c "context"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	st "financialcontrol/internal/store"
	pgs "financialcontrol/internal/store/pgstore"
	u "financialcontrol/internal/utils"
	tm "financialcontrol/internal/v1/transactions/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TransactionsRepository struct {
	store st.TransactionsStore
}

func NewTransactionsRepository(store st.TransactionsStore) tm.TransactionsRepository {
	return TransactionsRepository{store: store}
}

func (t TransactionsRepository) Create(context c.Context, transaction tm.CreateTransaction) (tm.Transaction, []e.ApiError) {
	value, err := u.Float64ToNumeric(transaction.Value)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	param := pgs.CreateTransactionParams{
		UserID:                    transaction.UserID,
		Name:                      transaction.Name,
		Date:                      pgtype.Timestamptz{Time: transaction.Date, Valid: true},
		Value:                     value,
		CategoryID:                transaction.CategoryID,
		CreditCardID:              u.UUIDToPgTypeUUID(transaction.CreditCardID),
		MonthlyTransactionsID:     u.UUIDToPgTypeUUID(transaction.MonthlyTransactionsID),
		AnnualTransactionsID:      u.UUIDToPgTypeUUID(transaction.AnnualTransactionsID),
		InstallmentTransactionsID: u.UUIDToPgTypeUUID(transaction.InstallmentTransactionsID),
	}

	createdTransaction, err := t.store.CreateTransaction(context, param)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	transactionModel, err := storeToModel(createdTransaction)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return transactionModel, nil
}

func (t TransactionsRepository) Read(context c.Context, params m.PaginatedParams) ([]tm.Transaction, int64, []e.ApiError) {
	args := pgs.ListTransactionsByUserIDPaginatedParams{
		UserID: params.UserID,
		Limit:  params.Limit,
		Offset: params.Offset,
	}

	transactions, err := t.store.ListTransactionsByUserIDPaginated(context, args)

	if err != nil {
		return []tm.Transaction{}, 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(transactions) == 0 {
		return []tm.Transaction{}, 0, nil
	}

	var transactionModels []tm.Transaction
	count := transactions[len(transactions)-1].TotalCount

	for _, transaction := range transactions {
		transactionModel, err := storeToModel(paginatedModelToStore(transaction))

		if err != nil {
			return []tm.Transaction{}, 0, []e.ApiError{e.StoreError{Message: err.Error()}}
		}

		transactionModels = append(transactionModels, transactionModel)
	}

	return transactionModels, count, nil
}

func (t TransactionsRepository) ReadInToDates(context c.Context, params m.PaginatedParamsWithDateRange) ([]tm.Transaction, int64, []e.ApiError) {
	args := pgs.ListTransactionsByUserAndDateParams{
		UserID: params.UserID,
		Limit:  params.Limit,
		Offset: params.Offset,
		Date:   pgtype.Timestamptz{Time: params.StartDate, Valid: true},
		Date_2: pgtype.Timestamptz{Time: params.EndDate, Valid: true},
	}

	transactions, err := t.store.ListTransactionsByUserAndDate(context, args)

	if err != nil {
		return []tm.Transaction{}, 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(transactions) == 0 {
		return []tm.Transaction{}, 0, nil
	}

	var transactionModels []tm.Transaction
	count := transactions[len(transactions)-1].TotalCount

	for _, transaction := range transactions {
		transactionModel, err := storeToModel(paginatedWithDateModelToStore(transaction))

		if err != nil {
			return []tm.Transaction{}, 0, []e.ApiError{e.StoreError{Message: err.Error()}}
		}

		transactionModels = append(transactionModels, transactionModel)
	}

	return transactionModels, count, nil
}

func (t TransactionsRepository) ReadById(context c.Context, id uuid.UUID) (tm.Transaction, []e.ApiError) {
	transaction, err := t.store.GetTransactionByID(context, id)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	model, err := storeToModel(transaction)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return model, nil
}

func (t TransactionsRepository) Update(context c.Context, transaction tm.Transaction) (tm.Transaction, []e.ApiError) {
	value, err := u.Float64ToNumeric(transaction.Value)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	params := pgs.UpdateTransactionParams{
		ID:                        transaction.ID,
		Name:                      transaction.Name,
		Date:                      pgtype.Timestamptz{Time: transaction.Date, Valid: true},
		Value:                     value,
		CategoryID:                transaction.CategoryID,
		CreditCardID:              u.UUIDToPgTypeUUID(transaction.CreditCardID),
		MonthlyTransactionsID:     u.UUIDToPgTypeUUID(transaction.MonthlyTransactionsID),
		AnnualTransactionsID:      u.UUIDToPgTypeUUID(transaction.AnnualTransactionsID),
		InstallmentTransactionsID: u.UUIDToPgTypeUUID(transaction.InstallmentTransactionsID),
	}

	transactionUpdated, err := t.store.UpdateTransaction(context, params)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	model, err := storeToModel(transactionUpdated)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return model, nil
}

func (t TransactionsRepository) Delete(context c.Context, id uuid.UUID) []e.ApiError {
	err := t.store.DeleteTransaction(context, id)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}

func paginatedModelToStore(transaction pgs.ListTransactionsByUserIDPaginatedRow) pgs.Transaction {
	return pgs.Transaction{
		ID:                        transaction.ID,
		UserID:                    transaction.UserID,
		Name:                      transaction.Name,
		Date:                      transaction.Date,
		Value:                     transaction.Value,
		CategoryID:                transaction.CategoryID,
		CreditCardID:              transaction.CreditCardID,
		MonthlyTransactionsID:     transaction.MonthlyTransactionsID,
		AnnualTransactionsID:      transaction.AnnualTransactionsID,
		InstallmentTransactionsID: transaction.InstallmentTransactionsID,
		CreatedAt:                 transaction.CreatedAt,
		UpdatedAt:                 transaction.UpdatedAt,
	}
}

func paginatedWithDateModelToStore(transaction pgs.ListTransactionsByUserAndDateRow) pgs.Transaction {
	return pgs.Transaction{
		ID:                        transaction.ID,
		UserID:                    transaction.UserID,
		Name:                      transaction.Name,
		Date:                      transaction.Date,
		Value:                     transaction.Value,
		CategoryID:                transaction.CategoryID,
		CreditCardID:              transaction.CreditCardID,
		MonthlyTransactionsID:     transaction.MonthlyTransactionsID,
		AnnualTransactionsID:      transaction.AnnualTransactionsID,
		InstallmentTransactionsID: transaction.InstallmentTransactionsID,
		CreatedAt:                 transaction.CreatedAt,
		UpdatedAt:                 transaction.UpdatedAt,
	}
}

func storeToModel(transaction pgs.Transaction) (tm.Transaction, error) {
	value, err := u.NumericToFloat64(transaction.Value)

	if err != nil {
		return tm.Transaction{}, err
	}

	return tm.Transaction{
		ID:                        transaction.ID,
		UserID:                    transaction.UserID,
		Name:                      transaction.Name,
		Date:                      transaction.Date.Time,
		Value:                     value,
		CategoryID:                transaction.CategoryID,
		CreditCardID:              u.PgTypeUUIDToUUID(transaction.CreditCardID),
		MonthlyTransactionsID:     u.PgTypeUUIDToUUID(transaction.MonthlyTransactionsID),
		AnnualTransactionsID:      u.PgTypeUUIDToUUID(transaction.AnnualTransactionsID),
		InstallmentTransactionsID: u.PgTypeUUIDToUUID(transaction.InstallmentTransactionsID),
		CreatedAt:                 transaction.CreatedAt.Time,
		UpdatedAt:                 transaction.UpdatedAt.Time,
	}, nil
}
