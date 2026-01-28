package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"financialcontrol/internal/commonsmodels"
	"financialcontrol/internal/constants"
	"financialcontrol/internal/dtos"
	e "financialcontrol/internal/errors"
	"financialcontrol/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Helper Functions
func setupTestContext(t *testing.T, method, path string, body interface{}) *gin.Context {
	var bodyBytes []byte
	if body != nil {
		var err error
		bodyBytes, err = json.Marshal(body)
		if err != nil {
			t.Fatalf("failed to marshal body: %v", err)
		}
	}

	req, _ := http.NewRequest(method, path, bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	return ctx
}

// ============= EXTRA TESTS MERGED FROM transactions_service_extra_test.go =============

func TestGetRelationsInvalidJSON(t *testing.T) {
	svc := transaction{
		repository: NewTransactionsRepositoryMock(),
	}

	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString("{invalid-json"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	setupContextWithCookie(ctx, uuid.New())

	_, status, errs := svc.getRelations(ctx)

	if status != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) == 0 {
		t.Fatalf("expected errors, got none")
	}
}

func TestGetRelationsCategoryInternalError(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Test",
		Date:       time.Now(),
		Value:      10,
		Paid:       false,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryError = errors.New("db error")

	svc := transaction{
		repository: mock,
	}

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := svc.getRelations(ctx)

	if status != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Fatalf("expected errors, got none")
	}
}

func TestGetRelationsCreditcardTotalAmountError(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test",
		Date:         time.Now(),
		Value:        100,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Credit,
	}
	mock.CreditcardResult = models.CreditCard{
		ID:     creditcardID,
		UserID: userID,
		Limit:  10000,
	}
	mock.Error = errors.New("get total amount failed")

	svc := transaction{
		repository: mock,
	}

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := svc.getRelations(ctx)

	if status != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Fatalf("expected errors, got none")
	}
}

func TestGetRelationsDebitWithCreditcard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test",
		Date:         time.Now(),
		Value:        100,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
	}
	mock.CreditcardResult = models.CreditCard{
		ID:     creditcardID,
		UserID: userID,
		Limit:  10000,
	}

	svc := transaction{
		repository: mock,
	}

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := svc.getRelations(ctx)

	if status != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) == 0 {
		t.Fatalf("expected errors, got none")
	}
}

func TestReadInvalidUUID(t *testing.T) {
	svc := transaction{}

	ctx := getTestContextForGet(t, "/transactions/not-a-uuid")
	setupContextWithCookie(ctx, uuid.New())
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "not-a-uuid"})

	_, status, errs := svc.read(ctx)

	if status != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) != 0 {
		t.Fatalf("expected no errors, got %v", errs)
	}
}

func TestReadRepositoryInternalError(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.Error = errors.New("db error")

	svc := transaction{
		repository: mock,
	}

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	setupContextWithCookie(ctx, userID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: transactionID.String()})

	_, status, errs := svc.read(ctx)

	if status != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Fatalf("expected errors, got none")
	}
}

func TestReadDatesFromInvalidStart(t *testing.T) {
	svc := transaction{}

	_, _, errs := svc.readDatesFrom("invalid-date", "2021-01-01")

	if len(errs) == 0 {
		t.Fatalf("expected errors for invalid start date, got none")
	}
}

func TestReadDatesFromInvalidEnd(t *testing.T) {
	svc := transaction{}

	_, _, errs := svc.readDatesFrom("2021-01-01", "invalid-date")

	if len(errs) == 0 {
		t.Fatalf("expected errors for invalid end date, got none")
	}
}

func TestReadTransactionsLimitTooLarge(t *testing.T) {
	userID := uuid.New()

	transactions := []models.Transaction{{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      "Tx",
		Date:      time.Now(),
		Value:     10,
		Paid:      false,
		Category:  models.ShortCategory{ID: uuid.New(), Name: "Cat"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionsResult = transactions
	mock.TransactionsCount = 1

	svc := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=100&page=1")
	setupContextWithCookie(ctx, userID)

	resp, status, errs := svc.Read(ctx)

	if status != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, status)
	}

	if len(errs) != 0 {
		t.Fatalf("expected no errors, got %v", errs)
	}

	if len(resp.Items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(resp.Items))
	}
}

