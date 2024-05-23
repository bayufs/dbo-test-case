package resources

import (
	"time"
)

// OrderDTO is the data transfer object for Order
type OrderDTO struct {
	ID          uint      `json:"id"`
	OrderDate   time.Time `json:"order_date"`
	CustomerID  uint      `json:"customer_id"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// OrderItemDTO is the data transfer object for OrderItem
type OrderItemDTO struct {
	ID        uint      `json:"id"`
	OrderID   uint      `json:"order_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unit_price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ProductDTO is the data transfer object for Product
type ProductDTO struct {
	ID          uint      `json:"id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
