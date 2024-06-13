package main

import (
	"ocean-pos/config"
	"ocean-pos/internal/controller"
	"ocean-pos/internal/middleware"
	"ocean-pos/internal/repository"
	"ocean-pos/internal/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.NewDB()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	authService := service.NewAuthService(userRepository, db)
	authController := controller.NewAuthController(authService)

	authMiddleware := middleware.AuthMiddleware()

	r := gin.Default()
	r.POST("/users", authMiddleware, userController.Register)
	r.POST("/auth", authController.Login)

	r.Run()
}
