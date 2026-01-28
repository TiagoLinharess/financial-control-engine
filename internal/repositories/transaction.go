package repositories

import (
	"context"
	"financialcontrol/internal/commonsmodels"
	"financialcontrol/internal/models"
	"financialcontrol/internal/store/pgstore"
	"financialcontrol/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Transaction interface {
	ReadCategoryByID(context context.Context, categoryID uuid.UUID) (models.Category, error)
	ReadCreditCardByID(context context.Context, creditCardId uuid.UUID) (models.CreditCard, error)
	CreateTransaction(context context.Context, transaction models.CreateTransaction) (models.ShortTransaction, error)
	ReadTransactions(context context.Context, params commonsmodels.PaginatedParams) ([]models.Transaction, int64, error)
	ReadTransactionsInToDates(context context.Context, params commonsmodels.PaginatedParamsWithDateRange) ([]models.Transaction, int64, error)
	ReadTransactionById(context context.Context, id uuid.UUID) (models.Transaction, error)
	UpdateTransaction(context context.Context, transaction models.Transaction) (models.ShortTransaction, error)
	DeleteTransaction(context context.Context, id uuid.UUID) error
	PayTransaction(context context.Context, id uuid.UUID, paid bool) error
	GetCreditcardTotalAmount(ctx context.Context, model models.TransactionsCreditCardTotal) (float64, error)
}

func (r Repository) CreateTransaction(context context.Context, transaction models.CreateTransaction) (models.ShortTransaction, error) {
	value := utils.Float64ToNumeric(transaction.Value)

	param := pgstore.CreateTransactionParams{
		UserID:                    transaction.UserID,
		Name:                      transaction.Name,
		Date:                      pgtype.Timestamptz{Time: transaction.Date, Valid: true},
		Value:                     value,
		Paid:                      transaction.Paid,
		CategoryID:                transaction.CategoryID,
		CreditCardID:              utils.UUIDToPgTypeUUID(transaction.CreditcardID),
		MonthlyTransactionsID:     utils.UUIDToPgTypeUUID(transaction.MonthlyTransactionsID),
		AnnualTransactionsID:      utils.UUIDToPgTypeUUID(transaction.AnnualTransactionsID),
		InstallmentTransactionsID: utils.UUIDToPgTypeUUID(transaction.InstallmentTransactionsID),
	}

	createdTransaction, err := r.store.CreateTransaction(context, param)

	if err != nil {
		return models.ShortTransaction{}, err
	}

	return models.ShortTransaction{
		ID:        createdTransaction.ID,
		Name:      createdTransaction.Name,
		Date:      createdTransaction.Date.Time,
		Value:     utils.NumericToFloat64(createdTransaction.Value),
		Paid:      createdTransaction.Paid,
		CreatedAt: createdTransaction.CreatedAt.Time,
		UpdatedAt: createdTransaction.UpdatedAt.Time,
	}, nil
}

func (r Repository) ReadTransactions(context context.Context, params commonsmodels.PaginatedParams) ([]models.Transaction, int64, error) {
	args := pgstore.ListTransactionsByUserIDPaginatedParams{
		UserID: params.UserID,
		Limit:  params.Limit,
		Offset: params.Offset,
	}

	transactions, err := r.store.ListTransactionsByUserIDPaginated(context, args)

	if err != nil {
		return []models.Transaction{}, 0, err
	}

	if len(transactions) == 0 {
		return []models.Transaction{}, 0, nil
	}

	var transactionModels []models.Transaction
	count := transactions[len(transactions)-1].TotalCount

	for _, transaction := range transactions {
		transactionModels = append(transactionModels, storeTransactionToTransaction(storeTransactionPaginatedToStoreTransaction(transaction)))
	}

	return transactionModels, count, nil
}

func (r Repository) ReadTransactionsInToDates(context context.Context, params commonsmodels.PaginatedParamsWithDateRange) ([]models.Transaction, int64, error) {
	args := pgstore.ListTransactionsByUserAndDateParams{
		UserID: params.UserID,
		Limit:  params.Limit,
		Offset: params.Offset,
		Date:   pgtype.Timestamptz{Time: params.StartDate, Valid: true},
		Date_2: pgtype.Timestamptz{Time: params.EndDate, Valid: true},
	}

	transactions, err := r.store.ListTransactionsByUserAndDate(context, args)

	if err != nil {
		return []models.Transaction{}, 0, err
	}

	if len(transactions) == 0 {
		return []models.Transaction{}, 0, nil
	}

	var transactionModels []models.Transaction
	count := transactions[len(transactions)-1].TotalCount

	for _, transaction := range transactions {
		transactionModels = append(transactionModels, storeTransactionToTransaction(storeTransactionListToStoreTransaction(transaction)))
	}

	return transactionModels, count, nil
}

