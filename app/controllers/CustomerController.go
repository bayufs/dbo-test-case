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

type CustomerController struct {
	customerService services.CustomerService
}

func NewCustomerController() *CustomerController {
	return &CustomerController{
		customerService: *services.NewCustomerService(),
	}
}

func (h *CustomerController) Index(c *gin.Context) {

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

	metaRes, result, err := h.customerService.GetCustomerList(queryStringParam)

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

	var customerResources []resources.CustomerDTO

	if len(result) > 0 {
		data := result
		for _, customer := range data {
			customerResource := resources.CustomerDTO{
				ID:        customer.ID,
				FirstName: customer.FirstName,
				LastName:  customer.LastName,
				Email:     customer.Email,
				Phone:     customer.Phone,
				Address:   customer.Address,
				CreatedAt: customer.CreatedAt,
				UpdatedAt: customer.UpdatedAt,
			}

			customerResources = append(customerResources, customerResource)
		}
	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success get customer list.", customerResources, meta),
	)

	return

}

func (h *CustomerController) Show(c *gin.Context) {

	customerID := c.Param("customerID")

	customerID64, err := strconv.ParseUint(customerID, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	customerIDUINT := uint(customerID64)

	result, err := h.customerService.GetCustomer(customerIDUINT)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", err.Error()),
		)
		return

	}

	var customerResource resources.CustomerDTO

	if result != nil {
		customerResource = resources.CustomerDTO{
			ID:        result.ID,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     result.Email,
			Phone:     result.Phone,
			Address:   result.Address,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success get customer.", customerResource),
	)

	return

}

func (h *CustomerController) Store(c *gin.Context) {

	var payload resources.StoreNewCustomer

	var errMsg string

	if err := c.ShouldBindJSON(&payload); err != nil {

		switch err.(type) {
		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "FirstName" && fieldError.Tag() == "required" {
					errMsg = "Firstname is required"
				} else if fieldError.Field() == "LastName" && fieldError.Tag() == "required" {
					errMsg = "Lastname required"
				} else if fieldError.Field() == "Email" && fieldError.Tag() == "required" {
					errMsg = "Email required"
				} else if fieldError.Field() == "Phone" && fieldError.Tag() == "required" {
					errMsg = "Phone required"
				} else if fieldError.Field() == "Address" && fieldError.Tag() == "required" {
					errMsg = "Address required"
				} else if fieldError.Field() == "Username" && fieldError.Tag() == "required" {
					errMsg = "Username required"
				} else if fieldError.Field() == "Password" && fieldError.Tag() == "required" {
					errMsg = "Password required"
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

	result := h.customerService.StoreNewCustomer(payload)

	if result != nil {

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)

		return

	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success store new customer."),
	)
	return
}

func (h *CustomerController) Update(c *gin.Context) {

	customerID := c.Param("customerID")

	customerID64, err := strconv.ParseUint(customerID, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	customerIDUINT := uint(customerID64)

	var payload resources.StoreNewCustomer

	var errMsg string

	if err := c.ShouldBindJSON(&payload); err != nil {

		switch err.(type) {
		case validator.ValidationErrors:
			validationErrors := err.(validator.ValidationErrors)
			for _, fieldError := range validationErrors {
				if fieldError.Field() == "FirstName" && fieldError.Tag() == "required" {
					errMsg = "Firstname is required"
				} else if fieldError.Field() == "LastName" && fieldError.Tag() == "required" {
					errMsg = "Lastname required"
				} else if fieldError.Field() == "Email" && fieldError.Tag() == "required" {
					errMsg = "Email required"
				} else if fieldError.Field() == "Phone" && fieldError.Tag() == "required" {
					errMsg = "Phone required"
				} else if fieldError.Field() == "Address" && fieldError.Tag() == "required" {
					errMsg = "Address required"
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

	result := h.customerService.UpdateCustomer(customerIDUINT, payload)

	if result != nil {

		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,
			helpers.CreateResponse(http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg),
		)

		return

	}

	c.AbortWithStatusJSON(http.StatusOK,
		helpers.CreateResponse(http.StatusOK, "success", "Success update customer data."),
	)
	return
}
