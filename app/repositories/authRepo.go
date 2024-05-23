package repositories

import (
	"dbo-test-case/app/models"
	"dbo-test-case/config"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authDbConnection struct {
	connection *gorm.DB
}

type AuthRepo interface {
	Authorize(username, password string) (models.Authentication, error)
}

func NewAuthRepo() AuthRepo {
	return &authDbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *authDbConnection) Authorize(username, password string) (models.Authentication, error) {

	var auth models.Authentication

	err := db.connection.Where("username = ?", username).First(&auth)

	fmt.Println(err, "Check if customer is exist")

	if errors.Is(err.Error, gorm.ErrRecordNotFound) {

		fmt.Println(err, "User does not exist")

		return auth, errors.New("Customer does not exist")
	}

	err2 := bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))

	if err2 != nil {
		fmt.Println("Hash password result compare : ", err2)
		return auth, errors.New("Password does not match")
	}

	return auth, nil

}
