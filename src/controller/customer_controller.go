package controller

import (
	"context"
	"customer-api/src/contract"
	"customer-api/src/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type customerController struct {
	notificationService contract.NotificationService
}

func NewCustomerController(notificationService contract.NotificationService) customerController {
	return customerController{
		notificationService: notificationService,
	}
}

func (c customerController) CreateCustomer(ctx echo.Context) error {
	var customer model.Customer

	err := ctx.Bind(&customer)
	if err != nil {
		fmt.Println(err)
		return err
	}

	customer.UUID, err = uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return err
	}

	customerJSON, err := json.Marshal(customer)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = c.notificationService.Publish(context.TODO(), string(customerJSON))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.JSON(http.StatusCreated, customer)
}
