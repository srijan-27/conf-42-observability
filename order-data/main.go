package main

import (
	"gofr.dev/pkg/gofr"

	"order-data/handler"
	"order-data/migration"
	"order-data/store"
)

func main() {
	app := gofr.New()

	app.Migrate(migration.All())

	s := store.New()
	h := handler.New(s)

	// Add required routes
	app.POST("/orders", h.Create)
	app.GET("/orders", h.GetAll)
	app.GET("/orders/{id}", h.GetByID)
	app.PUT("/orders/{id}", h.Update)
	app.DELETE("/orders/{id}", h.Delete)

	app.Run()
}
