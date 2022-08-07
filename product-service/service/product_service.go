package service

import (
	"context"
	"github.com/iruldev/warung-pintar-test/product-service/model/web"
)

type ProductService interface {
	FindById(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
