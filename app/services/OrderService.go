package services

import (
	"dbo-test-case/app/models"
	"dbo-test-case/app/repositories"
	"dbo-test-case/app/resources"
	"errors"
	"fmt"
)

type OrderService struct {
	orderRepo repositories.OrderRepository
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo: repositories.NewOrderRepository(),
	}
}

func (h *OrderService) GetOrderList(queryStringParam map[string]interface{}) (map[string]interface{}, []models.Order, error) {

	meta, resultGetOrder, errGetOrderList := h.orderRepo.GetOrderList(queryStringParam)

	if errGetOrderList != nil {
		fmt.Println("Something went wrong while trying get customer list, see logs below.")
		fmt.Println(errGetOrderList.Error())
		return nil, nil, errors.New(errGetOrderList.Error())
	}

	return meta, resultGetOrder, nil

}

func (h *OrderService) CreateOrder(payload resources.CreateOrder) error {

	_, errCreateOrder := h.orderRepo.CreateOrder(payload)

	if errCreateOrder != nil {
		fmt.Println("Something went wrong while trying create order, see logs below.")
		fmt.Println(errCreateOrder.Error())
		return errors.New(errCreateOrder.Error())
	}

	return nil

}

func (h *OrderService) UpdateOrder(orderID uint, payload resources.UpdateOrder) error {

	errUpdateOrder := h.orderRepo.UpdateOrder(orderID, payload)

	if errUpdateOrder != nil {
		fmt.Println("Something went wrong while trying update order, see logs below.")
		fmt.Println(errUpdateOrder.Error())
		return errors.New(errUpdateOrder.Error())
	}

	return nil

}

func (h *OrderService) GetOrder(orderId uint) (*models.Order, error) {

	resultGetOrder, errGetOrder := h.orderRepo.GetOrderByID(orderId)

	if errGetOrder != nil {
		fmt.Println("Something went wrong while trying get order, see logs below.")
		fmt.Println(errGetOrder.Error())
		return nil, errors.New(errGetOrder.Error())
	}

	return resultGetOrder, nil

}

func (h *OrderService) DeleteOrder(orderId uint) error {

	errGetOrder := h.orderRepo.DeleteOrder(orderId)

	if errGetOrder != nil {
		fmt.Println("Something went wrong while trying get order, see logs below.")
		fmt.Println(errGetOrder.Error())
		return errors.New(errGetOrder.Error())
	}

	return nil

}
