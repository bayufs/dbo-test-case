package models

import (
	"time"
)

// Order represents the orders entity
type Order struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderDate   time.Time  `gorm:"type:timestamptz;not null" json:"order_date"`
	CustomerID  uint       `json:"customer_id"`
	TotalAmount float64    `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	Status      string     `gorm:"type:varchar(50)" json:"status"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:timestamptz" json:"deleted_at,omitempty"`
}

// OrderItem represents the order_items entity
type OrderItem struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   uint       `json:"order_id"`
	ProductID uint       `json:"product_id"`
	Quantity  int        `gorm:"not null" json:"quantity"`
	UnitPrice float64    `gorm:"type:decimal(10,2);not null" json:"unit_price"`
	CreatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamptz" json:"deleted_at,omitempty"`
}

// Product represents the products entity
type Product struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductName string     `gorm:"type:varchar(255);not null" json:"product_name"`
	Description string     `json:"description"`
	Price       float64    `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int        `gorm:"not null" json:"stock"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:timestamptz" json:"deleted_at,omitempty"`
}
