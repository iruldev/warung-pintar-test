package app

import (
	"github.com/iruldev/warung-pintar-test/product-service/controller"
	"github.com/iruldev/warung-pintar-test/product-service/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/products", productController.FindAll)
	router.GET("/api/products/:productId", productController.FindById)

	router.PanicHandler = exception.ErrorHandler
	return router
}
