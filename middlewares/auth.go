package middlewares

import (
	"gonexwind/backend-api/config" // Mengambil konfigurasi dari file .env
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Ambil secret key dari environment variable
// Jika tidak ada, gunakan default "secret_key"
var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

type MyClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &MyClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Simpan user_id ke gin context supaya bisa diakses controller
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
