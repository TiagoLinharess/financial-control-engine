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

func TestReadCategoriesByUserId(t *testing.T) {
	userID := uuid.New()

	mock := stm.NewCategoriesStoreMock()
	mock.CategoriesResult = []pgs.Category{
		{
			ID:              uuid.New(),
			UserID:          userID,
			TransactionType: int32(m.Debit),
			Name:            "Category 1",
			Icon:            "icon-1",
			CreatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
			UpdatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
		},
		{
			ID:              uuid.New(),
			UserID:          userID,
			TransactionType: int32(m.Income),
			Name:            "Category 2",
			Icon:            "icon-2",
			CreatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
			UpdatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
		},
	}

	repository := NewCategoriesRepository(mock)

	categories, err := repository.Read(t.Context(), userID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(categories) != len(mock.CategoriesResult) {
		t.Fatalf("expected %d categories, got %d", len(mock.CategoriesResult), len(categories))
	}

	for i, category := range categories {
		expected := mock.CategoriesResult[i]

		if category.ID != expected.ID {
			t.Errorf("expected ID %v, got %v", expected.ID, category.ID)
		}

		if category.UserID != expected.UserID {
			t.Errorf("expected UserID %v, got %v", expected.UserID, category.UserID)
		}

		if int32(category.TransactionType) != expected.TransactionType {
			t.Errorf("expected TransactionType %v, got %v", expected.TransactionType, category.TransactionType)
		}

		if category.Name != expected.Name {
			t.Errorf("expected Name %v, got %v", expected.Name, category.Name)
		}

		if category.Icon != expected.Icon {
			t.Errorf("expected Icon %v, got %v", expected.Icon, category.Icon)
		}
	}
}

func TestReadCategoriesByUserIdEmpty(t *testing.T) {
	userID := uuid.New()

	mock := stm.NewCategoriesStoreMock()

	repository := NewCategoriesRepository(mock)

	categories, err := repository.Read(t.Context(), userID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(categories) != 0 {
		t.Fatalf("expected 0 categories, got %d", len(categories))
	}
}

func TestReadCategoriesByUserIdError(t *testing.T) {
	userID := uuid.New()

	mock := stm.NewCategoriesStoreMock()
	mock.Error = errors.New("forcing a error...")

	repository := NewCategoriesRepository(mock)

	categories, err := repository.Read(t.Context(), userID)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if len(categories) != 0 {
		t.Errorf("expected empty categories, got %v", categories)
	}

	if err[0].String() != mock.Error.Error() {
		t.Errorf("expected error message %v, got %v", mock.Error.Error(), err[0].String())
	}
}

func TestReadById(t *testing.T) {
	categoryID := uuid.New()
	data := pgs.Category{
		ID:              categoryID,
		UserID:          uuid.New(),
		TransactionType: 1,
		Name:            "Test Category",
		Icon:            "test-icon",
		CreatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	mock := stm.NewCategoriesStoreMock()
	mock.CategoryResult = data

	repository := NewCategoriesRepository(mock)

	category, err := repository.ReadByID(t.Context(), categoryID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if category.ID != data.ID {
		t.Errorf("expected ID %v, got %v", data.ID, category.ID)
	}

	if category.UserID != data.UserID {
		t.Errorf("expected UserID %v, got %v", data.UserID, category.UserID)
	}

	if int32(category.TransactionType) != data.TransactionType {
		t.Errorf("expected TransactionType %v, got %v", data.TransactionType, category.TransactionType)
	}

	if category.Name != data.Name {
		t.Errorf("expected Name %v, got %v", data.Name, category.Name)
	}

	if category.Icon != data.Icon {
		t.Errorf("expected Icon %v, got %v", data.Icon, category.Icon)
	}
}

func TestReadByIdError(t *testing.T) {
	categoryID := uuid.New()

	mock := stm.NewCategoriesStoreMock()
	mock.Error = errors.New("forcing a error...")

	repository := NewCategoriesRepository(mock)

	category, err := repository.ReadByID(t.Context(), categoryID)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if category != (cm.Category{}) {
		t.Errorf("expected empty category, got %v", category)
	}

	if err[0].String() != mock.Error.Error() {
		t.Errorf("expected error message %v, got %v", mock.Error.Error(), err[0].String())
	}
}

func TestCountByUser(t *testing.T) {
	userID := uuid.New()

	mock := stm.NewCategoriesStoreMock()
	mock.CategoriesCount = 5

	repository := NewCategoriesRepository(mock)

	count, err := repository.GetCountByUser(t.Context(), userID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if count != mock.CategoriesCount {
		t.Errorf("expected count %d, got %d", mock.CategoriesCount, count)
	}
}

func TestCountByUserError(t *testing.T) {
	userID := uuid.New()

	mock := stm.NewCategoriesStoreMock()
	mock.Error = errors.New("forcing a error...")

	repository := NewCategoriesRepository(mock)

	count, err := repository.GetCountByUser(t.Context(), userID)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}

	if err[0].String() != mock.Error.Error() {
		t.Errorf("expected error message %v, got %v", mock.Error.Error(), err[0].String())
	}
}

func TestUpdateCategory(t *testing.T) {
	data := cm.Category{
		ID:              uuid.New(),
		UserID:          uuid.New(),
		TransactionType: 1,
		Name:            "Test Category",
		Icon:            "test-icon",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	mock := stm.NewCategoriesStoreMock()
	mock.CategoryResult = pgs.Category{
		ID:              data.ID,
		UserID:          data.UserID,
		TransactionType: int32(data.TransactionType),
		Name:            data.Name,
		Icon:            data.Icon,
		CreatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:       pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	repository := NewCategoriesRepository(mock)
	updatedCategory, err := repository.Update(t.Context(), data)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updatedCategory.ID != data.ID {
		t.Errorf("expected ID %v, got %v", data.ID, updatedCategory.ID)
	}

	if updatedCategory.UserID != data.UserID {
		t.Errorf("expected UserID %v, got %v", data.UserID, updatedCategory.UserID)
	}

	if updatedCategory.TransactionType != data.TransactionType {
		t.Errorf("expected TransactionType %v, got %v", data.TransactionType, updatedCategory.TransactionType)
	}

	if updatedCategory.Name != data.Name {
		t.Errorf("expected Name %v, got %v", data.Name, updatedCategory.Name)
	}

	if updatedCategory.Icon != data.Icon {
		t.Errorf("expected Icon %v, got %v", data.Icon, updatedCategory.Icon)
	}
}

func TestUpdateCategoryError(t *testing.T) {
	data := cm.Category{}
	mock := stm.NewCategoriesStoreMock()
	mock.Error = errors.New("forcing a error...")

	repository := NewCategoriesRepository(mock)
	updatedCategory, err := repository.Update(t.Context(), data)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if updatedCategory != (cm.Category{}) {
		t.Errorf("expected empty category, got %v", updatedCategory)
	}

	if err[0].String() != mock.Error.Error() {
		t.Errorf("expected error message %v, got %v", mock.Error.Error(), err[0].String())
	}
}

func TestDeleteCategory(t *testing.T) {
	categoryID := uuid.New()

	mock := stm.NewCategoriesStoreMock()

	repository := NewCategoriesRepository(mock)
	err := repository.Delete(t.Context(), categoryID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestDeleteCategoryError(t *testing.T) {
	categoryID := uuid.New()

	mock := stm.NewCategoriesStoreMock()
	mock.Error = errors.New("forcing a error...")

	repository := NewCategoriesRepository(mock)
	err := repository.Delete(t.Context(), categoryID)

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}

	if err[0].String() != mock.Error.Error() {
		t.Errorf("expected error message %v, got %v", mock.Error.Error(), err[0].String())
	}
}
