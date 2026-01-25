package services

import (
	"financialcontrol/internal/constants"
	e "financialcontrol/internal/models/errors"
	sm "financialcontrol/internal/store/models"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreditCardsService struct {
	repository sm.CreditCardsRepository
}

func NewCreditCardsService(repository sm.CreditCardsRepository) cm.CreditCardsService {
	return &CreditCardsService{repository: repository}
}

func (c CreditCardsService) read(ctx *gin.Context) (cm.CreditCard, int, []e.ApiError) {
	creditcardNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.CreditcardNotFound}}
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return cm.CreditCard{}, http.StatusUnauthorized, errs
	}

	creditcardIdString := ctx.Param(constants.ID)

	creditcardId, err := uuid.Parse(creditcardIdString)

	if err != nil {
		return cm.CreditCard{}, http.StatusBadRequest, errs
	}

	creditcard, errs := c.repository.ReadCreditCardByID(ctx, creditcardId)

	if len(errs) > 0 {
		isNotFoundErr := u.FindIf(errs, func(err e.ApiError) bool {
			return err.String() == constants.StoreErrorNoRowsMsg
		})
		if isNotFoundErr {
			return cm.CreditCard{}, http.StatusNotFound, creditcardNotFoundErr
		}
		return cm.CreditCard{}, http.StatusInternalServerError, errs
	}

	if creditcard.UserID != userID {
		return cm.CreditCard{}, http.StatusForbidden, creditcardNotFoundErr
	}

	return creditcard, http.StatusOK, nil
}
