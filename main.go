package main

import (
	"context"
	"customer-api/src/controller"
	"customer-api/src/service"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()

	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	notificationService := service.NewNotificationService(
		config,
		os.Getenv("CreatedCustomerTopicARN"),
	)
	customerController := controller.NewCustomerController(
		notificationService,
	)

	e.POST("/customers", customerController.CreateCustomer)

	e.Start(":5000")
}
