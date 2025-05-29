package midlleware

import (
	"TurAgency/src/utils"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "Couldn't get cookie 'token'"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "Couldn't parse the token: " + err.Error()})
			c.Abort()
			return
		}
		c.Set("role", claims.Role)
		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(403, gin.H{"error": "Access denied, only admin can do this"})
			c.Abort()
			return
		}
		c.Next()
	}
}
