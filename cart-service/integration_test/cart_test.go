package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/iruldev/warung-pintar-test/cart-service/app"
	"github.com/iruldev/warung-pintar-test/cart-service/controller"
	"github.com/iruldev/warung-pintar-test/cart-service/service"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var BaseURL = fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))

func setupRouter() http.Handler {
	redis := app.NewDB()
	cartService := service.NewCartService(redis)
	cartController := controller.NewCartController(cartService)
	return app.NewRouter(cartController)
}

func flushRedis()  {
	redis := app.NewDB()
	redis.FlushAll()
}

func TestAddToCartSuccess(t *testing.T) {
	flushRedis()

	requestBody := strings.NewReader(`[
		{
			"product_id": 1,
			"quantity": 5
		},
		{
			"product_id": 5,
			"quantity": 5
		}
	]`)

	router := setupRouter()
	request := httptest.NewRequest(http.MethodPost, BaseURL + "/api/carts", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestAddToCartFailed(t *testing.T) {
	requestBody := strings.NewReader(`[
		{
			"product_id": 1,
			"quantity": 5
		},
		{
			"product_id": 9999,
			"quantity": 5
		}
	]`)

	router := setupRouter()
	request := httptest.NewRequest(http.MethodPost, BaseURL + "/api/carts", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}

func TestGetCartSuccess(t *testing.T) {
	flushRedis()

	// Add Cart
	redis := app.NewDB()
	cartService := service.NewCartService(redis)
	cart := cartService.AddToCart(context.Background(), []service.CartCreateRequest{
		{
			ProductID: 2,
			Quantity:  6,
		},
		{
			ProductID: 5,
			Quantity:  10,
		},
	})

	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/carts", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, cart.ID, responseBody["data"].(map[string]interface{})["id"])
}

func TestGetCartFailed(t *testing.T) {
	flushRedis()

	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/carts", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not found", responseBody["status"])
}
