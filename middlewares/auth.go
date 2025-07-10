package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lê o token do header Authorization
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não enviado"})
			return
		}

		// Extrai o token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Valida o token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica se o método de assinatura é o esperado (HS256)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("JWT error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			return
		}

		// Pega o ID do usuário (sub) e salva no contexto
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["sub"] == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			return
		}

		// Armazena o user ID no contexto (como uint)
		c.Set("userID", uint(claims["sub"].(float64)))

		// Continua a requisição
		c.Next()
	}
}
