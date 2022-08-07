package main

import (
	"fmt"
	"github.com/iruldev/warung-pintar-test/cart-service/app"
	"github.com/iruldev/warung-pintar-test/cart-service/controller"
	"github.com/iruldev/warung-pintar-test/cart-service/helper"
	"github.com/iruldev/warung-pintar-test/cart-service/service"
	"net/http"
	"os"
)

func main()  {
	redis := app.NewDB()
	cartService := service.NewCartService(redis)
	cartController := controller.NewCartController(cartService)

	router := app.NewRouter(cartController)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}