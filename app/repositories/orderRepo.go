package repositories

import (
	"dbo-test-case/app/models"
	"dbo-test-case/app/resources"
	"dbo-test-case/config"
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

const (
	PENDING = "pending"
	SUCCESS = "success"
	FAILED  = "failed"
)

type orderDbConnection struct {
	connection *gorm.DB
}

type OrderRepository interface {
	CreateOrder(payload resources.CreateOrder) (uint, error)
	UpdateOrder(orderID uint, payload resources.UpdateOrder) error
	GetOrderList(queryStringParam map[string]interface{}) (map[string]interface{}, []models.Order, error)
	GetOrderByID(orderID uint) (*models.Order, error)
	DeleteOrder(orderID uint) error
}

func NewOrderRepository() OrderRepository {
	return &orderDbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *orderDbConnection) GetOrderList(queryStringParam map[string]interface{}) (map[string]interface{}, []models.Order, error) {

	var orders []models.Order

	query := db.connection.Model(&models.Order{}).Debug().
		Preload("Customer").
		Preload("OrderItem").
		Preload("OrderItem.Product")

	if sortParam, ok := queryStringParam["sort"]; ok {
		query = query.Order("updated_at " + sortParam.(string))
	}

	if searchParam, ok := queryStringParam["search"]; ok {
		search := searchParam.(string)
		if search != "" {
			query = query.Joins("JOIN order_items ON order_items.order_id = orders.id").
				Joins("JOIN products ON products.id = order_items.product_id").
				Where("orders.status ILIKE ? OR orders.id::text ILIKE ? OR products.product_name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
		}
	}

	var (
		page  int = 1
		limit int = 10
	)

	if pageParam, ok := queryStringParam["page"]; ok {
		pageStr, _ := pageParam.(string)
		page, _ = strconv.Atoi(pageStr)
	}

	if limitParam, ok := queryStringParam["limit"]; ok {
		limitStr, _ := limitParam.(string)
		limit, _ = strconv.Atoi(limitStr)
	}

	offset := (page - 1) * limit

	var total int64
	query.Count(&total)

	query = query.Limit(limit).Offset(offset)

	if err := query.Find(&orders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, fmt.Errorf("no records found")
		}
		return nil, nil, err
	}

	meta := map[string]interface{}{
		"count":  len(orders),
		"offset": offset,
		"limit":  limit,
		"total":  total,
	}

	return meta, orders, nil
}

func (db *orderDbConnection) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order

	if err := db.connection.
		Preload("Customer").
		Preload("OrderItem").
		Preload("OrderItem.Product").
		First(&order, orderID).Error; err != nil {
		return nil, fmt.Errorf("error getting order by ID: %v", err)
	}

	return &order, nil
}

func (db *orderDbConnection) CreateOrder(payload resources.CreateOrder) (uint, error) {

	tx := db.connection.Begin()
	if tx.Error != nil {
		return 0, fmt.Errorf("error starting transaction: %v", tx.Error)
	}

	order := models.Order{
		OrderDate:   time.Now(),
		CustomerID:  payload.CustomerID,
		TotalAmount: payload.TotalAmount,
		Status:      PENDING,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error creating order: %v", err)
	}

	orderItem := models.OrderItem{
		OrderID:   order.ID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
		UnitPrice: payload.UnitPrice,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := tx.Create(&orderItem).Error; err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("error creating order item: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return 0, fmt.Errorf("error committing transaction: %v", err)
	}

	return order.ID, nil
}

func (db *orderDbConnection) UpdateOrder(orderID uint, payload resources.UpdateOrder) error {

	tx := db.connection.Begin()

	if tx.Error != nil {
		return fmt.Errorf("error starting transaction: %v", tx.Error)
	}

	var order models.Order

	if err := tx.First(&order, orderID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error finding order: %v", err)
	}

	order.TotalAmount = payload.TotalAmount
	order.Status = payload.Status
	order.UpdatedAt = time.Now()

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating order: %v", err)
	}

	var orderItem models.OrderItem
	if err := tx.Where("order_id = ?", orderID).First(&orderItem).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error finding order item: %v", err)
	}

	orderItem.ProductID = payload.ProductID
	orderItem.Quantity = payload.Quantity
	orderItem.UnitPrice = payload.UnitPrice
	orderItem.UpdatedAt = time.Now()

	if err := tx.Save(&orderItem).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating order item: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (db *orderDbConnection) DeleteOrder(orderID uint) error {
	tx := db.connection.Begin()
	if tx.Error != nil {
		return fmt.Errorf("error starting transaction: %v", tx.Error)
	}

	var order models.Order
	if err := tx.Preload("OrderItem").First(&order, orderID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order not found")
		}
		return fmt.Errorf("error finding order: %v", err)
	}

	if err := tx.Where("order_id = ?", orderID).Delete(&models.OrderItem{}).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting order items: %v", err)
	}

	if err := tx.Delete(&order).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting order: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}
