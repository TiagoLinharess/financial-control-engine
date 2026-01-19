package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	m "financialcontrol/internal/models"
	pgs "financialcontrol/internal/store/pgstore"
	"financialcontrol/internal/utils"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/categories/models"
	cr "financialcontrol/internal/v1/creditcards/models"
	tm "financialcontrol/internal/v1/transactions/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func createTestUUID() uuid.UUID {
	return uuid.New()
}

func createTestTimestamp(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, Valid: true}
}

func createTestNumeric(value float64) pgtype.Numeric {
	return u.Float64ToNumeric(value)
}

func TestStoreDateToStore(t *testing.T) {
	userID := createTestUUID()
	categoryID := createTestUUID()
	now := time.Now()

	dateRow := pgs.ListTransactionsByUserAndDateRow{
		ID:                        createTestUUID(),
		UserID:                    userID,
		Name:                      "Test Date Transaction",
		Date:                      createTestTimestamp(now),
		Value:                     createTestNumeric(200.00),
		Paid:                      false,
		CreatedAt:                 createTestTimestamp(now),
		UpdatedAt:                 createTestTimestamp(now),
		CategoryID:                pgtype.UUID{Bytes: categoryID, Valid: true},
		CategoryTransactionType:   pgtype.Int4{Int32: 2, Valid: true},
		CategoryName:              pgtype.Text{String: "Transport", Valid: true},
		CategoryIcon:              pgtype.Text{String: "🚗", Valid: true},
		CreditcardID:              pgtype.UUID{},
		MonthlyTransactionsID:     pgtype.UUID{},
		AnnualTransactionsID:      pgtype.UUID{},
		InstallmentTransactionsID: pgtype.UUID{},
		TotalCount:                50,
	}

	result := storeDateToStore(dateRow)

	assert.Equal(t, dateRow.ID, result.ID)
	assert.Equal(t, dateRow.UserID, result.UserID)
	assert.Equal(t, dateRow.Name, result.Name)
	assert.Equal(t, dateRow.CategoryID, result.CategoryID)
	assert.Equal(t, dateRow.Value, result.Value)
}

func TestStorePaginatedToStore(t *testing.T) {
	userID := createTestUUID()
	categoryID := createTestUUID()
	creditCardID := createTestUUID()
	now := time.Now()

	paginatedRow := pgs.ListTransactionsByUserIDPaginatedRow{
		ID:                         createTestUUID(),
		UserID:                     userID,
		Name:                       "Test Transaction",
		Date:                       createTestTimestamp(now),
		Value:                      createTestNumeric(150.75),
		Paid:                       true,
		CreatedAt:                  createTestTimestamp(now),
		UpdatedAt:                  createTestTimestamp(now),
		CategoryID:                 pgtype.UUID{Bytes: categoryID, Valid: true},
		CategoryTransactionType:    pgtype.Int4{Int32: 1, Valid: true},
		CategoryName:               pgtype.Text{String: "Food", Valid: true},
		CategoryIcon:               pgtype.Text{String: "🍔", Valid: true},
		CreditcardID:               pgtype.UUID{Bytes: creditCardID, Valid: true},
		CreditcardName:             pgtype.Text{String: "Visa", Valid: true},
		CreditcardFirstFourNumbers: pgtype.Text{String: "4111", Valid: true},
		CreditcardCreditLimit:      pgtype.Float8{Float64: 100000.00, Valid: true},
		CreditcardCloseDay:         pgtype.Int4{Int32: 10, Valid: true},
		CreditcardExpireDay:        pgtype.Int4{Int32: 20, Valid: true},
		CreditcardBackgroundColor:  pgtype.Text{String: "#FFFFFF", Valid: true},
		CreditcardTextColor:        pgtype.Text{String: "#000000", Valid: true},
		MonthlyTransactionsID:      pgtype.UUID{},
		AnnualTransactionsID:       pgtype.UUID{},
		InstallmentTransactionsID:  pgtype.UUID{},
		TotalCount:                 100,
	}

	result := storePaginatedToStore(paginatedRow)

	assert.Equal(t, paginatedRow.ID, result.ID)
	assert.Equal(t, paginatedRow.UserID, result.UserID)
	assert.Equal(t, paginatedRow.Name, result.Name)
	assert.Equal(t, paginatedRow.Value, result.Value)
	assert.Equal(t, paginatedRow.Paid, result.Paid)
	assert.Equal(t, paginatedRow.CategoryID, result.CategoryID)
}