func TestReadTransactionsPageZero(t *testing.T) {
	userID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.TransactionsResult = []models.Transaction{}
	mock.TransactionsCount = 0

	svc := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=10&page=0")
	setupContextWithCookie(ctx, userID)

	resp, status, errs := svc.Read(ctx)

	if status != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, status)
	}

	if len(errs) != 0 {
		t.Fatalf("expected no errors, got %v", errs)
	}

	if resp.Page != 1 {
		t.Fatalf("expected page 1, got %d", resp.Page)
	}
}

func TestUpdateTransactionWithCreditcard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()
	transactionID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Updated",
		Date:         time.Now(),
		Value:        123,
		Paid:         true,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{ID: categoryID, UserID: userID, TransactionType: models.Credit}
	mock.CreditcardResult = models.CreditCard{ID: creditcardID, UserID: userID, Limit: 10000}
	mock.TransactionFullResult = models.Transaction{ID: transactionID, UserID: userID}
	mock.TransactionResult = models.ShortTransaction{ID: transactionID}

	svc := NewTransactionsService(mock)

	ctx := setupTestContext(t, "PUT", "/transactions/"+transactionID.String(), request)
	setupContextWithCookie(ctx, userID)
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}

	resp, status, errs := svc.Update(ctx)

	if status != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, status)
	}

	if len(errs) != 0 {
		t.Fatalf("expected no errors, got %v", errs)
	}

	if resp.ID != transactionID {
		t.Fatalf("expected id %v, got %v", transactionID, resp.ID)
	}
}

func TestUpdateTransactionUpdateError(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Updated",
		Date:       time.Now(),
		Value:      123,
		Paid:       true,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{ID: categoryID, UserID: userID, TransactionType: models.Debit}
	mock.TransactionFullResult = models.Transaction{ID: transactionID, UserID: userID}
	mock.UpdateError = errors.New("update failed")

	svc := NewTransactionsService(mock)

	ctx := setupTestContext(t, "PUT", "/transactions/"+transactionID.String(), request)
	setupContextWithCookie(ctx, userID)
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}

	_, status, errs := svc.Update(ctx)

	if status != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Fatalf("expected errors, got none")
	}
}

func TestReadTransactionsRepositoryError(t *testing.T) {
	userID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.Error = errors.New("db error")

	svc := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=10&page=1")
	setupContextWithCookie(ctx, userID)

	_, status, errs := svc.Read(ctx)

	if status != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Fatalf("expected errors, got none")
	}
}

type TransactionsRepositoryMock struct {
	Error                 error
	CategoryError         error
	CreditcardError       error
	DeleteError           error
	PayError              error
	UpdateError           error
	TransactionResult     models.ShortTransaction
	TransactionFullResult models.Transaction
	TransactionsResult    []models.Transaction
	TransactionsCount     int64
	Amount                float64
	CategoryResult        models.Category
	CreditcardResult      models.CreditCard
}

func NewTransactionsRepositoryMock() *TransactionsRepositoryMock {
	return &TransactionsRepositoryMock{
		Error:                 nil,
		CategoryError:         nil,
		CreditcardError:       nil,
		UpdateError:           nil,
		TransactionResult:     models.ShortTransaction{},
		TransactionFullResult: models.Transaction{},
		TransactionsResult:    []models.Transaction{},
		TransactionsCount:     0,
		Amount:                0,
		CategoryResult:        models.Category{},
		CreditcardResult:      models.CreditCard{},
	}
}

func (t *TransactionsRepositoryMock) ReadCategoryByID(ctx context.Context, categoryID uuid.UUID) (models.Category, []e.ApiError) {
	if t.CategoryError != nil {
		return models.Category{}, []e.ApiError{e.CustomError{Message: t.CategoryError.Error()}}
	}
	return t.CategoryResult, nil
}

func (t *TransactionsRepositoryMock) ReadCreditCardByID(ctx context.Context, creditCardId uuid.UUID) (models.CreditCard, []e.ApiError) {
	if t.CreditcardError != nil {
		return models.CreditCard{}, []e.ApiError{e.CustomError{Message: t.CreditcardError.Error()}}
	}
	return t.CreditcardResult, nil
}

func (t *TransactionsRepositoryMock) GetCreditcardTotalAmount(ctx context.Context, model models.TransactionsCreditCardTotal) (float64, error) {
	if t.Error != nil {
		return 0, t.Error
	}
	return t.Amount, nil
}

