package midlleware

import (
	"TurAgency/internal/database"
	"TurAgency/internal/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// проверка в Redis
		_, err = database.RedisDB.Get(context.Background(), cookie).Result()
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		// кладём в контекст
		c.Set("role", claims.Role)
		c.Set("userID", claims.UserID)

		c.Next()
	}
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "Администратор" {

			// Перенаправляем на главную страницу или другую страницу
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}
