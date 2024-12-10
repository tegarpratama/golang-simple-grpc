package message

import (
	"context"
	"project-one/internal/model/message"

	"github.com/gin-gonic/gin"
)

type messageService interface {
	SendMessage(ctx context.Context, req message.MessageReq) (*message.MessageRes, error)
}

type handler struct {
	api        *gin.Engine
	messageSvc messageService
}

func NewHandler(api *gin.Engine, messageSvc messageService) *handler {
	return &handler{
		api:        api,
		messageSvc: messageSvc,
	}
}

func (h *handler) RouteList() {
	route := h.api.Group("/message")
	route.POST("/", h.SendMessage)
}