func (t *TransactionsRepositoryMock) CreateTransaction(ctx context.Context, transaction models.CreateTransaction) (models.ShortTransaction, []e.ApiError) {
	if t.Error != nil {
		return models.ShortTransaction{}, []e.ApiError{e.CustomError{Message: t.Error.Error()}}
	}
	return t.TransactionResult, nil
}

func (t *TransactionsRepositoryMock) ReadTransactions(ctx context.Context, params commonsmodels.PaginatedParams) ([]models.Transaction, int64, []e.ApiError) {
	if t.Error != nil {
		return []models.Transaction{}, 0, []e.ApiError{e.CustomError{Message: t.Error.Error()}}
	}
	return t.TransactionsResult, t.TransactionsCount, nil
}

func (t *TransactionsRepositoryMock) ReadTransactionsInToDates(ctx context.Context, params commonsmodels.PaginatedParamsWithDateRange) ([]models.Transaction, int64, []e.ApiError) {
	if t.Error != nil {
		return []models.Transaction{}, 0, []e.ApiError{e.CustomError{Message: t.Error.Error()}}
	}
	return t.TransactionsResult, t.TransactionsCount, nil
}

func (t *TransactionsRepositoryMock) ReadTransactionById(ctx context.Context, id uuid.UUID) (models.Transaction, []e.ApiError) {
	if t.Error != nil {
		return models.Transaction{}, []e.ApiError{e.CustomError{Message: t.Error.Error()}}
	}
	return t.TransactionFullResult, nil
}

func (t *TransactionsRepositoryMock) UpdateTransaction(ctx context.Context, transaction models.Transaction) (models.ShortTransaction, []e.ApiError) {
	if t.UpdateError != nil {
		return models.ShortTransaction{}, []e.ApiError{e.CustomError{Message: t.UpdateError.Error()}}
	}
	if t.Error != nil {
		return models.ShortTransaction{}, []e.ApiError{e.CustomError{Message: t.Error.Error()}}
	}
	return t.TransactionResult, nil
}

func (t *TransactionsRepositoryMock) DeleteTransaction(ctx context.Context, id uuid.UUID) []e.ApiError {
	if t.DeleteError != nil {
		return []e.ApiError{e.CustomError{Message: t.DeleteError.Error()}}
	}
	if t.Error != nil {
		return []e.ApiError{e.CustomError{Message: t.Error.Error()}}
	}
	return nil
}

func (t *TransactionsRepositoryMock) PayTransaction(ctx context.Context, id uuid.UUID, paid bool) []e.ApiError {
	if t.PayError != nil {
		return []e.ApiError{e.CustomError{Message: t.PayError.Error()}}
	}
	if t.Error != nil {
		return []e.ApiError{e.CustomError{Message: t.Error.Error()}}
	}
	return nil
}

func setupContextWithCookie(ctx *gin.Context, userID uuid.UUID) {
	req := ctx.Request
	req.AddCookie(&http.Cookie{
		Name:  "user_id",
		Value: userID.String(),
	})
}

func getTestContextForGet(t *testing.T, path string) *gin.Context {
	req, _ := http.NewRequest("GET", path, nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	return ctx
}

// ============= CREATE TESTS =============

func TestCreateTransactionSuccess(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	transactionID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Test Transaction",
		Date:       time.Now(),
		Value:      100.00,
		Paid:       false,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
		Name:            "Test Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.TransactionResult = models.ShortTransaction{
		ID:        transactionID,
		Name:      request.Name,
		Date:      request.Date,
		Value:     request.Value,
		Paid:      request.Paid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	response, status, errs := service.Create(ctx)

	if status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}

	if response.ID != transactionID {
		t.Errorf("expected ID %v, got %v", transactionID, response.ID)
	}
}

func TestCreateTransactionMissingUserCookie(t *testing.T) {
	request := dtos.TransactionRequest{
		Name:       "Test Transaction",
		Date:       time.Now(),
		Value:      100.00,
		Paid:       false,
		CategoryID: uuid.New(),
	}

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)

	_, status, errs := service.Create(ctx)

	if status != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionCategoryNotFound(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Test Transaction",
		Date:       time.Now(),
		Value:      100.00,
		Paid:       false,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryError = errors.New(string(constants.StoreErrorNoRowsMsg))

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionCreditRequiresCreditcard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Test Transaction",
		Date:       time.Now(),
		Value:      100.00,
		Paid:       false,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Credit,
		Name:            "Credit Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, _ := service.Create(ctx)

	if status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}
}

