package repositories

import (
	c "context"
	e "financialcontrol/internal/models/errors"
	s "financialcontrol/internal/store"
	pgs "financialcontrol/internal/store/pgstore"
	cm "financialcontrol/internal/v1/creditcards/models"

	"github.com/google/uuid"
)

type CreditCardsRepository struct {
	store s.CreditCardsStore
}

func NewCreditCardsRepository(store s.CreditCardsStore) cm.CreditCardsRepository {
	return CreditCardsRepository{store: store}
}

func (c CreditCardsRepository) Create(context c.Context, creditCard cm.CreditCard) (cm.CreditCard, []e.ApiError) {
	param := pgs.CreateCreditCardParams{
		UserID:           creditCard.UserID,
		Name:             creditCard.Name,
		FirstFourNumbers: creditCard.FirstFourNumbers,
		CreditLimit:      creditCard.Limit,
		CloseDay:         creditCard.CloseDay,
		ExpireDay:        creditCard.ExpireDay,
	}

	createdCreditCard, err := c.store.CreateCreditCard(context, param)

	if err != nil {
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return cm.CreditCard{
		ID:               createdCreditCard.ID,
		UserID:           createdCreditCard.UserID,
		Name:             createdCreditCard.Name,
		FirstFourNumbers: createdCreditCard.FirstFourNumbers,
		Limit:            createdCreditCard.CreditLimit,
		CloseDay:         createdCreditCard.CloseDay,
		ExpireDay:        createdCreditCard.ExpireDay,
		CreatedAt:        createdCreditCard.CreatedAt.Time,
		UpdatedAt:        createdCreditCard.UpdatedAt.Time,
	}, nil
}

func (c CreditCardsRepository) Read(context c.Context, userId uuid.UUID) ([]cm.CreditCard, []e.ApiError) {
	creditCards, err := c.store.ListCreditCards(context, userId)

	if err != nil {
		return nil, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	if len(creditCards) == 0 {
		return []cm.CreditCard{}, nil
	}

	creditCardsResponse := make([]cm.CreditCard, 0, len(creditCards))

	for _, creditCard := range creditCards {
		creditCardsResponse = append(creditCardsResponse, cm.CreditCard{
			ID:               creditCard.ID,
			UserID:           creditCard.UserID,
			Name:             creditCard.Name,
			FirstFourNumbers: creditCard.FirstFourNumbers,
			Limit:            creditCard.CreditLimit,
			CloseDay:         creditCard.CloseDay,
			ExpireDay:        creditCard.ExpireDay,
			CreatedAt:        creditCard.CreatedAt.Time,
			UpdatedAt:        creditCard.UpdatedAt.Time,
		})
	}

	return creditCardsResponse, nil
}

func (c CreditCardsRepository) ReadByID(context c.Context, creditCardId uuid.UUID) (cm.CreditCard, []e.ApiError) {
	creditCard, err := c.store.GetCreditCardByID(context, creditCardId)

	if err != nil {
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return cm.CreditCard{
		ID:               creditCard.ID,
		UserID:           creditCard.UserID,
		Name:             creditCard.Name,
		FirstFourNumbers: creditCard.FirstFourNumbers,
		Limit:            creditCard.CreditLimit,
		CloseDay:         creditCard.CloseDay,
		ExpireDay:        creditCard.ExpireDay,
		CreatedAt:        creditCard.CreatedAt.Time,
		UpdatedAt:        creditCard.UpdatedAt.Time,
	}, nil
}

func (c CreditCardsRepository) ReadCountByUser(context c.Context, userId uuid.UUID) (int, []e.ApiError) {
	count, err := c.store.CountCreditCardsByUserID(context, userId)

	if err != nil {
		return 0, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return int(count), nil
}

func (c CreditCardsRepository) Update(context c.Context, creditCard cm.CreditCard) (cm.CreditCard, []e.ApiError) {
	param := pgs.UpdateCreditCardParams{
		ID:               creditCard.ID,
		Name:             creditCard.Name,
		FirstFourNumbers: creditCard.FirstFourNumbers,
		CreditLimit:      creditCard.Limit,
		CloseDay:         creditCard.CloseDay,
		ExpireDay:        creditCard.ExpireDay,
	}

	updatedCreditCard, err := c.store.UpdateCreditCard(context, param)

	if err != nil {
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return cm.CreditCard{
		ID:               updatedCreditCard.ID,
		UserID:           updatedCreditCard.UserID,
		Name:             updatedCreditCard.Name,
		FirstFourNumbers: updatedCreditCard.FirstFourNumbers,
		Limit:            updatedCreditCard.CreditLimit,
		CloseDay:         updatedCreditCard.CloseDay,
		ExpireDay:        updatedCreditCard.ExpireDay,
		CreatedAt:        updatedCreditCard.CreatedAt.Time,
		UpdatedAt:        updatedCreditCard.UpdatedAt.Time,
	}, nil
}

func (c CreditCardsRepository) Delete(context c.Context, creditCardId uuid.UUID) []e.ApiError {
	err := c.store.DeleteCreditCard(context, creditCardId)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}
