package cmd

import (
	healthHandler "project-two/internal/handler/health"
	healthRepo "project-two/internal/repository/health"
	healthSvc "project-two/internal/service/health"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	r := gin.Default()

	healthRepo := healthRepo.NewRepository()
	healthService := healthSvc.NewService(healthRepo)
	healthHandler := healthHandler.NewHandler(r, healthService)
	healthHandler.RouteList()

	r.Run(":8081")
}