func TestCreateTransactionDebitCannotHaveCreditcard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test Transaction",
		Date:         time.Now(),
		Value:        100.00,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
		Name:            "Debit Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.CreditcardResult = models.CreditCard{
		ID:     creditcardID,
		UserID: userID,
		Name:   "Test Card",
	}

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, _ := service.Create(ctx)

	if status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}
}

func TestCreateTransactionWithValidCreditcard(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()
	transactionID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test Transaction",
		Date:         time.Now(),
		Value:        100.00,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Credit,
		Name:            "Credit Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.CreditcardResult = models.CreditCard{
		ID:     creditcardID,
		UserID: userID,
		Name:   "Test Card",
		Limit:  10000,
	}
	mock.TransactionResult = models.ShortTransaction{
		ID:        transactionID,
		Name:      request.Name,
		Date:      request.Date,
		Value:     request.Value,
		Paid:      request.Paid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	response, status, errs := service.Create(ctx)

	if status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}

	if response.ID != transactionID {
		t.Errorf("expected ID %v, got %v", transactionID, response.ID)
	}
}

// ============= READ TESTS =============

func TestReadTransactionsSuccess(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	transactionID1 := uuid.New()
	transactionID2 := uuid.New()

	transactions := []models.Transaction{
		{
			ID:        transactionID1,
			UserID:    userID,
			Name:      "Transaction 1",
			Date:      time.Now(),
			Value:     100.00,
			Paid:      false,
			Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        transactionID2,
			UserID:    userID,
			Name:      "Transaction 2",
			Date:      time.Now(),
			Value:     200.00,
			Paid:      true,
			Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionsResult = transactions
	mock.TransactionsCount = 2

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=10&page=1")
	setupContextWithCookie(ctx, userID)

	response, status, errs := service.Read(ctx)

	if status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}

	if len(response.Items) != 2 {
		t.Errorf("expected 2 items, got %d", len(response.Items))
	}
}

func TestReadTransactionsMissingUserCookie(t *testing.T) {
	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=10&page=1")

	_, status, errs := service.Read(ctx)

	if status != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestReadTransactionsInvalidPage(t *testing.T) {
	userID := uuid.New()

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=10&page=invalid")
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Read(ctx)

	if status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestReadTransactionsWithDateRange(t *testing.T) {
	userID := uuid.New()
	now := time.Now()

	transactions := []models.Transaction{
		{
			ID:        uuid.New(),
			UserID:    userID,
			Name:      "Transaction 1",
			Date:      now,
			Value:     100.00,
			Paid:      false,
			Category:  models.ShortCategory{ID: uuid.New(), Name: "Category"},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionsResult = transactions
	mock.TransactionsCount = 1

	service := NewTransactionsService(mock)

	startDate := now.Format(time.DateOnly)
	endDate := now.Format(time.DateOnly)
	ctx := getTestContextForGet(t, "/transactions?limit=10&page=1&start_date="+startDate+"&end_date="+endDate)
	setupContextWithCookie(ctx, userID)

	response, status, errs := service.Read(ctx)

	if status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}

	if len(response.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(response.Items))
	}
}

// ============= READ BY ID TESTS =============

func TestReadByIdTransactionSuccess(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	transaction := models.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Name:      "Test Transaction",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = transaction

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	response, status, errs := service.ReadById(ctx)

	if status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}

	if response.ID != transactionID {
		t.Errorf("expected ID %v, got %v", transactionID, response.ID)
	}
}

func TestReadByIdTransactionNotFound(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.Error = errors.New(string(constants.StoreErrorNoRowsMsg))

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.ReadById(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

// ============= UPDATE TESTS =============

func TestUpdateTransactionSuccess(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Updated Transaction",
		Date:       time.Now(),
		Value:      200.00,
		Paid:       true,
		CategoryID: categoryID,
	}

	originalTransaction := models.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Name:      "Original Transaction",
		Date:      time.Now().AddDate(0, 0, -1),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
		Name:            "Test Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.TransactionFullResult = originalTransaction
	mock.TransactionResult = models.ShortTransaction{
		ID:        transactionID,
		Name:      request.Name,
		Date:      request.Date,
		Value:     request.Value,
		Paid:      request.Paid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "PUT", "/transactions/"+transactionID.String(), request)
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	response, status, errs := service.Update(ctx)

	if status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}

	if response.ID != transactionID {
		t.Errorf("expected ID %v, got %v", transactionID, response.ID)
	}
}

func TestUpdateTransactionNotFound(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Updated Transaction",
		Date:       time.Now(),
		Value:      200.00,
		Paid:       true,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
		Name:            "Test Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.Error = errors.New(string(constants.StoreErrorNoRowsMsg))

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "PUT", "/transactions/"+transactionID.String(), request)
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Update(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

// ============= DELETE TESTS =============

func TestDeleteTransactionSuccess(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	transaction := models.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Name:      "Test Transaction",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = transaction

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	status, errs := service.Delete(ctx)

	if status != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}
}

func TestDeleteTransactionNotFound(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.Error = errors.New(string(constants.StoreErrorNoRowsMsg))

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	status, errs := service.Delete(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

// ============= PAY TESTS =============

func TestPayTransactionSuccess(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	transaction := models.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Name:      "Test Transaction",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = transaction

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String()+"/pay")
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	status, errs := service.Pay(ctx)

	if status != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}
}

func TestPayTransactionNotFound(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.Error = errors.New(string(constants.StoreErrorNoRowsMsg))

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String()+"/pay")
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	status, errs := service.Pay(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestPayTransactionToggleStatus(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	transaction := models.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Name:      "Test Transaction",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      true,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = transaction

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String()+"/pay")
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	status, errs := service.Pay(ctx)

	if status != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, status)
	}

	if len(errs) != 0 {
		t.Errorf("expected no errors, got %v", errs)
	}
}

