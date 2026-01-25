package categories

import (
	"financialcontrol/internal/models"
	"financialcontrol/internal/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(ctx *gin.Context) {
	data, status, err := h.service.Create(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Handler) Read(ctx *gin.Context) {
	data, status, err := h.service.Read(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Handler) ReadByID(ctx *gin.Context) {
	data, status, err := h.service.ReadByID(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Handler) Update(ctx *gin.Context) {
	data, status, err := h.service.Update(ctx)
	utils.SendResponse(ctx, data, status, err)
}

func (h *Handler) Delete(ctx *gin.Context) {
	status, err := h.service.Delete(ctx)
	utils.SendResponse(ctx, models.NewResponseSuccess(), status, err)
}
