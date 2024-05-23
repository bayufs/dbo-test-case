package resources

import (
	"time"
)

// OrderDTO is the data transfer object for Order
type OrderDTO struct {
	ID          uint         `json:"id"`
	OrderDate   time.Time    `json:"order_date"`
	CustomerID  uint         `json:"customer_id"`
	TotalAmount float64      `json:"total_amount"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	OrderItem   OrderItemDTO `json:"order_item"`
	Customer    CustomerDTO  `json:"customer"`
}

// OrderItemDTO is the data transfer object for OrderItem
type OrderItemDTO struct {
	ID        uint       `json:"id"`
	ProductID uint       `json:"product_id"`
	Quantity  int        `json:"quantity"`
	UnitPrice float64    `json:"unit_price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Product   ProductDTO `json:"product"`
}

type ProductDTO struct {
	ID          uint      `json:"id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateOrder struct {
	CustomerID  uint    `json:"customer_id" binding:"required"`
	TotalAmount float64 `json:"total_amount"  binding:"required"`
	ProductID   uint    `json:"product_id" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	UnitPrice   float64 `json:"unit_price" binding:"required"`
}

type UpdateOrder struct {
	CustomerID  uint    `json:"customer_id" binding:"required"`
	TotalAmount float64 `json:"total_amount"  binding:"required"`
	ProductID   uint    `json:"product_id" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	UnitPrice   float64 `json:"unit_price" binding:"required"`
	Status      string  `json:"status" binding:"required"`
}