// ============= CREATE ERROR TESTS =============

func TestCreateTransactionRepositoryError(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Test Transaction",
		Date:       time.Now(),
		Value:      100.00,
		Paid:       false,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
		Name:            "Test Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.Error = errors.New("database error")

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionInvalidJSON(t *testing.T) {
	userID := uuid.New()

	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	setupContextWithCookie(ctx, userID)

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	_, status, errs := service.Create(ctx)

	if status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionCategoryRepositoryError(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Test Transaction",
		Date:       time.Now(),
		Value:      100.00,
		Paid:       false,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryError = errors.New("database error")

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionCreditcardRepositoryError(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test Transaction",
		Date:         time.Now(),
		Value:        100.00,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Credit,
		Name:            "Credit Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.CreditcardError = errors.New("database error")

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestReadTransactionsInvalidStartDate(t *testing.T) {
	userID := uuid.New()

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=10&page=1&start_date=invalid&end_date=2024-01-01")
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Read(ctx)

	if status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestUpdateTransactionRepositoryError(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Updated Transaction",
		Date:       time.Now(),
		Value:      200.00,
		Paid:       true,
		CategoryID: categoryID,
	}

	originalTransaction := models.Transaction{
		ID:        transactionID,
		UserID:    userID,
		Name:      "Original",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
		Name:            "Test Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.TransactionFullResult = originalTransaction
	mock.UpdateError = errors.New("database error")

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "PUT", "/transactions/"+transactionID.String(), request)
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Update(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestReadTransactionsInvalidEndDate(t *testing.T) {
	userID := uuid.New()

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions?limit=10&page=1&start_date=2024-01-01&end_date=invalid")
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Read(ctx)

	if status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestReadByIdTransactionRepositoryError(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.Error = errors.New("database error")

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.ReadById(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionWrongCategoryUser(t *testing.T) {
	userID := uuid.New()
	otherUserID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Test Transaction",
		Date:       time.Now(),
		Value:      100.00,
		Paid:       false,
		CategoryID: categoryID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          otherUserID,
		TransactionType: models.Debit,
		Name:            "Test Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionWrongCreditcardUser(t *testing.T) {
	userID := uuid.New()
	otherUserID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test Transaction",
		Date:         time.Now(),
		Value:        100.00,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Credit,
		Name:            "Credit Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.CreditcardResult = models.CreditCard{
		ID:     creditcardID,
		UserID: otherUserID,
		Name:   "Test Card",
	}

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestUpdateTransactionWrongUser(t *testing.T) {
	userID := uuid.New()
	otherUserID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Updated",
		Date:       time.Now(),
		Value:      200.00,
		Paid:       true,
		CategoryID: categoryID,
	}

	transaction := models.Transaction{
		ID:        transactionID,
		UserID:    otherUserID,
		Name:      "Original",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Debit,
		Name:            "Test Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.TransactionFullResult = transaction

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "PUT", "/transactions/"+transactionID.String(), request)
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Update(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestDeleteTransactionWrongUser(t *testing.T) {
	userID := uuid.New()
	otherUserID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	transaction := models.Transaction{
		ID:        transactionID,
		UserID:    otherUserID,
		Name:      "Test",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = transaction

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	status, errs := service.Delete(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestPayTransactionWrongUser(t *testing.T) {
	userID := uuid.New()
	otherUserID := uuid.New()
	transactionID := uuid.New()
	categoryID := uuid.New()

	transaction := models.Transaction{
		ID:        transactionID,
		UserID:    otherUserID,
		Name:      "Test",
		Date:      time.Now(),
		Value:     100.00,
		Paid:      false,
		Category:  models.ShortCategory{ID: categoryID, Name: "Category"},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = transaction

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String()+"/pay")
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	status, errs := service.Pay(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestUpdateTransactionInvalidJSON(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	req, _ := http.NewRequest("PUT", "/transactions/"+transactionID.String(), bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}
	setupContextWithCookie(ctx, userID)

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	_, status, errs := service.Update(ctx)

	if status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestDeleteTransactionMissingUserCookie(t *testing.T) {
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}

	status, errs := service.Delete(ctx)

	if status != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestPayTransactionMissingUserCookie(t *testing.T) {
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String()+"/pay")
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}

	status, errs := service.Pay(ctx)

	if status != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestUpdateTransactionMissingUserCookie(t *testing.T) {
	transactionID := uuid.New()

	request := dtos.TransactionRequest{
		Name:       "Updated",
		Date:       time.Now(),
		Value:      200.00,
		Paid:       true,
		CategoryID: uuid.New(),
	}

	mock := NewTransactionsRepositoryMock()

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "PUT", "/transactions/"+transactionID.String(), request)
	ctx.Params = gin.Params{{Key: "id", Value: transactionID.String()}}

	_, status, errs := service.Update(ctx)

	if status != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestReadCreditcardError(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test Transaction",
		Date:         time.Now(),
		Value:        100.00,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Credit,
		Name:            "Credit Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.CreditcardError = errors.New("database error")

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestCreateTransactionReadCreditcardNotFound(t *testing.T) {
	userID := uuid.New()
	categoryID := uuid.New()
	creditcardID := uuid.New()

	request := dtos.TransactionRequest{
		Name:         "Test Transaction",
		Date:         time.Now(),
		Value:        100.00,
		Paid:         false,
		CategoryID:   categoryID,
		CreditcardID: &creditcardID,
	}

	mock := NewTransactionsRepositoryMock()
	mock.CategoryResult = models.Category{
		ID:              categoryID,
		UserID:          userID,
		TransactionType: models.Credit,
		Name:            "Credit Category",
		Icon:            "icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mock.CreditcardError = errors.New(string(constants.StoreErrorNoRowsMsg))

	service := NewTransactionsService(mock)

	ctx := setupTestContext(t, "POST", "/transactions", request)
	setupContextWithCookie(ctx, userID)

	_, status, errs := service.Create(ctx)

	if status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestDeleteTransactionRepositoryError(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = models.Transaction{
		ID:     transactionID,
		UserID: userID,
	}
	mock.DeleteError = errors.New("delete failed")

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	setupContextWithCookie(ctx, userID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: transactionID.String()})

	status, errs := service.Delete(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}
}

func TestPayTransactionRepositoryError(t *testing.T) {
	userID := uuid.New()
	transactionID := uuid.New()

	mock := NewTransactionsRepositoryMock()
	mock.TransactionFullResult = models.Transaction{
		ID:     transactionID,
		UserID: userID,
	}
	mock.PayError = errors.New("pay failed")

	service := NewTransactionsService(mock)

	ctx := getTestContextForGet(t, "/transactions/"+transactionID.String())
	setupContextWithCookie(ctx, userID)
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: transactionID.String()})

	status, errs := service.Pay(ctx)

	if status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, status)
	}

	if len(errs) == 0 {
		t.Errorf("expected errors, got none")
	}

}
