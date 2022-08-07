package app

import (
	"github.com/iruldev/warung-pintar-test/cart-service/controller"
	"github.com/iruldev/warung-pintar-test/cart-service/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(cartController controller.ICartController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/carts", cartController.GetCart)
	router.POST("/api/carts", cartController.AddToCart)

	router.PanicHandler = exception.ErrorHandler
	return router
}