package handlers

import (
	"net/http"

	"github.com/feribeirods/voz-api/config"
	"github.com/feribeirods/voz-api/models"
	"github.com/feribeirods/voz-api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Validação simples
	if input.Name == "" || input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os campos são obrigatórios"})
		return
	}

	// Gerar hash seguro da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	// Criar usuário com senha já protegida
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar usuário"})
		return
	}

	// Retorno sem expor a senha
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário registrado com sucesso",
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

func LoginHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "E-mail não encontrado"})
		return
	}

	// Comparar senha com o hash salvo
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha incorreta"})
		return
	}

	// Gerar token JWT
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	// Retornar token
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
