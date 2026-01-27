package repositories

import (
	"context"
	"financialcontrol/internal/errors"
	"financialcontrol/internal/models"
	"financialcontrol/internal/repositories/dtos"
	"financialcontrol/internal/store/pgstore"
	"financialcontrol/internal/utils"

	"github.com/google/uuid"
)

type CreditCard interface {
	CreateCreditCard(context context.Context, creditCard models.CreateCreditCard) (models.CreditCard, []errors.ApiError)
	ReadCreditCards(context context.Context, userId uuid.UUID) ([]models.CreditCard, []errors.ApiError)
	ReadCountByUser(context context.Context, userId uuid.UUID) (int, []errors.ApiError)
	ReadCreditCardByID(context context.Context, creditCardId uuid.UUID) (models.CreditCard, []errors.ApiError)
	UpdateCreditCard(context context.Context, creditCard models.CreditCard) (models.CreditCard, []errors.ApiError)
	DeleteCreditCard(context context.Context, creditCardId uuid.UUID) []errors.ApiError
	HasTransactionsByCreditCard(context context.Context, creditCardID uuid.UUID) (bool, []errors.ApiError)
}

func (r Repository) CreateCreditCard(context context.Context, creditCard models.CreateCreditCard) (models.CreditCard, []errors.ApiError) {
	param := pgstore.CreateCreditCardParams{
		UserID:           creditCard.UserID,
		Name:             creditCard.Name,
		FirstFourNumbers: creditCard.FirstFourNumbers,
		CreditLimit:      creditCard.Limit,
		CloseDay:         creditCard.CloseDay,
		ExpireDay:        creditCard.ExpireDay,
		BackgroundColor:  creditCard.BackgroundColor,
		TextColor:        creditCard.TextColor,
	}

	createdCreditCard, err := r.store.CreateCreditCard(context, param)

	if err != nil {
		return models.CreditCard{}, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCreditcardToCreditcard(createdCreditCard), nil
}

func (r Repository) ReadCreditCards(context context.Context, userId uuid.UUID) ([]models.CreditCard, []errors.ApiError) {
	creditCards, err := r.store.ListCreditCards(context, userId)

	if err != nil {
		return nil, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	if len(creditCards) == 0 {
		return []models.CreditCard{}, nil
	}

	creditCardsResponse := make([]models.CreditCard, 0, len(creditCards))

	for _, creditCard := range creditCards {
		creditCardsResponse = append(creditCardsResponse, dtos.StoreCreditcardToCreditcard(creditCard))
	}

	return creditCardsResponse, nil
}

func (r Repository) ReadCreditCardByID(context context.Context, creditCardId uuid.UUID) (models.CreditCard, []errors.ApiError) {
	creditCard, err := r.store.GetCreditCardByID(context, creditCardId)

	if err != nil {
		return models.CreditCard{}, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCreditcardToCreditcard(creditCard), nil
}

func (r Repository) ReadCountByUser(context context.Context, userId uuid.UUID) (int, []errors.ApiError) {
	count, err := r.store.CountCreditCardsByUserID(context, userId)

	if err != nil {
		return 0, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return int(count), nil
}

func (r Repository) UpdateCreditCard(context context.Context, creditCard models.CreditCard) (models.CreditCard, []errors.ApiError) {
	param := pgstore.UpdateCreditCardParams{
		ID:               creditCard.ID,
		Name:             creditCard.Name,
		FirstFourNumbers: creditCard.FirstFourNumbers,
		CreditLimit:      creditCard.Limit,
		CloseDay:         creditCard.CloseDay,
		ExpireDay:        creditCard.ExpireDay,
		BackgroundColor:  creditCard.BackgroundColor,
		TextColor:        creditCard.TextColor,
	}

	updatedCreditCard, err := r.store.UpdateCreditCard(context, param)

	if err != nil {
		return models.CreditCard{}, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCreditcardToCreditcard(updatedCreditCard), nil
}

func (r Repository) DeleteCreditCard(context context.Context, creditCardId uuid.UUID) []errors.ApiError {
	err := r.store.DeleteCreditCard(context, creditCardId)

	if err != nil {
		return []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return nil
}

func (r Repository) HasTransactionsByCreditCard(context context.Context, creditCardID uuid.UUID) (bool, []errors.ApiError) {
	hasTransactions, err := r.store.HasTransactionsByCreditCard(context, utils.UUIDToPgTypeUUID(&creditCardID))

	if err != nil {
		return false, []errors.ApiError{errors.StoreError{Message: err.Error()}}
	}

	return hasTransactions, nil
}
