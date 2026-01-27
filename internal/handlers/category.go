package handlers

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/services"
	"financialcontrol/internal/utils"

	"github.com/gin-gonic/gin"
)

type Category struct {
	service services.Category
}

func NewCategoriesHandler(service services.Category) *Category {
	return &Category{service: service}
}

func (h *Category) Create(ctx *gin.Context) {
	data, status, err := h.service.Create(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Category) Read(ctx *gin.Context) {
	data, status, err := h.service.Read(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Category) ReadByID(ctx *gin.Context) {
	data, status, err := h.service.ReadByID(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Category) Update(ctx *gin.Context) {
	data, status, err := h.service.Update(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Category) Delete(ctx *gin.Context) {
	status, err := h.service.Delete(ctx)
	utils.SendResponse(ctx, models.NewResponseSuccess(), status, err)
}
