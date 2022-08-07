package integration_test

import (
	"encoding/json"
	"fmt"
	"github.com/iruldev/warung-pintar-test/product-service/app"
	"github.com/iruldev/warung-pintar-test/product-service/controller"
	"github.com/iruldev/warung-pintar-test/product-service/repository"
	"github.com/iruldev/warung-pintar-test/product-service/service"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var BaseURL = fmt.Sprintf("http://localhost:%s", os.Getenv("PORT"))

func setupRouter() http.Handler {
	productRepository := repository.NewProductRepository()
	productyService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productyService)
	return app.NewRouter(productController)
}

func TestGetProductSuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/products/1", nil)
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

func TestGetProductFailed(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/products/404", nil)
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

func TestListProductsSuccess(t *testing.T) {
	router := setupRouter()

	request := httptest.NewRequest(http.MethodGet, BaseURL + "/api/products", nil)
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