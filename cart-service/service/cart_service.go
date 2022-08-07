package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/iruldev/warung-pintar-test/cart-service/exception"
	"github.com/iruldev/warung-pintar-test/cart-service/helper"
	"io"
	"net/http"
	"os"
	"strconv"
)

type ICartService interface {
	GetExistingCart() (string, error)
	GetCart(ctx context.Context) CartResponse
	AddToCart(ctx context.Context, cartCreateRequest []CartCreateRequest) CartResponse
}

type CartServiceImpl struct {
	*redis.Client
}

func NewCartService(client *redis.Client) ICartService {
	return &CartServiceImpl{
		Client: client,
	}
}

type CartCreateRequest struct {
	ProductID	int `json:"product_id"`
	Quantity	int `json:"quantity"`
}

type Product struct {
	Id		int		`json:"id"`
	Name 	string 	`json:"name"`
	Price	int 	`json:"price"`
	Stock	int 	`json:"stock"`
	Weight	int 	`json:"weight"`
}

type Cart struct {
	ProductID	int 		`json:"product_id"`
	Quantity	int 		`json:"quantity"`
	Product 	Product 	`json:"product"`
}

type CartResponse struct {
	ID			string	`json:"id"`
	Carts		[]Cart 	`json:"carts"`
}

func (service *CartServiceImpl) GetExistingCart() (string, error) {
	val, err := service.Client.Keys("*").Result()
	if err != nil {
		return "", err
	}

	if len(val) == 0 {
		return "", nil
	}

	return val[0], nil
}

func (service *CartServiceImpl) GetCart(ctx context.Context) CartResponse {
	existingCart, err := service.GetExistingCart()
	helper.PanicIfError(err)

	if existingCart == "" {
		panic(exception.NewNotFoundError("Cart not found"))
	}

	cart, err := service.Client.Get(existingCart).Result()
	helper.PanicIfError(err)

	var carts []Cart
	json.Unmarshal([]byte(cart), &carts)

	return CartResponse{
		ID:    existingCart,
		Carts: carts,
	}
}


func (service *CartServiceImpl) AddToCart(ctx context.Context, cartCreateRequest []CartCreateRequest) CartResponse {
	var carts []Cart

	baseURL := fmt.Sprintf("http://%s", os.Getenv("PRODUCT_SERVICE_HOST"))
	for _, request := range cartCreateRequest {
		res, err := http.Get(baseURL + "/api/products/" + strconv.Itoa(request.ProductID))
		helper.PanicIfError(err)
		if res.StatusCode != 200 {
			panic(exception.NewNotFoundError("Product is not found"))
		}

		body, _ := io.ReadAll(res.Body)
		var responseBody map[string]interface{}
		json.Unmarshal(body, &responseBody)
		product := responseBody["data"].(map[string]interface{})

		cart := Cart {
			ProductID: int(product["id"].(float64)),
			Quantity:  request.Quantity,
			Product:  Product{
				Id:     int(product["id"].(float64)),
				Name:   product["name"].(string),
				Price:  int(product["price"].(float64)),
				Stock:  int(product["stock"].(float64)),
				Weight: int(product["weight"].(float64)),
			},
		}
		carts = append(carts, cart)
	}
	cartsByte, err := json.Marshal(carts)
	helper.PanicIfError(err)

	service.Client.FlushAll()

	uid := uuid.NewString()
	service.Client.Set(uid, string(cartsByte), 0)

	return CartResponse{
		ID:    uid,
		Carts: carts,
	}
}

