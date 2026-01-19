package repositories

import (
	c "context"
	e "financialcontrol/internal/models/errors"
	s "financialcontrol/internal/store"
	pgs "financialcontrol/internal/store/pgstore"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"

	"github.com/google/uuid"
)

type CreditCardsRepository struct {
	store s.CreditCardsStore
}

func NewCreditCardsRepository(store s.CreditCardsStore) cm.CreditCardsRepository {
	return CreditCardsRepository{store: store}
}

func (c CreditCardsRepository) Create(context c.Context, creditCard cm.CreateCreditCard) (cm.CreditCard, []e.ApiError) {
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

	createdCreditCard, err := c.store.CreateCreditCard(context, param)

	if err != nil {
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return storeModelToModel(createdCreditCard), nil
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
		creditCardsResponse = append(creditCardsResponse, storeModelToModel(creditCard))
	}

	return creditCardsResponse, nil
}

func (c CreditCardsRepository) ReadByID(context c.Context, creditCardId uuid.UUID) (cm.CreditCard, []e.ApiError) {
	creditCard, err := c.store.GetCreditCardByID(context, creditCardId)

	if err != nil {
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return storeModelToModel(creditCard), nil
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
		BackgroundColor:  creditCard.BackgroundColor,
		TextColor:        creditCard.TextColor,
	}

	updatedCreditCard, err := c.store.UpdateCreditCard(context, param)

	if err != nil {
		return cm.CreditCard{}, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return storeModelToModel(updatedCreditCard), nil
}

func (c CreditCardsRepository) Delete(context c.Context, creditCardId uuid.UUID) []e.ApiError {
	err := c.store.DeleteCreditCard(context, creditCardId)

	if err != nil {
		return []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return nil
}

func (c CreditCardsRepository) HasTransactionsByCreditCard(context c.Context, creditCardID uuid.UUID) (bool, []e.ApiError) {
	hasTransactions, err := c.store.HasTransactionsByCreditCard(context, u.UUIDToPgTypeUUID(&creditCardID))

	if err != nil {
		return false, []e.ApiError{e.StoreError{Message: err.Error()}}
	}

	return hasTransactions, nil
}

func storeModelToModel(storeModel pgs.CreditCard) cm.CreditCard {
	return cm.CreditCard{
		ID:               storeModel.ID,
		UserID:           storeModel.UserID,
		Name:             storeModel.Name,
		FirstFourNumbers: storeModel.FirstFourNumbers,
		Limit:            storeModel.CreditLimit,
		CloseDay:         storeModel.CloseDay,
		ExpireDay:        storeModel.ExpireDay,
		BackgroundColor:  storeModel.BackgroundColor,
		TextColor:        storeModel.TextColor,
		CreatedAt:        storeModel.CreatedAt.Time,
		UpdatedAt:        storeModel.UpdatedAt.Time,
	}
}
