package http_handlers

import (
	"chat-backend/internal/domain"
	http_types "chat-backend/internal/http/types"
	"chat-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// AddChat @Summary Add chat
// @Description Add chat
// @ID add-chat
// @Accept  json
// @Produce  json
// @Param   body body http_types.CreateChat true "ping request"
// @Success 201 {object} http_types.ReadChat
// @Router /api/v1/chat [post]
func AddChat(c *gin.Context) {
	var createChat http_types.CreateChat

	if err := c.ShouldBindJSON(&createChat); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	chat, err := service.AddChat(createChat.Title, createChat.UserId)
	if err != nil {
		// TODO: add log
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, http_types.NewReadChatFromDomain(chat))
}

func getChat(c *gin.Context) {
	c.JSON(200, domain.Chat{})
}