func TestStoreToModel_WithCreditCard(t *testing.T) {
	userID := createTestUUID()
	categoryID := createTestUUID()
	creditCardID := createTestUUID()
	now := time.Now()

	storeRow := pgs.GetTransactionByIDRow{
		ID:                         createTestUUID(),
		UserID:                     userID,
		Name:                       "Test Transaction",
		Date:                       createTestTimestamp(now),
		Value:                      createTestNumeric(100.50),
		Paid:                       true,
		CreatedAt:                  createTestTimestamp(now),
		UpdatedAt:                  createTestTimestamp(now),
		CategoryID:                 pgtype.UUID{Bytes: categoryID, Valid: true},
		CategoryTransactionType:    pgtype.Int4{Int32: 1, Valid: true},
		CategoryName:               pgtype.Text{String: "Food", Valid: true},
		CategoryIcon:               pgtype.Text{String: "🍔", Valid: true},
		CreditcardID:               pgtype.UUID{Bytes: creditCardID, Valid: true},
		CreditcardName:             pgtype.Text{String: "Visa", Valid: true},
		CreditcardFirstFourNumbers: pgtype.Text{String: "4111", Valid: true},
		CreditcardCreditLimit:      pgtype.Float8{Float64: 100000.00, Valid: true},
		CreditcardCloseDay:         pgtype.Int4{Int32: 10, Valid: true},
		CreditcardExpireDay:        pgtype.Int4{Int32: 20, Valid: true},
		CreditcardBackgroundColor:  pgtype.Text{String: "#FFFFFF", Valid: true},
		CreditcardTextColor:        pgtype.Text{String: "#000000", Valid: true},
	}

	result := storeToModel(storeRow)

	assert.Equal(t, storeRow.ID, result.ID)
	assert.Equal(t, storeRow.UserID, result.UserID)
	assert.Equal(t, "Test Transaction", result.Name)
	assert.Equal(t, 100.50, result.Value)
	assert.True(t, result.Paid)
	assert.NotNil(t, result.Category)
	assert.Equal(t, "Food", result.Category.Name)
	assert.Equal(t, m.Debit, result.Category.TransactionType)
	assert.NotNil(t, result.Creditcard)
	assert.Equal(t, "Visa", result.Creditcard.Name)
	assert.Equal(t, "4111", result.Creditcard.FirstFourNumbers)
	assert.Equal(t, 100000.00, result.Creditcard.Limit)
	assert.Equal(t, int32(10), result.Creditcard.CloseDay)
	assert.Equal(t, int32(20), result.Creditcard.ExpireDay)
}

func TestStoreToModel_WithoutCreditCard(t *testing.T) {
	categoryID := createTestUUID()
	now := time.Now()

	storeRow := pgs.GetTransactionByIDRow{
		ID:                      createTestUUID(),
		UserID:                  createTestUUID(),
		Name:                    "Test Transaction",
		Date:                    createTestTimestamp(now),
		Value:                   createTestNumeric(100.50),
		Paid:                    false,
		CreatedAt:               createTestTimestamp(now),
		UpdatedAt:               createTestTimestamp(now),
		CategoryID:              pgtype.UUID{Bytes: categoryID, Valid: true},
		CategoryTransactionType: pgtype.Int4{Int32: 2, Valid: true},
		CategoryName:            pgtype.Text{String: "Transport", Valid: true},
		CategoryIcon:            pgtype.Text{String: "🚗", Valid: true},
		CreditcardID:            pgtype.UUID{Valid: false},
	}

	result := storeToModel(storeRow)

	assert.Equal(t, storeRow.ID, result.ID)
	assert.NotNil(t, result.Category)
	assert.Equal(t, "Transport", result.Category.Name)
	assert.Equal(t, m.Credit, result.Category.TransactionType)
	assert.Nil(t, result.Creditcard)
	assert.Nil(t, result.MonthlyTransaction)
	assert.Nil(t, result.AnnualTransaction)
	assert.Nil(t, result.InstallmentTransaction)
	assert.False(t, result.Paid)
}

