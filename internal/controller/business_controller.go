package controller

import (
	"net/http"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type BusinessController interface {
	RegisterBusiness(c *gin.Context)
}

type BusinessControllerImpl struct {
	BusinessService service.BusinessService
}

func NewBusinessController(service service.BusinessService) BusinessController {
	return &BusinessControllerImpl{
		BusinessService: service,
	}
}

func (controller *BusinessControllerImpl) RegisterBusiness(c *gin.Context) {
	var request dto.BusinessRequest
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

	response, err := controller.BusinessService.RegisterBusiness(c, request, userId.(int))
	if err != nil {
		if strings.HasPrefix(err.Error(), "validation error") {
			c.JSON(http.StatusBadRequest, dto.CommonResponse{
				Code:        http.StatusBadRequest,
				Status:      "BAD_REQUEST",
				Description: err.Error(),
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
