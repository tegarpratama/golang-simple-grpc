package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CheckHealth(c *gin.Context) {
	res := h.healthSvc.CheckHealth()
	c.JSON(http.StatusOK, gin.H{
		"message": res,
	})
}