func TestStoreToModel_WithAllFields(t *testing.T) {
	userID := createTestUUID()
	categoryID := createTestUUID()
	creditCardID := createTestUUID()
	monthlyTxID := createTestUUID()
	annualTxID := createTestUUID()
	installmentTxID := createTestUUID()
	now := time.Now()

	storeRow := pgs.GetTransactionByIDRow{
		ID:                                 createTestUUID(),
		UserID:                             userID,
		Name:                               "Full Transaction",
		Date:                               createTestTimestamp(now),
		Value:                              createTestNumeric(500.00),
		Paid:                               true,
		CreatedAt:                          createTestTimestamp(now),
		UpdatedAt:                          createTestTimestamp(now),
		CategoryID:                         pgtype.UUID{Bytes: categoryID, Valid: true},
		CategoryTransactionType:            pgtype.Int4{Int32: 0, Valid: true},
		CategoryName:                       pgtype.Text{String: "Salary", Valid: true},
		CategoryIcon:                       pgtype.Text{String: "💰", Valid: true},
		CreditcardID:                       pgtype.UUID{Bytes: creditCardID, Valid: true},
		CreditcardName:                     pgtype.Text{String: "Mastercard", Valid: true},
		CreditcardFirstFourNumbers:         pgtype.Text{String: "5555", Valid: true},
		CreditcardCreditLimit:              pgtype.Float8{Float64: 50000.00, Valid: true},
		CreditcardCloseDay:                 pgtype.Int4{Int32: 15, Valid: true},
		CreditcardExpireDay:                pgtype.Int4{Int32: 25, Valid: true},
		CreditcardBackgroundColor:          pgtype.Text{String: "#FF0000", Valid: true},
		CreditcardTextColor:                pgtype.Text{String: "#FFFFFF", Valid: true},
		MonthlyTransactionsID:              pgtype.UUID{Bytes: monthlyTxID, Valid: true},
		MonthlyTransactionsDay:             pgtype.Int4{Int32: 5, Valid: true},
		AnnualTransactionsID:               pgtype.UUID{Bytes: annualTxID, Valid: true},
		AnnualTransactionsMonth:            pgtype.Int4{Int32: 12, Valid: true},
		AnnualTransactionsDay:              pgtype.Int4{Int32: 25, Valid: true},
		InstallmentTransactionsID:          pgtype.UUID{Bytes: installmentTxID, Valid: true},
		InstallmentTransactionsInitialDate: createTestTimestamp(now),
		InstallmentTransactionsFinalDate:   createTestTimestamp(now.AddDate(0, 6, 0)),
	}

	result := storeToModel(storeRow)

	assert.Equal(t, storeRow.ID, result.ID)
	assert.Equal(t, userID, result.UserID)
	assert.Equal(t, "Full Transaction", result.Name)
	assert.Equal(t, 500.00, result.Value)
	assert.True(t, result.Paid)
	assert.NotNil(t, result.Category)
	assert.Equal(t, m.Income, result.Category.TransactionType)
	assert.NotNil(t, result.Creditcard)
	assert.Equal(t, "Mastercard", result.Creditcard.Name)
}

