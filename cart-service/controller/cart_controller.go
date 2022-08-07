package controller

import (
	"encoding/json"
	"github.com/iruldev/warung-pintar-test/cart-service/helper"
	"github.com/iruldev/warung-pintar-test/cart-service/model/web"
	"github.com/iruldev/warung-pintar-test/cart-service/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ICartController interface {
	GetCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CartController struct{
	CartService service.ICartService
}

func NewCartController(cartService service.ICartService) ICartController {
	return &CartController{
		CartService: cartService,
	}
}

func (controller *CartController) GetCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartResponse := controller.CartService.GetCart(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CartController) AddToCart(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var cartCreateRequest []service.CartCreateRequest

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&cartCreateRequest)
	helper.PanicIfError(err)

	cartResponse := controller.CartService.AddToCart(request.Context(), cartCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   cartResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
