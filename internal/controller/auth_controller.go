package controller

import (
	"net/http"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: service,
	}
}

func (controller *AuthControllerImpl) Login(c *gin.Context) {
	var request dto.AuthRequest
	c.BindJSON(&request)

	response, err := controller.AuthService.Login(c, request)
	if err != nil {
		switch err.Error() {
		case "EMAIL_NOT_FOUND":
			c.IndentedJSON(http.StatusNotFound,
				dto.CommonResponse{
					Code:        http.StatusNotFound,
					Status:      "NOT_FOUND",
					Description: "email not registered",
				})
		case "INCORRECT_CREDENTIAL":
			c.IndentedJSON(http.StatusBadRequest,
				dto.CommonResponse{
					Code:        http.StatusBadRequest,
					Status:      "BAD_REQUEST",
					Description: "incorrect password",
				})
		default:
			c.IndentedJSON(http.StatusInternalServerError,
				dto.CommonResponse{
					Code:        http.StatusInternalServerError,
					Status:      "INTERNAL_SERVER_ERROR",
					Description: "something went wrong",
				})
		}
	} else {
		c.IndentedJSON(http.StatusOK,
			dto.CommonResponse{
				Code:   http.StatusOK,
				Status: "SUCCESS",
				Data:   response,
			})
	}
}