func TestCreate_Success(t *testing.T) {
	now := time.Now()
	userID := createTestUUID()
	categoryID := createTestUUID()
	transactionID := createTestUUID()

	mockStore := &mockStore{}
	mockStore.onCreateTransaction = func(ctx context.Context, params pgs.CreateTransactionParams) (pgs.CreateTransactionRow, error) {
		return pgs.CreateTransactionRow{
			ID:        transactionID,
			Name:      "Test Transaction",
			Date:      createTestTimestamp(now),
			Value:     createTestNumeric(100.00),
			Paid:      false,
			CreatedAt: createTestTimestamp(now),
			UpdatedAt: createTestTimestamp(now),
		}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	result, errs := repo.Create(context.Background(), tm.CreateTransaction{
		UserID:                    userID,
		Name:                      "Test Transaction",
		Date:                      now,
		Value:                     100.00,
		Paid:                      false,
		CategoryID:                categoryID,
		CreditcardID:              nil,
		MonthlyTransactionsID:     nil,
		AnnualTransactionsID:      nil,
		InstallmentTransactionsID: nil,
	})

	assert.Empty(t, errs)
	assert.Equal(t, transactionID, result.ID)
	assert.Equal(t, "Test Transaction", result.Name)
	assert.Equal(t, 100.00, result.Value)
}

func TestCreate_StoreError(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onCreateTransaction = func(ctx context.Context, params pgs.CreateTransactionParams) (pgs.CreateTransactionRow, error) {
		return pgs.CreateTransactionRow{}, errors.New("database error")
	}

	repo := NewTransactionsRepository(mockStore)

	result, errs := repo.Create(context.Background(), tm.CreateTransaction{
		UserID:                    createTestUUID(),
		Name:                      "Test",
		Date:                      time.Now(),
		Value:                     50.00,
		Paid:                      false,
		CategoryID:                createTestUUID(),
		CreditcardID:              nil,
		MonthlyTransactionsID:     nil,
		AnnualTransactionsID:      nil,
		InstallmentTransactionsID: nil,
	})

	assert.NotEmpty(t, errs)
	assert.Equal(t, uuid.UUID{}, result.ID)
}

func TestRead_Success(t *testing.T) {
	userID := createTestUUID()
	now := time.Now()
	categoryID := createTestUUID()
	creditCardID := createTestUUID()

	mockStore := &mockStore{}
	mockStore.onListTransactionsByUserIDPaginated = func(ctx context.Context, params pgs.ListTransactionsByUserIDPaginatedParams) ([]pgs.ListTransactionsByUserIDPaginatedRow, error) {
		return []pgs.ListTransactionsByUserIDPaginatedRow{
			{
				ID:                         createTestUUID(),
				UserID:                     userID,
				Name:                       "Test Transaction",
				Date:                       createTestTimestamp(now),
				Value:                      createTestNumeric(100.00),
				Paid:                       true,
				CreatedAt:                  createTestTimestamp(now),
				UpdatedAt:                  createTestTimestamp(now),
				CategoryID:                 pgtype.UUID{Bytes: categoryID, Valid: true},
				CategoryTransactionType:    pgtype.Int4{Int32: 1, Valid: true},
				CategoryName:               pgtype.Text{String: "Food", Valid: true},
				CategoryIcon:               pgtype.Text{String: "🍔", Valid: true},
				CreditcardID:               pgtype.UUID{Bytes: creditCardID, Valid: true},
				CreditcardName:             pgtype.Text{String: "Visa", Valid: true},
				CreditcardFirstFourNumbers: pgtype.Text{String: "4111", Valid: true},
				CreditcardCreditLimit:      pgtype.Float8{Float64: 10000.00, Valid: true},
				CreditcardCloseDay:         pgtype.Int4{Int32: 10, Valid: true},
				CreditcardExpireDay:        pgtype.Int4{Int32: 20, Valid: true},
				CreditcardBackgroundColor:  pgtype.Text{String: "#FFFFFF", Valid: true},
				CreditcardTextColor:        pgtype.Text{String: "#000000", Valid: true},
				TotalCount:                 1,
			},
		}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	result, count, errs := repo.Read(context.Background(), m.PaginatedParams{
		UserID: userID,
		Limit:  10,
		Offset: 0,
	})

	assert.Empty(t, errs)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), count)
	assert.Equal(t, "Test Transaction", result[0].Name)
}

func TestRead_EmptyResult(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onListTransactionsByUserIDPaginated = func(ctx context.Context, params pgs.ListTransactionsByUserIDPaginatedParams) ([]pgs.ListTransactionsByUserIDPaginatedRow, error) {
		return []pgs.ListTransactionsByUserIDPaginatedRow{}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	result, count, errs := repo.Read(context.Background(), m.PaginatedParams{
		UserID: createTestUUID(),
		Limit:  10,
		Offset: 0,
	})

	assert.Empty(t, errs)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), count)
}

