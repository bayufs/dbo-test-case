package repositories

import (
	"dbo-test-case/app/models"
	"dbo-test-case/app/resources"
	"dbo-test-case/config"
	"errors"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type customerDbConnection struct {
	connection *gorm.DB
}

type CustomerRepository interface {
	GetCustomerList(queryStringParam map[string]interface{}) (map[string]interface{}, []models.Customer, error)
	GetCustomerByID(customerID uint) (*models.Customer, error)
	StoreNewCustomer(payload resources.StoreNewCustomer) (uint, error)
	UpdateCustomer(customerID uint, payload resources.StoreNewCustomer) error
}

func NewCustomerRepository() CustomerRepository {
	return &customerDbConnection{
		connection: config.ConnectDB(),
	}
}

func (db *customerDbConnection) GetCustomerList(queryStringParam map[string]interface{}) (map[string]interface{}, []models.Customer, error) {

	var customer []models.Customer

	query := db.connection

	query = query.Debug().Model(&models.Customer{})

	if sortParam, ok := queryStringParam["sort"]; ok {
		query = query.Order("updated_at " + sortParam.(string))
	}

	if searchParam, ok := queryStringParam["search"]; ok {
		search := searchParam.(string)
		if search != "" {
			query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")

		}
	}

	var (
		page  int
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

	if err := query.Find(&customer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, errors.New("No records found")
		}
		return nil, nil, err
	}

	meta := map[string]interface{}{
		"count":  len(customer),
		"offset": offset,
		"limit":  limit,
		"total":  total,
	}

	return meta, customer, nil

}

func (db *customerDbConnection) GetCustomerByID(customerID uint) (*models.Customer, error) {
	var customer models.Customer

	if err := db.connection.First(&customer, customerID).Error; err != nil {
		return nil, fmt.Errorf("error getting customer by ID: %v", err)
	}

	return &customer, nil
}

func (db *customerDbConnection) StoreNewCustomer(payload resources.StoreNewCustomer) (uint, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("error hashing password: %v", err)
	}

	customer := models.Customer{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Phone:     payload.Phone,
		Address:   payload.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.connection.Create(&customer).Error; err != nil {
		return 0, fmt.Errorf("error creating customer: %v", err)
	}

	auth := models.Authentication{
		CustomerID: customer.ID,
		Username:   payload.Username,
		Password:   string(hashedPassword),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := db.connection.Create(&auth).Error; err != nil {
		return 0, fmt.Errorf("error creating authentication record: %v", err)
	}

	return customer.ID, nil

}

func (db *customerDbConnection) UpdateCustomer(customerID uint, payload resources.StoreNewCustomer) error {
	var existingCustomer models.Customer
	if err := db.connection.First(&existingCustomer, customerID).Error; err != nil {
		return fmt.Errorf("error finding customer: %v", err)
	}

	existingCustomer.FirstName = payload.FirstName
	existingCustomer.LastName = payload.LastName
	existingCustomer.Email = payload.Email
	existingCustomer.Phone = payload.Phone
	existingCustomer.Address = payload.Address
	existingCustomer.UpdatedAt = time.Now()

	if err := db.connection.Save(&existingCustomer).Error; err != nil {
		return fmt.Errorf("error updating customer: %v", err)
	}

	if payload.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing password: %v", err)
		}

		var existingAuth models.Authentication
		if err := db.connection.Where("customer_id = ?", customerID).First(&existingAuth).Error; err != nil {
			return fmt.Errorf("error finding authentication record: %v", err)
		}
		existingAuth.Username = payload.Username
		existingAuth.Password = string(hashedPassword)
		existingAuth.UpdatedAt = time.Now()

		if err := db.connection.Save(&existingAuth).Error; err != nil {
			return fmt.Errorf("error updating authentication record: %v", err)
		}
	}

	return nil
}
