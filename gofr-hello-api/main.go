package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()

	app.GET("/hello", func(c *gofr.Context) (interface{}, error) {
		return "Hello World!", nil
	})

	app.Run()
}