func (r Repository) ReadTransactionById(context context.Context, id uuid.UUID) (models.Transaction, error) {
	transaction, err := r.store.GetTransactionByID(context, id)

	if err != nil {
		return models.Transaction{}, err
	}

	return storeTransactionToTransaction(transaction), nil
}

func (r Repository) UpdateTransaction(context context.Context, transaction models.Transaction) (models.ShortTransaction, error) {
	value := utils.Float64ToNumeric(transaction.Value)
	var creditCardID pgtype.UUID
	if transaction.Creditcard != nil {
		creditCardID = utils.UUIDToPgTypeUUID(&transaction.Creditcard.ID)
	}

	params := pgstore.UpdateTransactionParams{
		ID:                        transaction.ID,
		Name:                      transaction.Name,
		Date:                      pgtype.Timestamptz{Time: transaction.Date, Valid: true},
		Value:                     value,
		Paid:                      transaction.Paid,
		CategoryID:                transaction.Category.ID,
		CreditCardID:              creditCardID,
		MonthlyTransactionsID:     utils.UUIDToPgTypeUUID(transaction.MonthlyTransaction),
		AnnualTransactionsID:      utils.UUIDToPgTypeUUID(transaction.AnnualTransaction),
		InstallmentTransactionsID: utils.UUIDToPgTypeUUID(transaction.InstallmentTransaction),
	}

	transactionUpdated, err := r.store.UpdateTransaction(context, params)

	if err != nil {
		return models.ShortTransaction{}, err
	}

	return models.ShortTransaction{
		ID:        transactionUpdated.ID,
		Name:      transactionUpdated.Name,
		Date:      transactionUpdated.Date.Time,
		Value:     utils.NumericToFloat64(transactionUpdated.Value),
		Paid:      transactionUpdated.Paid,
		CreatedAt: transactionUpdated.CreatedAt.Time,
		UpdatedAt: transactionUpdated.UpdatedAt.Time,
	}, nil
}

func (r Repository) DeleteTransaction(context context.Context, id uuid.UUID) error {
	err := r.store.DeleteTransaction(context, id)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) PayTransaction(context context.Context, id uuid.UUID, paid bool) error {
	params := pgstore.PayTransactionParams{
		ID:   id,
		Paid: paid,
	}

	err := r.store.PayTransaction(context, params)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetCreditcardTotalAmount(ctx context.Context, model models.TransactionsCreditCardTotal) (float64, error) {
	normalizedDate := time.Date(model.Date.Year(), model.Date.Month(), model.Date.Day(), 0, 0, 0, 0, model.Date.Location())
	dateOnly := pgtype.Date{Time: normalizedDate, Valid: true}

	params := pgstore.GetTotalFromCreditTransactionsByUserAndMonthParams{
		Column1:      dateOnly,
		CreditCardID: utils.UUIDToPgTypeUUID(&model.CreditcardID),
		UserID:       model.UserID,
	}

	total, err := r.store.GetTotalFromCreditTransactionsByUserAndMonth(ctx, params)

	if err != nil {
		return 0, err
	}

	if !total.Valid {
		return 0, nil
	}

	floatValue := utils.NumericToFloat64(total)

	return floatValue, nil
}

func storeTransactionListToStoreTransaction(transaction pgstore.ListTransactionsByUserAndDateRow) pgstore.GetTransactionByIDRow {
	return storeTransactionPaginatedToStoreTransaction(pgstore.ListTransactionsByUserIDPaginatedRow(transaction))
}

func storeTransactionPaginatedToStoreTransaction(transaction pgstore.ListTransactionsByUserIDPaginatedRow) pgstore.GetTransactionByIDRow {
	return pgstore.GetTransactionByIDRow{
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

func storeTransactionToTransaction(transaction pgstore.GetTransactionByIDRow) models.Transaction {
	category := models.ShortCategory{
		ID:              *utils.PgTypeUUIDToUUID(transaction.CategoryID),
		TransactionType: models.TransactionType(transaction.CategoryTransactionType.Int32),
		Name:            transaction.CategoryName.String,
		Icon:            transaction.CategoryIcon.String,
	}

	var creditcard *models.ShortCreditCard
	if transaction.CreditcardID.Valid {
		creditcardValue := models.ShortCreditCard{
			ID:               *utils.PgTypeUUIDToUUID(transaction.CreditcardID),
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

	return models.Transaction{
		ID:                     transaction.ID,
		UserID:                 transaction.UserID,
		Name:                   transaction.Name,
		Date:                   transaction.Date.Time,
		Value:                  utils.NumericToFloat64(transaction.Value),
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
