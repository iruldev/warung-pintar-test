package repository

import (
	"context"
	"github.com/iruldev/warung-pintar-test/product-service/model/domain"
)

type ProductRepository interface {
	FindById(ctx context.Context, productId int) (domain.Product, error)
	FindAll(ctx context.Context) []domain.Product
}
