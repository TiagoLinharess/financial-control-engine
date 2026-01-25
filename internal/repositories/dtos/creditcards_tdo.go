package dtos

import (
	pgs "financialcontrol/internal/store/pgstore"
	cm "financialcontrol/internal/v1/creditcards/models"
)

func StoreCreditcardToCreditcard(storeModel pgs.CreditCard) cm.CreditCard {
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
