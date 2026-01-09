package repositories

import (
	"errors"
	m "financialcontrol/internal/models"
	pgs "financialcontrol/internal/store/pgstore"
	stm "financialcontrol/internal/store/storemocks"
	cm "financialcontrol/internal/v1/categories/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestCreateCategory(t *testing.T) {
	data := cm.CreateCategory{
		UserID:          uuid.New(),
		TransactionType: m.Debit,
		Name:            "Test Category",
		Icon:            "test-icon",
	}

	mock := stm.NewCategoriesStoreMock()
	mock.CategoryResult = pgs.Category{
		ID:              uuid.New(),
		UserID:          data.UserID,
		TransactionType: int32(data.TransactionType),
		Name:            data.Name,
		Icon:            data.Icon,
		CreatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	repository := NewCategoriesRepository(mock)

	model, err := repository.Create(t.Context(), data)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if model.UserID != data.UserID {
		t.Errorf("expected UserID %v, got %v", data.UserID, model.UserID)
	}

	if model.TransactionType != data.TransactionType {
		t.Errorf("expected TransactionType %v, got %v", data.TransactionType, model.TransactionType)
	}

	if model.Name != data.Name {
		t.Errorf("expected Name %v, got %v", data.Name, model.Name)
	}

	if model.Icon != data.Icon {
		t.Errorf("expected Icon %v, got %v", data.Icon, model.Icon)
	}
}

func TestCreateCategoryErrorFromStore(t *testing.T) {
	data := cm.CreateCategory{
		UserID:          uuid.New(),
		TransactionType: m.Debit,
		Name:            "Test Category",
		Icon:            "test-icon",
	}

	mock := stm.NewCategoriesStoreMock()
	mock.Error = errors.New("forcing a error...")

	repository := NewCategoriesRepository(mock)

	model, err := repository.Create(t.Context(), data)

	if err == nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if model != (cm.Category{}) {
		t.Errorf("expected empty model, got %v", model)
	}

	if err[0].String() != mock.Error.Error() {
		t.Errorf("expected error message %v, got %v", mock.Error.Error(), err[0].String())
	}
}
