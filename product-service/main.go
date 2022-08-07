package main

import (
	"fmt"
	"github.com/iruldev/warung-pintar-test/product-service/app"
	"github.com/iruldev/warung-pintar-test/product-service/controller"
	"github.com/iruldev/warung-pintar-test/product-service/helper"
	"github.com/iruldev/warung-pintar-test/product-service/repository"
	"github.com/iruldev/warung-pintar-test/product-service/service"
	"net/http"
	"os"
)

func main() {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(productController)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
