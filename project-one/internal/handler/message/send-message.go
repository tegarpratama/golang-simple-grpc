package message

import (
	"net/http"

	"project-one/internal/model/message"

	"github.com/gin-gonic/gin"
)

func (h *handler) SendMessage(c *gin.Context) {
	var input message.MessageReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.messageSvc.SendMessage(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
