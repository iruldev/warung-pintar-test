package service

import (
	"context"
	"github.com/iruldev/warung-pintar-test/product-service/exception"
	"github.com/iruldev/warung-pintar-test/product-service/model/web"
	"github.com/iruldev/warung-pintar-test/product-service/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	product, err := service.ProductRepository.FindById(ctx, productId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ProductResponse{
		Id:     product.Id,
		Name:   product.Name,
		Price: 	product.Price,
		Stock:  product.Stock,
		Weight: product.Weight,
	}
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	products := service.ProductRepository.FindAll(ctx)

	var ProductRepository []web.ProductResponse
	for _, product := range products {
		ProductRepository = append(ProductRepository, web.ProductResponse{
			Id:     product.Id,
			Name:   product.Name,
			Price: 	product.Price,
			Stock:  product.Stock,
			Weight: product.Weight,
		})
	}

	return ProductRepository
}
