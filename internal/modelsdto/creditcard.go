package modelsdto

import (
	"financialcontrol/internal/dtos"
	"financialcontrol/internal/models"

	"github.com/google/uuid"
)

func ShortCreditCardResponseFromShortCreditCard(model *models.ShortCreditCard) *dtos.ShortCreditCardResponse {
	return &dtos.ShortCreditCardResponse{
		ID:               model.ID,
		Name:             model.Name,
		FirstFourNumbers: model.FirstFourNumbers,
		Limit:            model.Limit,
		CloseDay:         model.CloseDay,
		ExpireDay:        model.ExpireDay,
		BackgroundColor:  model.BackgroundColor,
		TextColor:        model.TextColor,
	}
}

func CreateCreditCardFromCreditCardRequest(request dtos.CreditCardRequest, userID uuid.UUID) models.CreateCreditCard {
	return models.CreateCreditCard{
		UserID:           userID,
		Name:             request.Name,
		FirstFourNumbers: request.FirstFourNumbers,
		Limit:            request.Limit,
		CloseDay:         request.CloseDay,
		ExpireDay:        request.ExpireDay,
		BackgroundColor:  request.BackgroundColor,
		TextColor:        request.TextColor,
	}
}

func CreditCardResponseFromCreditCard(model models.CreditCard) dtos.CreditCardResponse {
	return dtos.CreditCardResponse{
		ID:               model.ID,
		Name:             model.Name,
		FirstFourNumbers: model.FirstFourNumbers,
		Limit:            model.Limit,
		CloseDay:         model.CloseDay,
		ExpireDay:        model.ExpireDay,
		BackgroundColor:  model.BackgroundColor,
		TextColor:        model.TextColor,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
	}
}

func ShortCreditCardResponseFromCreditCard(model models.CreditCard) dtos.ShortCreditCardResponse {
	return dtos.ShortCreditCardResponse{
		ID:               model.ID,
		Name:             model.Name,
		FirstFourNumbers: model.FirstFourNumbers,
		Limit:            model.Limit,
		CloseDay:         model.CloseDay,
		ExpireDay:        model.ExpireDay,
		BackgroundColor:  model.BackgroundColor,
		TextColor:        model.TextColor,
	}
}
