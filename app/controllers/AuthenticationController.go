package controllers

import (
	"dbo-test-case/app/helpers"
	"dbo-test-case/app/resources"
	"dbo-test-case/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{
		authService: *services.NewAuthService(),
	}
}

func (h *AuthController) Login(c *gin.Context) {
	var payload resources.Login

	var errMsg string

	if err := c.ShouldBindJSON(&payload); err != nil {

		switch err.(type) {
		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "Username" && fieldError.Tag() == "required" {
					errMsg = "Username is required"
				} else if fieldError.Field() == "Password" && fieldError.Tag() == "required" {
					errMsg = "Password required"
				} else {
					errMsg = "Invalid request"
				}
			}

		default:
			errMsg = "Invalid request"

		}

		c.AbortWithStatusJSON(http.StatusOK,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Credentials error", errMsg),
		)

		return

	}

	res, err := h.authService.Authorize(payload.Username, payload.Password)

	if err != nil {

		c.AbortWithStatusJSON(http.StatusOK,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Credentials error", "Invalid username or password"),
		)

		return
	}

	access_token, err := helpers.GenerateToken(res)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Credentials error", "Invalid username or password"),
		)

		return
	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success login.", access_token),
	)
	return

}
