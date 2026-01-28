package repositories

import (
	"context"
	"financialcontrol/internal/models"
	"financialcontrol/internal/store/pgstore"
	"financialcontrol/internal/utils"

	"github.com/google/uuid"
)

type CreditCard interface {
	CreateCreditCard(context context.Context, creditCard models.CreateCreditCard) (models.CreditCard, error)
	ReadCreditCards(context context.Context, userId uuid.UUID) ([]models.CreditCard, error)
	ReadCountByUser(context context.Context, userId uuid.UUID) (int, error)
	ReadCreditCardByID(context context.Context, creditCardId uuid.UUID) (models.CreditCard, error)
	UpdateCreditCard(context context.Context, creditCard models.CreditCard) (models.CreditCard, error)
	DeleteCreditCard(context context.Context, creditCardId uuid.UUID) error
	HasTransactionsByCreditCard(context context.Context, creditCardID uuid.UUID) (bool, error)
}

func (r Repository) CreateCreditCard(context context.Context, creditCard models.CreateCreditCard) (models.CreditCard, error) {
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
		return models.CreditCard{}, err
	}

	return storeCreditcardToCreditcard(createdCreditCard), nil
}

func (r Repository) ReadCreditCards(context context.Context, userId uuid.UUID) ([]models.CreditCard, error) {
	creditCards, err := r.store.ListCreditCards(context, userId)

	if err != nil {
		return nil, err
	}

	if len(creditCards) == 0 {
		return []models.CreditCard{}, nil
	}

	creditCardsResponse := make([]models.CreditCard, 0, len(creditCards))

	for _, creditCard := range creditCards {
		creditCardsResponse = append(creditCardsResponse, storeCreditcardToCreditcard(creditCard))
	}

	return creditCardsResponse, nil
}

func (r Repository) ReadCreditCardByID(context context.Context, creditCardId uuid.UUID) (models.CreditCard, error) {
	creditCard, err := r.store.GetCreditCardByID(context, creditCardId)

	if err != nil {
		return models.CreditCard{}, err
	}

	return storeCreditcardToCreditcard(creditCard), nil
}

func (r Repository) ReadCountByUser(context context.Context, userId uuid.UUID) (int, error) {
	count, err := r.store.CountCreditCardsByUserID(context, userId)

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r Repository) UpdateCreditCard(context context.Context, creditCard models.CreditCard) (models.CreditCard, error) {
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
		return models.CreditCard{}, err
	}

	return storeCreditcardToCreditcard(updatedCreditCard), nil
}

func (r Repository) DeleteCreditCard(context context.Context, creditCardId uuid.UUID) error {
	err := r.store.DeleteCreditCard(context, creditCardId)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) HasTransactionsByCreditCard(context context.Context, creditCardID uuid.UUID) (bool, error) {
	hasTransactions, err := r.store.HasTransactionsByCreditCard(context, utils.UUIDToPgTypeUUID(&creditCardID))

	if err != nil {
		return false, err
	}

	return hasTransactions, nil
}

func storeCreditcardToCreditcard(storeModel pgstore.CreditCard) models.CreditCard {
	return models.CreditCard{
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
