package content_item_hdlr

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-content-manager/pkg/ports"
)

func AuthMiddleware(userService ports.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <api_key>"})
			c.Abort()
			return
		}

		// Extract the API key from the header
		apiKey := strings.TrimPrefix(authHeader, "Bearer ")

		user, err := userService.Auth(apiKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Ooops! Something went wrong, please try again later"})
			c.Abort()
		}
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
		}

		c.Next()
	}
}
