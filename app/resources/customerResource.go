package resources

import (
	"time"
)

// CustomerDTO is the data transfer object for Customer
type CustomerDTO struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthenticationDTO is the data transfer object for Authentication
type AuthenticationDTO struct {
	ID         uint      `json:"id"`
	CustomerID uint      `json:"customer_id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type StoreNewCustomer struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
