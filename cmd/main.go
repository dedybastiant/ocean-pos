package main

import (
	"log"
	"ocean-pos/config"
	"ocean-pos/internal/controller"
	"ocean-pos/internal/middleware"
	"ocean-pos/internal/repository"
	"ocean-pos/internal/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	viperConfig, err := config.NewViper()
	if err != nil {
		log.Fatalf("Error initializing config: %v", err)
	}

	db := config.NewDB(viperConfig)
	rdb := config.NewRdb()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	authService := service.NewAuthService(userRepository, db, rdb, viperConfig)
	authController := controller.NewAuthController(authService)

	authMiddleware := middleware.AuthMiddleware(rdb, viperConfig)

	r := gin.Default()
	r.POST("/auth/login", authController.Login)
	r.POST("/auth/logout", authMiddleware, authController.Logout)

	r.POST("/users", authMiddleware, userController.Register)

	r.Run()
}
