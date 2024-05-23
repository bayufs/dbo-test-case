package models

import (
	"time"
)

// Customer represents the customers entity
type Customer struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string     `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string     `gorm:"type:varchar(255);not null" json:"last_name"`
	Email     string     `gorm:"type:varchar(255);not null" json:"email"`
	Phone     string     `gorm:"type:varchar(50)" json:"phone"`
	Address   string     `gorm:"type:varchar(255)" json:"address"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamptz" json:"deleted_at,omitempty"`
}

// Authentication represents the authentications entity
type Authentication struct {
	ID         uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID uint       `json:"customer_id"`
	Username   string     `gorm:"type:varchar(255);not null" json:"username"`
	Password   string     `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt  time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time `gorm:"type:timestamptz" json:"deleted_at,omitempty"`
}
