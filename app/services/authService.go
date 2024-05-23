package services

import (
	"dbo-test-case/app/models"
	"dbo-test-case/app/repositories"
	"errors"
	"fmt"
)

type AuthService struct {
	authRepo repositories.AuthRepo
}

func NewAuthService() *AuthService {
	return &AuthService{
		authRepo: repositories.NewAuthRepo(),
	}
}

func (h *AuthService) Authorize(username, password string) (*models.Authentication, error) {

	result, err := h.authRepo.Authorize(username, password)

	if err != nil {
		fmt.Println("Something went wrong while trying login, see logs below.")
		fmt.Println(err.Error())
		return nil, errors.New(err.Error())
	}

	return &result, nil

}
