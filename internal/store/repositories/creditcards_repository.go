package repositories

import (
	c "context"
	e "financialcontrol/internal/models/errors"
	"financialcontrol/internal/store/dtos"
	pgs "financialcontrol/internal/store/pgstore"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"

	"github.com/google/uuid"
)

func (r Repository) CreateCreditCard(context c.Context, creditCard cm.CreateCreditCard) (cm.CreditCard, []e.ApiError) {
	param := pgs.CreateCreditCardParams{
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
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCreditcardToCreditcard(createdCreditCard), nil
}

func (r Repository) ReadCreditCards(context c.Context, userId uuid.UUID) ([]cm.CreditCard, []e.ApiError) {
	creditCards, err := r.store.ListCreditCards(context, userId)

	if err != nil {
		return nil, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(creditCards) == 0 {
		return []cm.CreditCard{}, nil
	}

	creditCardsResponse := make([]cm.CreditCard, 0, len(creditCards))

	for _, creditCard := range creditCards {
		creditCardsResponse = append(creditCardsResponse, dtos.StoreCreditcardToCreditcard(creditCard))
	}

	return creditCardsResponse, nil
}

func (r Repository) ReadCreditCardByID(context c.Context, creditCardId uuid.UUID) (cm.CreditCard, []e.ApiError) {
	creditCard, err := r.store.GetCreditCardByID(context, creditCardId)

	if err != nil {
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCreditcardToCreditcard(creditCard), nil
}

func (r Repository) ReadCountByUser(context c.Context, userId uuid.UUID) (int, []e.ApiError) {
	count, err := r.store.CountCreditCardsByUserID(context, userId)

	if err != nil {
		return 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return int(count), nil
}

func (r Repository) UpdateCreditCard(context c.Context, creditCard cm.CreditCard) (cm.CreditCard, []e.ApiError) {
	param := pgs.UpdateCreditCardParams{
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
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return dtos.StoreCreditcardToCreditcard(updatedCreditCard), nil
}

func (r Repository) DeleteCreditCard(context c.Context, creditCardId uuid.UUID) []e.ApiError {
	err := r.store.DeleteCreditCard(context, creditCardId)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}

func (r Repository) HasTransactionsByCreditCard(context c.Context, creditCardID uuid.UUID) (bool, []e.ApiError) {
	hasTransactions, err := r.store.HasTransactionsByCreditCard(context, u.UUIDToPgTypeUUID(&creditCardID))

	if err != nil {
		return false, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return hasTransactions, nil
}
