package repositories

import (
	c "context"
	m "financialcontrol/internal/models"
	e "financialcontrol/internal/models/errors"
	st "financialcontrol/internal/store"
	pgs "financialcontrol/internal/store/pgstore"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
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

func (t TransactionsRepository) Create(context c.Context, transaction tm.CreateTransaction) (tm.ShortTransaction, []e.ApiError) {
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

	createdTransaction, err := t.store.CreateTransaction(context, param)

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
		transactionModels = append(transactionModels, storeToModel(storePaginatedToStore(transaction)))
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
		transactionModels = append(transactionModels, storeToModel(storeDateToStore(transaction)))
	}

	return transactionModels, count, nil
}

func (t TransactionsRepository) ReadById(context c.Context, id uuid.UUID) (tm.Transaction, []e.ApiError) {
	transaction, err := t.store.GetTransactionByID(context, id)

	if err != nil {
		return tm.Transaction{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return storeToModel(transaction), nil
}

func (t TransactionsRepository) Update(context c.Context, transaction tm.Transaction) (tm.ShortTransaction, []e.ApiError) {
	value := u.Float64ToNumeric(transaction.Value)

	params := pgs.UpdateTransactionParams{
		ID:                        transaction.ID,
		Name:                      transaction.Name,
		Date:                      pgtype.Timestamptz{Time: transaction.Date, Valid: true},
		Value:                     value,
		Paid:                      transaction.Paid,
		CategoryID:                transaction.Category.ID,
		CreditCardID:              u.UUIDToPgTypeUUID(&transaction.Creditcard.ID),
		MonthlyTransactionsID:     u.UUIDToPgTypeUUID(transaction.MonthlyTransaction),
		AnnualTransactionsID:      u.UUIDToPgTypeUUID(transaction.AnnualTransaction),
		InstallmentTransactionsID: u.UUIDToPgTypeUUID(transaction.InstallmentTransaction),
	}

	transactionUpdated, err := t.store.UpdateTransaction(context, params)

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

func (t TransactionsRepository) Delete(context c.Context, id uuid.UUID) []e.ApiError {
	err := t.store.DeleteTransaction(context, id)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}

func (t TransactionsRepository) Pay(context c.Context, id uuid.UUID, paid bool) []e.ApiError {
	params := pgs.PayTransactionParams{
		ID:   id,
		Paid: paid,
	}

	err := t.store.PayTransaction(context, params)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}

func storeDateToStore(transaction pgs.ListTransactionsByUserAndDateRow) pgs.GetTransactionByIDRow {
	return storePaginatedToStore(pgs.ListTransactionsByUserIDPaginatedRow(transaction))
}

func storePaginatedToStore(transaction pgs.ListTransactionsByUserIDPaginatedRow) pgs.GetTransactionByIDRow {
	return pgs.GetTransactionByIDRow{
		ID:                                 transaction.ID,
		UserID:                             transaction.UserID,
		Name:                               transaction.Name,
		Date:                               transaction.Date,
		Value:                              transaction.Value,
		Paid:                               transaction.Paid,
		CreatedAt:                          transaction.CreatedAt,
		UpdatedAt:                          transaction.UpdatedAt,
		CategoryID:                         transaction.CategoryID,
		CategoryTransactionType:            transaction.CategoryTransactionType,
		CategoryName:                       transaction.CategoryName,
		CategoryIcon:                       transaction.CategoryIcon,
		CreditcardID:                       transaction.CreditcardID,
		CreditcardName:                     transaction.CreditcardName,
		CreditcardFirstFourNumbers:         transaction.CreditcardFirstFourNumbers,
		CreditcardCreditLimit:              transaction.CreditcardCreditLimit,
		CreditcardCloseDay:                 transaction.CreditcardCloseDay,
		CreditcardExpireDay:                transaction.CreditcardExpireDay,
		CreditcardBackgroundColor:          transaction.CreditcardBackgroundColor,
		CreditcardTextColor:                transaction.CreditcardTextColor,
		MonthlyTransactionsID:              transaction.MonthlyTransactionsID,
		MonthlyTransactionsDay:             transaction.MonthlyTransactionsDay,
		AnnualTransactionsID:               transaction.AnnualTransactionsID,
		AnnualTransactionsMonth:            transaction.AnnualTransactionsMonth,
		AnnualTransactionsDay:              transaction.AnnualTransactionsDay,
		InstallmentTransactionsID:          transaction.InstallmentTransactionsID,
		InstallmentTransactionsInitialDate: transaction.InstallmentTransactionsInitialDate,
		InstallmentTransactionsFinalDate:   transaction.InstallmentTransactionsFinalDate,
	}
}

func storeToModel(transaction pgs.GetTransactionByIDRow) tm.Transaction {
	category := cm.ShortCategory{
		ID:              *u.PgTypeUUIDToUUID(transaction.CategoryID),
		TransactionType: m.TransactionType(transaction.CategoryTransactionType.Int32),
		Name:            transaction.CategoryName.String,
		Icon:            transaction.CategoryIcon.String,
	}

	var creditcard *cr.ShortCreditCard
	if transaction.CreditcardID.Valid {
		creditcardValue := cr.ShortCreditCard{
			ID:               *u.PgTypeUUIDToUUID(transaction.CreditcardID),
			Name:             transaction.CreditcardName.String,
			FirstFourNumbers: transaction.CreditcardFirstFourNumbers.String,
			Limit:            transaction.CreditcardCreditLimit.Float64,
			CloseDay:         transaction.CreditcardCloseDay.Int32,
			ExpireDay:        transaction.CreditcardExpireDay.Int32,
			BackgroundColor:  transaction.CreditcardBackgroundColor.String,
			TextColor:        transaction.CreditcardTextColor.String,
		}

		creditcard = &creditcardValue
	}

	return tm.Transaction{
		ID:                     transaction.ID,
		UserID:                 transaction.UserID,
		Name:                   transaction.Name,
		Date:                   transaction.Date.Time,
		Value:                  u.NumericToFloat64(transaction.Value),
		Paid:                   transaction.Paid,
		Category:               category,
		Creditcard:             creditcard,
		MonthlyTransaction:     nil,
		AnnualTransaction:      nil,
		InstallmentTransaction: nil,
		CreatedAt:              transaction.CreatedAt.Time,
		UpdatedAt:              transaction.UpdatedAt.Time,
	}
}
