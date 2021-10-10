package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"io/ioutil"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Product contains descriptionrmation about one product
type Product struct {
	Slug    string   `json:"slug"`
	Name  string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	Summary  string  `json:"summary"`
	StartAgeFrom float64 `json:"price"`
}

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"slug": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"summary": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var products = getJson("https://staging-product.superyou.co.id/api/v1/products")
var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string) []Product {
    r, err := myClient.Get(url)
    if err != nil {
        fmt.Println(err)
    }
    defer r.Body.Close()
	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the responseData to type string

    var responseObject []Product
    json.Unmarshal(responseData, &responseObject)

    // for i := 0; i < len(responseObject); i++ {
    //     fmt.Println(responseObject[i])
    // }
	// sb := string(responseData)

    return responseObject
}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single product by slug
			   http://localhost:8080/product?query={product(slug:1){name,description,price}}
			*/
			"product": &graphql.Field{
				Type:        productType,
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
						for _, product := range products {
							if string(product.Slug) == slug {
								return product, nil
							}
						}
					}
					return nil, nil
				},
			},
			/* Get (read) product list
			   http://localhost:8080/product?query={list{slug,name,description,price}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get product list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return products, nil
				},
			},
		},
	})

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		/* Create new product item
		http://localhost:8080/product?query=mutation+_{create(name:"Inca Kola",description:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){slug,name,description,price}}
		*/
		"create": &graphql.Field{
			Type:        productType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
				product := Product{
					Slug:    params.Args["slug"].(string),
					Name:  params.Args["name"].(string),
					Description:  params.Args["description"].(string),
					Summary:  params.Args["summary"].(string),
					StartAgeFrom: params.Args["price"].(float64),
				}
				products = append(products, product)
				return product, nil
			},
		},

		/* Update product by slug
		   http://localhost:8080/product?query=mutation+_{update(slug:1,price:3.95){slug,name,description,price}}
		*/
		"update": &graphql.Field{
			Type:        productType,
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
				product := Product{}
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
							products[i].StartAgeFrom = price
						}
						product = products[i]
						break
					}
				}
				return product, nil
			},
		},

		/* Delete product by slug
		   http://localhost:8080/product?query=mutation+_{delete(slug:1){slug,name,description,price}}
		*/
		"delete": &graphql.Field{
			Type:        productType,
			Description: "Delete product by slug",
			Args: graphql.FieldConfigArgument{
				"slug": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				slug, _ := params.Args["slug"].(string)
				product := Product{}
				for i, p := range products {
					if string(slug) == p.Slug {
						product = products[i]
						// Remove from product list
						products = append(products[:i], products[i+1:]...)
					}
				}

				return product, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

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

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Println("Server ready at http://localhost/8080/graphql")
	http.ListenAndServe(":8080", nil)
}