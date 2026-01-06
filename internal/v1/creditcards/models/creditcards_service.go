package models

import (
	"financialcontrol/internal/models/errors"
	"net/http"
)

type CreditCardsService interface {
	Create(w http.ResponseWriter, r *http.Request) (CreditCardResponse, int, []errors.ApiError)
	Read(w http.ResponseWriter, r *http.Request) ([]CreditCardResponse, int, []errors.ApiError)
	ReadAt(w http.ResponseWriter, r *http.Request) (CreditCardResponse, int, []errors.ApiError)
	Update(w http.ResponseWriter, r *http.Request) (CreditCardResponse, int, []errors.ApiError)
	Delete(w http.ResponseWriter, r *http.Request) (int, []errors.ApiError)
}
