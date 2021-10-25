package usecases

import (
	"graphql-poc-go/product/entities"
	"graphql-poc-go/app/utils"
)

type ProductService interface {
	// SaveProduct(entities.Product) (int, error)
	// UpdateProduct(entities.Product) error
	// DeleteProduct(entities.Product) error
	GetAllProducts() *[]entities.Product
	GetProduct(id string) *entities.Product
}

type productService struct {}

func NewProductService() ProductService {
	return &productService{}
}

func (usecase *productService) GetAllProducts() *[]entities.Product {
	products := new([]entities.Product)
	utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)
	return products
}

func (usecase *productService) GetProduct(slug string) *entities.Product {
	product := new(entities.Product)
	utils.GetJson("http://0.0.0.0:8016/api/v1/products/slug/"+slug, product)
	if string(product.Slug) == slug {
		return product
	}
	return nil
}

// func (usecases *productService) SaveProduct(product entities.Product) (int, error) {
// 	return usecases.repositories.SaveProduct(product)
// }

// func (usecases *productService) UpdateProduct(product entities.Product) error {
// 	return usecases.repositories.UpdateProduct(product)
// }

// func (usecases *productService) DeleteProduct(product entities.Product) error {
// 	return usecases.repositories.DeleteProduct(product)
// }