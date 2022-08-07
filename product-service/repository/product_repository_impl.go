package repository

import (
	"context"
	"errors"
	"github.com/iruldev/warung-pintar-test/product-service/model/domain"
)

type ProductRepositoryImpl struct {}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

var products = []domain.Product{
	{
		Id : 1,
		Name : "Indomie",
		Price: 3000,
		Stock : 50,
		Weight: 200,
	},
	{
		Id : 2,
		Name : "Oreo",
		Price: 10000,
		Stock : 44,
		Weight: 150,
	},
	{
		Id : 3,
		Name : "Mie Sedaap",
		Price: 2000,
		Stock : 100,
		Weight: 200,
	},
	{
		Id : 4,
		Name : "Sarimi isi 2",
		Price: 4000,
		Stock : 80,
		Weight: 400,
	},
	{
		Id : 5,
		Name : "Bimoli",
		Price: 14000,
		Stock : 66,
		Weight: 1000,
	},
	{
		Id : 6,
		Name : "Belibis",
		Price: 45000,
		Stock : 55,
		Weight: 1000,
	},
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId int) (domain.Product, error) {
	for _, product := range products {
		if productId == product.Id {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("product is not found")
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) []domain.Product {
	return products
}