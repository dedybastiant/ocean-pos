package middleware

import (
	"net/http"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			_, err := service.VerifyToken(authToken, "$2a$12$MNLqNZYZTTnS2/dvFrmmL..W6vrKSCxNS8BQaAv/jGPg6MJUCGIDm")
			if err != nil {
				c.JSON(http.StatusUnauthorized, dto.CommonResponse{Code: http.StatusUnauthorized, Status: "UNAUTHORIZED", Description: err.Error()})
				c.Abort()
				return
			}
			c.Set("x-user-id", 1)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Code: http.StatusUnauthorized, Status: "UNAUTHORIZED"})
		c.Abort()
	}
}
