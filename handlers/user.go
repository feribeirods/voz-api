package handlers

import (
	"net/http"

	"github.com/feribeirods/voz-api/config"
	"github.com/feribeirods/voz-api/models"
	"github.com/gin-gonic/gin"
)

func MeHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Usuário não encontrado no contexto"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}