func TestReadById_Success(t *testing.T) {
	transactionID := createTestUUID()
	userID := createTestUUID()
	categoryID := createTestUUID()
	creditCardID := createTestUUID()
	now := time.Now()

	mockStore := &mockStore{}
	mockStore.onGetTransactionByID = func(ctx context.Context, id uuid.UUID) (pgs.GetTransactionByIDRow, error) {
		return pgs.GetTransactionByIDRow{
			ID:                         transactionID,
			UserID:                     userID,
			Name:                       "Test Transaction",
			Date:                       createTestTimestamp(now),
			Value:                      createTestNumeric(100.50),
			Paid:                       true,
			CreatedAt:                  createTestTimestamp(now),
			UpdatedAt:                  createTestTimestamp(now),
			CategoryID:                 pgtype.UUID{Bytes: categoryID, Valid: true},
			CategoryTransactionType:    pgtype.Int4{Int32: 1, Valid: true},
			CategoryName:               pgtype.Text{String: "Food", Valid: true},
			CategoryIcon:               pgtype.Text{String: "🍔", Valid: true},
			CreditcardID:               pgtype.UUID{Bytes: creditCardID, Valid: true},
			CreditcardName:             pgtype.Text{String: "Visa", Valid: true},
			CreditcardFirstFourNumbers: pgtype.Text{String: "4111", Valid: true},
			CreditcardCreditLimit:      pgtype.Float8{Float64: 100000.00, Valid: true},
			CreditcardCloseDay:         pgtype.Int4{Int32: 10, Valid: true},
			CreditcardExpireDay:        pgtype.Int4{Int32: 20, Valid: true},
			CreditcardBackgroundColor:  pgtype.Text{String: "#FFFFFF", Valid: true},
			CreditcardTextColor:        pgtype.Text{String: "#000000", Valid: true},
		}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	result, errs := repo.ReadById(context.Background(), transactionID)

	assert.Empty(t, errs)
	assert.Equal(t, transactionID, result.ID)
	assert.NotNil(t, result.Creditcard)
}

func TestReadById_NotFound(t *testing.T) {
	transactionID := createTestUUID()

	mockStore := &mockStore{}
	mockStore.onGetTransactionByID = func(ctx context.Context, id uuid.UUID) (pgs.GetTransactionByIDRow, error) {
		return pgs.GetTransactionByIDRow{}, errors.New("no rows in result set")
	}

	repo := NewTransactionsRepository(mockStore)

	result, errs := repo.ReadById(context.Background(), transactionID)

	assert.NotEmpty(t, errs)
	assert.Equal(t, uuid.UUID{}, result.ID)
}

func TestRead_StoreError(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onListTransactionsByUserIDPaginated = func(ctx context.Context, params pgs.ListTransactionsByUserIDPaginatedParams) ([]pgs.ListTransactionsByUserIDPaginatedRow, error) {
		return nil, errors.New("database error")
	}

	repo := NewTransactionsRepository(mockStore)

	result, count, errs := repo.Read(context.Background(), m.PaginatedParams{
		UserID: createTestUUID(),
		Limit:  10,
		Offset: 0,
	})

	assert.NotEmpty(t, errs)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), count)
}

func TestDelete_Success(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onDeleteTransaction = func(ctx context.Context, id uuid.UUID) error {
		return nil
	}

	repo := NewTransactionsRepository(mockStore)

	errs := repo.Delete(context.Background(), createTestUUID())

	assert.Empty(t, errs)
}

func TestDelete_Error(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onDeleteTransaction = func(ctx context.Context, id uuid.UUID) error {
		return errors.New("delete failed")
	}

	repo := NewTransactionsRepository(mockStore)

	errs := repo.Delete(context.Background(), createTestUUID())

	assert.NotEmpty(t, errs)
}

func TestPay_Success(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onPayTransaction = func(ctx context.Context, params pgs.PayTransactionParams) error {
		return nil
	}

	repo := NewTransactionsRepository(mockStore)

	errs := repo.Pay(context.Background(), createTestUUID(), true)

	assert.Empty(t, errs)
}

