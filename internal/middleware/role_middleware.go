package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(
	allowedRoles ...string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		role, exists :=
			c.Get("role")

		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Access denied",
			})
			c.Abort()
			return
		}

		userRole := role.(string)

		for _, allowed := range allowedRoles {

			if userRole == allowed {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Forbidden access",
		})

		c.Abort()
	}
}
