package main

import (
	"gofr.dev/pkg/gofr"

	"order-func/handler"
	"order-func/service"
)

func main() {
	app := gofr.New()

	app.AddHTTPService("order-data", app.Config.Get("DATA_SERVICE"))

	svc := service.New()
	h := handler.New(svc)

	// Add required routes
	app.POST("/orders", h.Create)
	app.GET("/orders", h.GetAll)
	app.GET("/orders/{id}", h.GetByID)
	app.PUT("/orders/{id}", h.Update)
	app.DELETE("/orders/{id}", h.Delete)

	app.Run()
}
