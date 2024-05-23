package controllers

import (
	"dbo-test-case/app/helpers"
	"dbo-test-case/app/resources"
	"dbo-test-case/app/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		orderService: *services.NewOrderService(),
	}
}

func (h *OrderController) Index(c *gin.Context) {

	sortParam := c.DefaultQuery("sort", "DESC")

	pageParam := c.Query("page")

	limitParam := c.DefaultQuery("limit", "10")

	offsetParam := c.Query("offset")

	search := c.Query("search")

	var queryStringParam = map[string]interface{}{

		"sort":   sortParam,
		"page":   pageParam,
		"limit":  limitParam,
		"offset": offsetParam,
		"search": search,
	}

	metaRes, result, err := h.orderService.GetOrderList(queryStringParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error()),
		)
		return

	}

	meta := map[string]interface{}{
		"count":  metaRes["count"],
		"offset": metaRes["offset"],
		"limit":  metaRes["limit"],
		"total":  metaRes["total"],
	}

	var orderResource resources.OrderDTO

	if len(result) > 0 {
		data := result
		for _, order := range data {
			customerResource := resources.CustomerDTO{
				ID:        order.Customer.ID,
				FirstName: order.Customer.FirstName,
				LastName:  order.Customer.LastName,
				Email:     order.Customer.Email,
				Phone:     order.Customer.Phone,
				Address:   order.Customer.Address,
				CreatedAt: order.Customer.CreatedAt,
				UpdatedAt: order.Customer.UpdatedAt,
			}

			productResource := resources.ProductDTO{
				ID:          order.OrderItem.Product.ID,
				ProductName: order.OrderItem.Product.ProductName,
				Description: order.OrderItem.Product.Description,
				Price:       order.OrderItem.Product.Price,
				Stock:       order.OrderItem.Product.Stock,
				CreatedAt:   order.OrderItem.Product.CreatedAt,
				UpdatedAt:   order.OrderItem.Product.UpdatedAt,
			}

			orderItemResource := resources.OrderItemDTO{
				ID:        order.OrderItem.ID,
				Product:   productResource,
				Quantity:  order.OrderItem.Quantity,
				UnitPrice: order.OrderItem.UnitPrice,
				CreatedAt: order.OrderItem.CreatedAt,
				UpdatedAt: order.OrderItem.UpdatedAt,
			}

			orderResource = resources.OrderDTO{
				ID:          order.ID,
				OrderDate:   order.OrderDate,
				Customer:    customerResource,
				TotalAmount: order.TotalAmount,
				Status:      order.Status,
				OrderItem:   orderItemResource,
				CreatedAt:   order.CreatedAt,
				UpdatedAt:   order.UpdatedAt,
			}

		}
	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success get order list.", orderResource, meta),
	)

	return

}

func (h *OrderController) Show(c *gin.Context) {

	orderID := c.Param("orderID")

	orderID64, err := strconv.ParseUint(orderID, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	orderID64UINT := uint(orderID64)

	result, err := h.orderService.GetOrder(orderID64UINT)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error()),
		)
		return

	}

	var orderResource resources.OrderDTO

	if result != nil {

		order := result

		customerResource := resources.CustomerDTO{
			ID:        order.Customer.ID,
			FirstName: order.Customer.FirstName,
			LastName:  order.Customer.LastName,
			Email:     order.Customer.Email,
			Phone:     order.Customer.Phone,
			Address:   order.Customer.Address,
			CreatedAt: order.Customer.CreatedAt,
			UpdatedAt: order.Customer.UpdatedAt,
		}

		productResource := resources.ProductDTO{
			ID:          order.OrderItem.Product.ID,
			ProductName: order.OrderItem.Product.ProductName,
			Description: order.OrderItem.Product.Description,
			Price:       order.OrderItem.Product.Price,
			Stock:       order.OrderItem.Product.Stock,
			CreatedAt:   order.OrderItem.Product.CreatedAt,
			UpdatedAt:   order.OrderItem.Product.UpdatedAt,
		}

		orderItemResource := resources.OrderItemDTO{
			ID:        order.OrderItem.ID,
			Product:   productResource,
			Quantity:  order.OrderItem.Quantity,
			UnitPrice: order.OrderItem.UnitPrice,
			CreatedAt: order.OrderItem.CreatedAt,
			UpdatedAt: order.OrderItem.UpdatedAt,
		}

		orderResource = resources.OrderDTO{
			ID:          order.ID,
			OrderDate:   order.OrderDate,
			Customer:    customerResource,
			TotalAmount: order.TotalAmount,
			Status:      order.Status,
			OrderItem:   orderItemResource,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		}

	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success get order.", orderResource),
	)

	return

}

func (h *OrderController) Store(c *gin.Context) {

	var payload resources.CreateOrder

	var errMsg string

	if err := c.ShouldBindJSON(&payload); err != nil {

		switch err.(type) {
		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "CustomerID" && fieldError.Tag() == "required" {
					errMsg = "Customer ID is required"
				} else if fieldError.Field() == "TotalAmount" && fieldError.Tag() == "required" {
					errMsg = "Total amount required"
				} else if fieldError.Field() == "ProductID" && fieldError.Tag() == "required" {
					errMsg = "ProductID required"
				} else if fieldError.Field() == "Quantity" && fieldError.Tag() == "required" {
					errMsg = "Quantity required"
				} else if fieldError.Field() == "UnitPrice" && fieldError.Tag() == "required" {
					errMsg = "UnitPrice required"
				} else {
					errMsg = "Invalid request"
				}
			}

		default:
			errMsg = "Invalid request"

		}

		c.AbortWithStatusJSON(http.StatusOK,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Credentials error", errMsg),
		)

		return

	}

	result := h.orderService.CreateOrder(payload)

	if result != nil {

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)

		return

	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success creat order."),
	)
	return
}

func (h *OrderController) Update(c *gin.Context) {

	orderID := c.Param("orderID")

	orderID64, err := strconv.ParseUint(orderID, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	orderID64UINT := uint(orderID64)

	var payload resources.UpdateOrder

	var errMsg string

	if err := c.ShouldBindJSON(&payload); err != nil {

		switch err.(type) {
		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "CustomerID" && fieldError.Tag() == "required" {
					errMsg = "Customer ID is required"
				} else if fieldError.Field() == "TotalAmount" && fieldError.Tag() == "required" {
					errMsg = "Total amount required"
				} else if fieldError.Field() == "ProductID" && fieldError.Tag() == "required" {
					errMsg = "ProductID required"
				} else if fieldError.Field() == "Quantity" && fieldError.Tag() == "required" {
					errMsg = "Quantity required"
				} else if fieldError.Field() == "UnitPrice" && fieldError.Tag() == "required" {
					errMsg = "UnitPrice required"
				} else if fieldError.Field() == "Status" && fieldError.Tag() == "required" {
					errMsg = "Status required"
				} else {
					errMsg = "Invalid request"
				}
			}

		default:
			errMsg = "Invalid request"

		}

		c.AbortWithStatusJSON(http.StatusOK,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Credentials error", errMsg),
		)

		return

	}

	result := h.orderService.UpdateOrder(orderID64UINT, payload)

	if result != nil {

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)

		return

	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success update order."),
	)
	return
}

func (h *OrderController) Delete(c *gin.Context) {

	orderID := c.Param("orderID")

	orderID64, err := strconv.ParseUint(orderID, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	orderID64UINT := uint(orderID64)

	errDelete := h.orderService.DeleteOrder(orderID64UINT)

	if errDelete != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error()),
		)
		return

	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success delete order."),
	)

	return

}
