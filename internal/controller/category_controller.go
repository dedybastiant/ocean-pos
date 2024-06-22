package controller

import (
	"net/http"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	CreateCategory(c *gin.Context)
}

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (controller *CategoryControllerImpl) CreateCategory(c *gin.Context) {
	var request dto.CreateCategoryRequest
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

	response, err := controller.CategoryService.CreateCategory(c, request, userId.(int))
	if err != nil {
		if strings.HasPrefix(err.Error(), "validation error") {
			c.JSON(http.StatusBadRequest, dto.CommonResponse{
				Code:        http.StatusBadRequest,
				Status:      "BAD_REQUEST",
				Description: err.Error(),
			})
			return
		}
		if err.Error() == "FORBIDDEN" {
			c.JSON(http.StatusForbidden, dto.CommonResponse{
				Code:        http.StatusForbidden,
				Status:      "FORBIDDEN",
				Description: "you have no permission",
			})
			return
		}
		if err.Error() == "DUPLICATE_CATEGORY" {
			c.IndentedJSON(http.StatusConflict,
				dto.CommonResponse{
					Code:        http.StatusConflict,
					Status:      "ERROR_DUPLICATE",
					Description: "caterogy already exsist",
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
