package main

import (
	"log"
	"ocean-pos/config"
	"ocean-pos/internal/controller"
	"ocean-pos/internal/middleware"
	"ocean-pos/internal/repository"
	"ocean-pos/internal/service"

	"github.com/go-playground/validator/v10"
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
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db)
	userController := controller.NewUserController(userService)

	businessRepository := repository.NewBusinessRepository()
	businessService := service.NewBusinessService(businessRepository, db, validate)
	businessController := controller.NewBusinessController(businessService)

	storeRepository := repository.NewStoreRepository()
	storeService := service.NewStoreService(storeRepository, businessRepository, db, validate)
	storeController := controller.NewStoreController(storeService)

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, businessRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	menuRepository := repository.NewMenuRepository()
	menuService := service.NewMenuService(menuRepository, businessRepository, categoryRepository, db, validate)
	menuController := controller.NewMenuController(menuService)

	storeMenuRepository := repository.NewStoreMenuRepository()
	storeMenuService := service.NewStoreMenuService(storeMenuRepository, businessRepository, storeRepository, categoryRepository, menuRepository, db, validate)
	storeMenuController := controller.NewStoreMenuController(storeMenuService)

	authService := service.NewAuthService(userRepository, db, rdb, viperConfig)
	authController := controller.NewAuthController(authService)

	authMiddleware := middleware.AuthMiddleware(rdb, viperConfig)

	r := gin.Default()
	r.POST("/auth/login", authController.Login)
	r.POST("/auth/logout", authMiddleware, authController.Logout)

	r.POST("/users", userController.Register)
	r.GET("/users/:userId", authMiddleware, userController.FindUserById)

	r.POST("/businesses", authMiddleware, businessController.RegisterBusiness)
	r.GET("/businesses/:businessId", authMiddleware, businessController.GetBusinessById)

	r.POST("/stores", authMiddleware, storeController.RegisterStore)

	r.POST("/categories", authMiddleware, categoryController.CreateCategory)

	r.POST("/menus", authMiddleware, menuController.AddNewMenu)

	r.POST("/store-menus", authMiddleware, storeMenuController.AddNewStoreMenu)

	r.Run()
}
