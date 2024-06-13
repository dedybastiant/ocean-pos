package controller

import (
	"net/http"
	"ocean-pos/internal/dto"
	"ocean-pos/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
}

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		UserService: service,
	}
}

func (controller *UserControllerImpl) Register(c *gin.Context) {
	var request dto.UserRequest
	c.BindJSON(&request)

	response, err := controller.UserService.Register(c, request)
	if err != nil {
		switch err.Error() {
		case "EMAIL_ALREADY_USED":
			c.IndentedJSON(http.StatusConflict,
				dto.CommonResponse{
					Code:        http.StatusConflict,
					Status:      "ERROR_DUPLICATE",
					Description: "email already used",
				})
		case "PHONE_NUMBER_ALREADY_USED":
			c.IndentedJSON(http.StatusConflict,
				dto.CommonResponse{
					Code:        http.StatusConflict,
					Status:      "ERROR_DUPLICATE",
					Description: "phone number already used",
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
		c.IndentedJSON(http.StatusCreated,
			dto.CommonResponse{
				Code:   http.StatusCreated,
				Status: "SUCCESS",
				Data:   response,
			})
	}
}
