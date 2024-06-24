package controller

import (
	"net/http"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type MenuController interface {
	AddNewMenu(c *gin.Context)
}

type MenuControllerImpl struct {
	MenuService service.MenuService
}

func NewMenuController(menuService service.MenuService) MenuController {
	return &MenuControllerImpl{
		MenuService: menuService,
	}
}

func (controller *MenuControllerImpl) AddNewMenu(c *gin.Context) {
	var request dto.AddMenuRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest,
			dto.CommonResponse{
				Code:        http.StatusBadRequest,
				Status:      "BAD_REQUEST",
				Description: "Invalid request payload",
			})
		return
	}

	userId, ok := c.Get("x-user-id")
	if !ok {
		c.IndentedJSON(http.StatusUnauthorized,
			dto.CommonResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			})
	}

	response, err := controller.MenuService.AddNewMenu(c, request, userId.(int))
	if err != nil {
		if strings.HasPrefix(err.Error(), "validation error") {
			c.JSON(http.StatusBadRequest, dto.CommonResponse{
				Code:        http.StatusBadRequest,
				Status:      "BAD_REQUEST",
				Description: err.Error(),
			})
			return
		}
		if err.Error() == "BUSINESS_NOT_FOUND" || err.Error() == "CATEGORY_NOT_FOUND" {
			c.JSON(http.StatusForbidden, dto.CommonResponse{
				Code:        http.StatusForbidden,
				Status:      "FORBIDDEN",
				Description: "have no permission to access",
			})
			return
		}
		if err.Error() == "DUPLICATE_MENU" {
			c.IndentedJSON(http.StatusConflict,
				dto.CommonResponse{
					Code:        http.StatusConflict,
					Status:      "ERROR_DUPLICATE",
					Description: "caterogy already exist",
				})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CommonResponse{
		Code:   http.StatusCreated,
		Status: "SUCCESS",
		Data:   response,
	})
}
