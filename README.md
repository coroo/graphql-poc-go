Installation

```
go run main.go
```

```
# Write your query or mutation here
{
  product(slug: "super-care-protection") {
    name,
    summary,
    available_claim_methods,
    bundling_with_rider,
    slug,
    riders {
      name,
      slug,
      is_active
    },
    category {
      name
    },
    insurance_type {
      name
    },
    benefit_groups {
      order,
      name,
      benefits {
        name,
        order
      }
    }
  }
}
```

```
mutation{
  createProduct(name:"test", slug:"name", price:1){
    id
  }
}
```





NOTES 
SET TO GLOBAL VARIABLE
```
package product

import(
	"fmt"

	"github.com/graphql-go/graphql"
	"graphql-poc-go/product/types"
	"graphql-poc-go/app/utils"
)


// var products = getJson("https://staging-product.superyou.co.id/api/v1/products")
var products = new([]Product)
var temp = utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)

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
				// product := new(Product)
				// utils.GetJson("https://staging-product.superyou.co.id/api/v1/products/slug/"+slug, product)
				// if string(product.Slug) == slug {
				// 	return product, nil
				// }
				// products := new([]Product)
				// utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)
				// fmt.Println(products)

				for _, product := range *products {
					if string(product.Slug) == slug {
						return product, nil
					}
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
			// products := new([]Product)
			// utils.GetJson("http://0.0.0.0:8016/api/v1/products/", products)
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
```