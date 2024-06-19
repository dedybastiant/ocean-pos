package middleware

import (
	"context"
	"net/http"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func AuthMiddleware(rdb *redis.Client, viperConfig *viper.Viper) gin.HandlerFunc {
	secretKey := viperConfig.GetString("ACCESS_TOKEN_KEY")

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if authHeader == "" || len(t) < 2 {
			c.JSON(http.StatusUnauthorized, dto.CommonResponse{Code: http.StatusUnauthorized, Status: "UNAUTHORIZED"})
			c.Abort()
			return
		}

		exists, _ := rdb.Exists(context.Background(), t[1]).Result()

		if exists == 1 {
			c.JSON(http.StatusUnauthorized, dto.CommonResponse{Code: http.StatusUnauthorized, Status: "UNAUTHORIZED"})
			c.Abort()
			return
		}
		if len(t) == 2 {
			authToken := t[1]
			claims, err := service.VerifyToken(authToken, secretKey)
			if err != nil {
				c.JSON(http.StatusUnauthorized, dto.CommonResponse{Code: http.StatusUnauthorized, Status: "UNAUTHORIZED"})
				c.Abort()
				return
			}

			sub := (*claims)["sub"].(float64)
			userId := int(sub)

			c.Set("x-user-id", userId)
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Code: http.StatusUnauthorized, Status: "UNAUTHORIZED"})
		c.Abort()
	}
}
