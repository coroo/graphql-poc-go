package product

import(
	"fmt"

	"github.com/graphql-go/graphql"
	"graphql-poc-go/product/types"
	"graphql-poc-go/product/entities"
	"graphql-poc-go/app/utils"
)

func GetProductQuery() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProductType,
		Description: "Get product by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			slug, ok := p.Args["slug"].(string)
			if ok {
				// Find product
				product := new(entities.Product)
				utils.GetJson("http://0.0.0.0:8016/api/v1/products/slug/"+slug, product)
				if string(product.Slug) == slug {
					return product, nil
				}
			}
			return nil, nil
		},
	}
}

func GetProductsQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.ProductType),
		Description: "Get product list",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			products := new([]entities.Product)
			utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)
			return products, nil
		},
	}
}

func GetRootQueryFields() graphql.Fields {
	return graphql.Fields{
		"product": GetProductQuery(),
		"list": GetProductsQuery(),
	}
}


func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}