func TestReadInToDates_Success(t *testing.T) {
	userID := createTestUUID()
	now := time.Now()
	startDate := now.AddDate(0, -1, 0)
	endDate := now
	categoryID := createTestUUID()

	mockStore := &mockStore{}
	mockStore.onListTransactionsByUserAndDate = func(ctx context.Context, params pgs.ListTransactionsByUserAndDateParams) ([]pgs.ListTransactionsByUserAndDateRow, error) {
		return []pgs.ListTransactionsByUserAndDateRow{
			{
				ID:                         createTestUUID(),
				UserID:                     userID,
				Name:                       "Transaction in range",
				Date:                       createTestTimestamp(now),
				Value:                      createTestNumeric(75.50),
				Paid:                       false,
				CreatedAt:                  createTestTimestamp(now),
				UpdatedAt:                  createTestTimestamp(now),
				CategoryID:                 pgtype.UUID{Bytes: categoryID, Valid: true},
				CategoryTransactionType:    pgtype.Int4{Int32: 0, Valid: true},
				CategoryName:               pgtype.Text{String: "Income", Valid: true},
				CategoryIcon:               pgtype.Text{String: "📈", Valid: true},
				CreditcardID:               pgtype.UUID{Valid: false},
				CreditcardName:             pgtype.Text{Valid: false},
				CreditcardFirstFourNumbers: pgtype.Text{Valid: false},
				CreditcardCreditLimit:      pgtype.Float8{Valid: false},
				CreditcardCloseDay:         pgtype.Int4{Valid: false},
				CreditcardExpireDay:        pgtype.Int4{Valid: false},
				CreditcardBackgroundColor:  pgtype.Text{Valid: false},
				CreditcardTextColor:        pgtype.Text{Valid: false},
				TotalCount:                 1,
			},
		}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	result, count, errs := repo.ReadInToDates(context.Background(), m.PaginatedParamsWithDateRange{
		UserID:    userID,
		StartDate: startDate,
		EndDate:   endDate,
		Limit:     10,
		Offset:    0,
	})

	assert.Empty(t, errs)
	assert.Len(t, result, 1)
	assert.Equal(t, int64(1), count)
	assert.Equal(t, "Transaction in range", result[0].Name)
}

func TestReadInToDates_EmptyResult(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onListTransactionsByUserAndDate = func(ctx context.Context, params pgs.ListTransactionsByUserAndDateParams) ([]pgs.ListTransactionsByUserAndDateRow, error) {
		return []pgs.ListTransactionsByUserAndDateRow{}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	result, count, errs := repo.ReadInToDates(context.Background(), m.PaginatedParamsWithDateRange{
		UserID:    createTestUUID(),
		StartDate: time.Now().AddDate(-1, 0, 0),
		EndDate:   time.Now(),
		Limit:     10,
		Offset:    0,
	})

	assert.Empty(t, errs)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), count)
}

func TestReadInToDates_StoreError(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onListTransactionsByUserAndDate = func(ctx context.Context, params pgs.ListTransactionsByUserAndDateParams) ([]pgs.ListTransactionsByUserAndDateRow, error) {
		return nil, errors.New("database error")
	}

	repo := NewTransactionsRepository(mockStore)

	result, count, errs := repo.ReadInToDates(context.Background(), m.PaginatedParamsWithDateRange{
		UserID:    createTestUUID(),
		StartDate: time.Now().AddDate(-1, 0, 0),
		EndDate:   time.Now(),
		Limit:     10,
		Offset:    0,
	})

	assert.NotEmpty(t, errs)
	assert.Empty(t, result)
	assert.Equal(t, int64(0), count)
}

