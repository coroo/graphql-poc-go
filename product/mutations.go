package product

import (
	"github.com/graphql-go/graphql"
	"graphql-poc-go/product/types"
	"graphql-poc-go/product/entities"
	"graphql-poc-go/product/usecases"
	dtoTypes "graphql-poc-go/app/dto/types"
)
type ProductMutation interface {
	CreateProductMutation() *graphql.Field
	UpdateProductMutation() *graphql.Field
	DeleteProductMutation() *graphql.Field
}

type productMutation struct {
	usecases		usecases.ProductService
}

func NewProductMutation(usecase usecases.ProductService) ProductMutation {
	return &productMutation{
		usecases: usecase,
	}
}

func (mutations *productMutation) CreateProductMutation() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProductType,
		Description: "Create new product",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			product, err := usecases.ProductService.SaveProduct(mutations.usecases, params)
			return product, err
		},
	}
}

func (mutations *productMutation) UpdateProductMutation() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProductType,
		Description: "Update product by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"description": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"price": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			slug, _ := params.Args["slug"].(string)
			name, nameOk := params.Args["name"].(string)
			description, descriptionOk := params.Args["description"].(string)
			summary, summaryOk := params.Args["summary"].(string)
			price, priceOk := params.Args["price"].(float64)
			product := entities.Product{}
			products := []entities.Product{}
			// utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)
			for i, p := range products {
				if string(slug) == p.Slug {
					if nameOk {
						products[i].Name = name
					}
					if descriptionOk {
						products[i].Description = description
					}
					if summaryOk {
						products[i].Summary = summary
					}
					if priceOk {
						products[i].StartPremiumFrom = price
					}
					product = products[i]
					break
				}
			}
			return product, nil
		},
	}
}

func (mutations *productMutation) DeleteProductMutation() *graphql.Field {
	return &graphql.Field{
		Type:        dtoTypes.DeleteType,
		Description: "Delete product by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			response, err := usecases.ProductService.DeleteProduct(mutations.usecases, params)
			return response, err
		},
	}
}