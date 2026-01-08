package controllers

import (
	m "financialcontrol/internal/models"
	"financialcontrol/internal/utils"
	cm "financialcontrol/internal/v1/creditcards/models"
	"net/http"
)

type CreditCardsController struct {
	service cm.CreditCardsService
}

func NewCreditCardsController(service cm.CreditCardsService) *CreditCardsController {
	return &CreditCardsController{service: service}
}

func (c *CreditCardsController) Create(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.Create(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CreditCardsController) Read(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.Read(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CreditCardsController) ReadAt(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.ReadAt(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CreditCardsController) Update(w http.ResponseWriter, r *http.Request) {
	data, status, err := c.service.Update(w, r)
	utils.SendResponse(w, data, status, err)
}

func (c *CreditCardsController) Delete(w http.ResponseWriter, r *http.Request) {
	status, err := c.service.Delete(w, r)
	utils.SendResponse(w, m.NewResponseSuccess(), status, err)
}
