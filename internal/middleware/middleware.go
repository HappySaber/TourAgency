package midlleware

import (
	"TurAgency/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			// Безопасно прерываем обработку, не отправляя других данных
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(cookie)
		if err != nil {
			// Тоже просто редиректим, не отправляя JSON — иначе возникнет конфликт
			c.Redirect(http.StatusSeeOther, "/login")
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
		if !exists || role != "Администратор" {

			// Перенаправляем на главную страницу или другую страницу
			c.Redirect(http.StatusSeeOther, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}
