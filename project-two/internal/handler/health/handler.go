package health

import "github.com/gin-gonic/gin"

type healthService interface {
	CheckHealth() string
}

type handler struct {
	api       *gin.Engine
	healthSvc healthService
}

func NewHandler(api *gin.Engine, healthSvc healthService) *handler {
	return &handler{
		api:       api,
		healthSvc: healthSvc,
	}
}

func (h *handler) RouteList() {
	route := h.api.Group("/health")
	route.GET("/", h.CheckHealth)
}
