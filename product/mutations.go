package product

import (
	"log"
	"encoding/json"

	"github.com/graphql-go/graphql"
	"graphql-poc-go/app/utils"
	"graphql-poc-go/product/types"
	"graphql-poc-go/product/entities"
)

func CreateProductMutation() *graphql.Field {
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
			// rand.Seed(time.Now().UnixNano())
			// product := Product{
			// 	Slug:    params.Args["slug"].(string),
			// 	Name:  params.Args["name"].(string),
			// 	// Description:  params.Args["description"].(string),
			// 	// Summary:  params.Args["summary"].(string),
			// 	// StartPremiumFrom: params.Args["price"].(float64),
			// }
			// fmt.Println(product)
			// products := []entities.Product{}
			// products = append(products, product)


			//Encode the data
			postBody, _ := json.Marshal(map[string]string{
				"slug":    			params.Args["slug"].(string),
				"name":    			params.Args["name"].(string),
				"doc_name":    		params.Args["name"].(string),
				"product_type":    	params.Args["name"].(string),
			})
			//Leverage Go's HTTP Post function to make request

			product := new(entities.Product)
			utils.PostJson("http://0.0.0.0:8016/api/v1/products/", product, postBody)
			if string(product.Id) == "" {
				log.Println("Error Occured")
			}
			return product, nil
			// return product, nil
		},
	}
}

func UpdateProductMutation() *graphql.Field {
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
			utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)
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

func DeleteProductMutation() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProductType,
		Description: "Delete product by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			slug, _ := params.Args["slug"].(string)
			product := entities.Product{}
			products := []*entities.Product{}
			utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)
			for i, p := range products {
				if string(slug) == p.Slug {
					// product = products[i]
					// Remove from product list
					products = append(products[:i], products[i+1:]...)
				}
			}

			return product, nil
		},
	}
}