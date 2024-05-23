package services

import (
	"dbo-test-case/app/models"
	"dbo-test-case/app/repositories"
	"dbo-test-case/app/resources"
	"errors"
	"fmt"
)

type CustomerService struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		customerRepo: repositories.NewCustomerRepository(),
	}
}

func (h *CustomerService) GetCustomerList(queryStringParam map[string]interface{}) (map[string]interface{}, []models.Customer, error) {

	meta, resultGetCustomer, errGetCustomerList := h.customerRepo.GetCustomerList(queryStringParam)

	if errGetCustomerList != nil {
		fmt.Println("Something went wrong while trying get customer list, see logs below.")
		fmt.Println(errGetCustomerList.Error())
		return nil, nil, errors.New(errGetCustomerList.Error())
	}

	return meta, resultGetCustomer, nil

}

func (h *CustomerService) GetCustomer(customerID uint) (*models.Customer, error) {

	resultGetCustomer, errGetCustome := h.customerRepo.GetCustomerByID(customerID)

	if errGetCustome != nil {
		fmt.Println("Something went wrong while trying get customer, see logs below.")
		fmt.Println(errGetCustome.Error())
		return nil, errors.New(errGetCustome.Error())
	}

	return resultGetCustomer, nil

}

func (h *CustomerService) StoreNewCustomer(payload resources.StoreNewCustomer) error {

	_, errStoreNewCustomer := h.customerRepo.StoreNewCustomer(payload)

	if errStoreNewCustomer != nil {
		fmt.Println("Something went wrong while trying store new customer, see logs below.")
		fmt.Println(errStoreNewCustomer.Error())
		return errors.New(errStoreNewCustomer.Error())
	}

	return nil

}

func (h *CustomerService) UpdateCustomer(customerID uint, payload resources.StoreNewCustomer) error {

	errUpdateCustomer := h.customerRepo.UpdateCustomer(customerID, payload)

	if errUpdateCustomer != nil {
		fmt.Println("Something went wrong while trying update customer data, see logs below.")
		fmt.Println(errUpdateCustomer.Error())
		return errors.New(errUpdateCustomer.Error())
	}

	return nil

}

func (h *CustomerService) DeleteCustomer(customerID uint) error {

	errDeleteOrder := h.customerRepo.DeleteCustomer(customerID)

	if errDeleteOrder != nil {
		fmt.Println("Something went wrong while trying delete customer, see logs below.")
		fmt.Println(errDeleteOrder.Error())
		return errors.New(errDeleteOrder.Error())
	}

	return nil

}