func TestUpdate_Success(t *testing.T) {
	now := time.Now()
	transactionID := createTestUUID()

	mockStore := &mockStore{}
	mockStore.onUpdateTransaction = func(ctx context.Context, params pgs.UpdateTransactionParams) (pgs.Transaction, error) {
		return pgs.Transaction{
			ID:        transactionID,
			Name:      "Updated Transaction",
			Date:      createTestTimestamp(now),
			Value:     createTestNumeric(200.00),
			Paid:      true,
			CreatedAt: createTestTimestamp(now),
			UpdatedAt: createTestTimestamp(now),
		}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	creditCardID := createTestUUID()

	result, errs := repo.Update(context.Background(), tm.Transaction{
		ID:         transactionID,
		UserID:     createTestUUID(),
		Name:       "Updated Transaction",
		Date:       now,
		Value:      200.00,
		Paid:       true,
		Category:   cm.ShortCategory{},
		Creditcard: &cr.ShortCreditCard{ID: creditCardID},
	})

	assert.Empty(t, errs)
	assert.Equal(t, transactionID, result.ID)
	assert.Equal(t, "Updated Transaction", result.Name)
}

func TestUpdate_StoreError(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onUpdateTransaction = func(ctx context.Context, params pgs.UpdateTransactionParams) (pgs.Transaction, error) {
		return pgs.Transaction{}, errors.New("update failed")
	}

	repo := NewTransactionsRepository(mockStore)

	result, errs := repo.Update(context.Background(), tm.Transaction{
		ID:       createTestUUID(),
		UserID:   createTestUUID(),
		Name:     "Test",
		Date:     time.Now(),
		Value:    100.00,
		Paid:     false,
		Category: cm.ShortCategory{},
	})

	assert.NotEmpty(t, errs)
	assert.Equal(t, uuid.UUID{}, result.ID)
}

func TestPay_Error(t *testing.T) {
	mockStore := &mockStore{}
	mockStore.onPayTransaction = func(ctx context.Context, params pgs.PayTransactionParams) error {
		return errors.New("pay failed")
	}

	repo := NewTransactionsRepository(mockStore)

	errs := repo.Pay(context.Background(), createTestUUID(), false)

	assert.NotEmpty(t, errs)
}

func TestGetCreditcardTotalAmount_Success(t *testing.T) {
	userID := createTestUUID()
	creditcardID := createTestUUID()
	now := time.Now()

	mockStore := &mockStore{}
	mockStore.onGetCreditcardTotalAmount = func(ctx context.Context, arg pgs.GetTotalFromCreditTransactionsByUserAndMonthParams) (pgtype.Numeric, error) {
		return utils.Float64ToNumeric(250.50), nil
	}

	repo := NewTransactionsRepository(mockStore)

	total, err := repo.GetCreditcardTotalAmount(context.Background(), tm.TransactionsCreditCardTotal{
		Date:         now,
		UserID:       userID,
		CreditcardID: creditcardID,
	})

	assert.Nil(t, err)
	assert.Equal(t, 250.50, total)
}

func TestGetCreditcardTotalAmount_StoreError(t *testing.T) {
	userID := createTestUUID()
	creditcardID := createTestUUID()
	now := time.Now()

	mockStore := &mockStore{}
	mockStore.onGetCreditcardTotalAmount = func(ctx context.Context, arg pgs.GetTotalFromCreditTransactionsByUserAndMonthParams) (pgtype.Numeric, error) {
		return pgtype.Numeric{}, errors.New("database error")
	}

	repo := NewTransactionsRepository(mockStore)

	total, err := repo.GetCreditcardTotalAmount(context.Background(), tm.TransactionsCreditCardTotal{
		Date:         now,
		UserID:       userID,
		CreditcardID: creditcardID,
	})

	assert.NotNil(t, err)
	assert.Equal(t, float64(0), total)
}

func TestGetCreditcardTotalAmount_ConversionError(t *testing.T) {
	userID := createTestUUID()
	creditcardID := createTestUUID()
	now := time.Now()

	mockStore := &mockStore{}
	mockStore.onGetCreditcardTotalAmount = func(ctx context.Context, arg pgs.GetTotalFromCreditTransactionsByUserAndMonthParams) (pgtype.Numeric, error) {
		// Return invalid Numeric to simulate conversion error
		return pgtype.Numeric{Valid: false}, nil
	}

	repo := NewTransactionsRepository(mockStore)

	total, err := repo.GetCreditcardTotalAmount(context.Background(), tm.TransactionsCreditCardTotal{
		Date:         now,
		UserID:       userID,
		CreditcardID: creditcardID,
	})

	assert.Nil(t, err)
	assert.Equal(t, float64(0), total)
}

func TestGetCreditcardTotalAmount_Error(t *testing.T) {
	userID := createTestUUID()
	creditcardID := createTestUUID()
	now := time.Now()

	mockStore := &mockStore{}
	mockStore.onGetCreditcardTotalAmount = func(ctx context.Context, arg pgs.GetTotalFromCreditTransactionsByUserAndMonthParams) (pgtype.Numeric, error) {
		return pgtype.Numeric{}, errors.New("database error")
	}

	repo := NewTransactionsRepository(mockStore)

	total, err := repo.GetCreditcardTotalAmount(context.Background(), tm.TransactionsCreditCardTotal{
		Date:         now,
		UserID:       userID,
		CreditcardID: creditcardID,
	})

	assert.NotNil(t, err)
	assert.Equal(t, float64(0), total)
}

type mockStore struct {
	onCreateTransaction                 func(ctx context.Context, params pgs.CreateTransactionParams) (pgs.CreateTransactionRow, error)
	onListTransactionsByUserIDPaginated func(ctx context.Context, params pgs.ListTransactionsByUserIDPaginatedParams) ([]pgs.ListTransactionsByUserIDPaginatedRow, error)
	onListTransactionsByUserAndDate     func(ctx context.Context, params pgs.ListTransactionsByUserAndDateParams) ([]pgs.ListTransactionsByUserAndDateRow, error)
	onGetTransactionByID                func(ctx context.Context, id uuid.UUID) (pgs.GetTransactionByIDRow, error)
	onUpdateTransaction                 func(ctx context.Context, params pgs.UpdateTransactionParams) (pgs.Transaction, error)
	onDeleteTransaction                 func(ctx context.Context, id uuid.UUID) error
	onPayTransaction                    func(ctx context.Context, params pgs.PayTransactionParams) error
	onGetCreditcardTotalAmount          func(ctx context.Context, arg pgs.GetTotalFromCreditTransactionsByUserAndMonthParams) (pgtype.Numeric, error)
}

func (m *mockStore) GetTotalFromCreditTransactionsByUserAndMonth(ctx context.Context, arg pgs.GetTotalFromCreditTransactionsByUserAndMonthParams) (pgtype.Numeric, error) {
	if m.onGetCreditcardTotalAmount != nil {
		return m.onGetCreditcardTotalAmount(ctx, arg)
	}
	return pgtype.Numeric{}, nil
}

func (m *mockStore) CreateTransaction(ctx context.Context, params pgs.CreateTransactionParams) (pgs.CreateTransactionRow, error) {
	if m.onCreateTransaction != nil {
		return m.onCreateTransaction(ctx, params)
	}
	return pgs.CreateTransactionRow{}, nil
}

func (m *mockStore) ListTransactionsByUserIDPaginated(ctx context.Context, params pgs.ListTransactionsByUserIDPaginatedParams) ([]pgs.ListTransactionsByUserIDPaginatedRow, error) {
	if m.onListTransactionsByUserIDPaginated != nil {
		return m.onListTransactionsByUserIDPaginated(ctx, params)
	}
	return nil, nil
}

func (m *mockStore) ListTransactionsByUserAndDate(ctx context.Context, params pgs.ListTransactionsByUserAndDateParams) ([]pgs.ListTransactionsByUserAndDateRow, error) {
	if m.onListTransactionsByUserAndDate != nil {
		return m.onListTransactionsByUserAndDate(ctx, params)
	}
	return nil, nil
}

func (m *mockStore) GetTransactionByID(ctx context.Context, id uuid.UUID) (pgs.GetTransactionByIDRow, error) {
	if m.onGetTransactionByID != nil {
		return m.onGetTransactionByID(ctx, id)
	}
	return pgs.GetTransactionByIDRow{}, nil
}

func (m *mockStore) UpdateTransaction(ctx context.Context, params pgs.UpdateTransactionParams) (pgs.Transaction, error) {
	if m.onUpdateTransaction != nil {
		return m.onUpdateTransaction(ctx, params)
	}
	return pgs.Transaction{}, nil
}

func (m *mockStore) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	if m.onDeleteTransaction != nil {
		return m.onDeleteTransaction(ctx, id)
	}
	return nil
}

func (m *mockStore) PayTransaction(ctx context.Context, params pgs.PayTransactionParams) error {
	if m.onPayTransaction != nil {
		return m.onPayTransaction(ctx, params)
	}
	return nil
}
