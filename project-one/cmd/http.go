package cmd

import (
	healthHandler "project-one/internal/handler/health"
	healthRepo "project-one/internal/repository/health"
	healthSvc "project-one/internal/service/health"

	messageHandler "project-one/internal/handler/message"
	messageRepo "project-one/internal/repository/message"
	messageSvc "project-one/internal/service/message"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	r := gin.Default()

	healthRepo := healthRepo.NewRepository()
	healthService := healthSvc.NewService(healthRepo)
	healthHandler := healthHandler.NewHandler(r, healthService)
	healthHandler.RouteList()

	messageRepo := messageRepo.NewRepository()
	messageService := messageSvc.NewService(messageRepo)
	messageHandler := messageHandler.NewHandler(r, messageService)
	messageHandler.RouteList()

	r.Run(":8080")
}
