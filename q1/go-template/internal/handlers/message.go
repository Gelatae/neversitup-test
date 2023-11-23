package handlers

import (
	ports "go-template/internal/core/ports/services"
	"go-template/internal/handlers/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type messageHandler struct {
	messageService ports.MessageService
}

func NewMessageHandler(ms ports.MessageService) *messageHandler {
	return &messageHandler{messageService: ms}
}

func (m *messageHandler) GetMessage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"message": err.Error()})
		return
	}
	message, err := m.messageService.GetMessage(c, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, model.NewMessage(message))
}
