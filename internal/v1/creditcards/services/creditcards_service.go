package services

import (
	e "financialcontrol/internal/models/errors"
	s "financialcontrol/internal/store"
	u "financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreditCardsService struct {
	repository cm.CreditCardsRepository
}

func NewCreditCardsService(repository cm.CreditCardsRepository) cm.CreditCardsService {
	return &CreditCardsService{repository: repository}
}

func (c CreditCardsService) read(ctx *gin.Context) (cm.CreditCard, int, []e.ApiError) {
	creditcardNotFoundErr := []e.ApiError{e.NotFoundError{Message: e.CreditcardNotFound}}
	userID, errs := u.ReadUserIdFromCookie(ctx)

	if len(errs) > 0 {
		return cm.CreditCard{}, http.StatusUnauthorized, errs
	}

	creditcardIdString := ctx.Param("id")

	creditcardId, err := uuid.Parse(creditcardIdString)

	if err != nil {
		return cm.CreditCard{}, http.StatusBadRequest, errs
	}

	creditcard, errs := c.repository.ReadByID(ctx, creditcardId)

	if len(errs) > 0 {
		isNotFoundErr := u.FindIf(errs, func(err e.ApiError) bool {
			return err.String() == string(s.ErrNoRows)
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
