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


type ProductRider struct {
    Id					string	`json:"id"`
    Slug				string	`json:"slug"`
    Name				string	`json:"name"`
    IsActive			bool	`json:"is_active"`
    ParentId			string	`json:"parent_id"`
    Summary				string	`json:"summary"`
    Description			string	`json:"description"`
    IconSvg				string	`json:"icon_svg"`
    CoveragePeriod		string	`json:"coverage_period"`
}

type ProductBenefitGroup struct {
    Id					string	`json:"id"`
    TooltipText			string	`json:"tooltip_text"`
    Name				string	`json:"name"`
    Order				int	`json:"order"`
	ProductBenefit		[]ProductBenefit `json:"benefits"`
}

type ProductBenefit struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
    IconSvg				string	`json:"icon_svg"`
    IconEtc				string	`json:"icon_etc"`
    ProductId			string	`json:"product_id"`
    ProductBenefitGroupId	string	`json:"product_benefit_group_id"`
    TooltipText			string	`json:"tooltip_text"`
    TooltipTextDescription	string	`json:"tooltip_text_description"`
    Order				int	`json:"order"`
}

type ProductCategory struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
}

type ProductInsuranceType struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
}
type Product struct {
    Id						string	`json:"id"`
    Slug					string	`json:"slug"`
    Name					string	`json:"name"`
    DocName					string	`json:"doc_name"`
    ParentId				string	`json:"parent_id"`
    IsActive		 		bool	`json:"is_active"`
    Featured		 		bool	`json:"featured"`
    BundlingWithRider		bool	`json:"bundling_with_rider"`
    Subheading				string	`json:"subheading"`
    Summary					string	`json:"summary"`
    Description				string	`json:"description"`
    IconSvg					string	`json:"icon_svg"`
    IconEtc					string	`json:"icon_etc"`
    RipLink					string	`json:"rip_link"`
    ProductType				string	`json:"product_type"`
    CoveragePeriod			string	`json:"coverage_period"`
    AvailableClaimMethods	[]string	`json:"available_claim_methods"`
    CovidCoverage		 	bool	`json:"covid_coverage"`
    StartAgeFrom			int	`json:"start_age_from"`
	StartPremiumFrom 		float64	`json:"start_premium_from"`
    Category				ProductCategory	`json:"category"`
    InsuranceType			ProductInsuranceType	`json:"insurance_type"`
    Riders					[]ProductRider	`json:"riders"`
    BenefitGroups			[]ProductBenefitGroup	`json:"benefit_groups"`
    // Tnc						List	`json:"tnc"`
    // Faq						List	`json:"faq"`
    // NotCoverage				List	`json:"not_coverage"`
    // Plans					List	`json:"plans"`
	// BELUM SELESAI
}

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{Type: graphql.String},
			"slug": &graphql.Field{Type: graphql.String},
			"name": &graphql.Field{Type: graphql.String},
			"doc_name": &graphql.Field{Type: graphql.String},
			"parent_id": &graphql.Field{Type: graphql.String},
			"is_active": &graphql.Field{Type: graphql.Boolean},
			"featured": &graphql.Field{Type: graphql.Boolean},
			"bundling_with_rider": &graphql.Field{Type: graphql.Boolean},
			"subheading": &graphql.Field{Type: graphql.String},
			"summary": &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"icon_svg": &graphql.Field{Type: graphql.String},
			"icon_etc": &graphql.Field{Type: graphql.String},
			"rip_link": &graphql.Field{Type: graphql.String},
			"product_type": &graphql.Field{Type: graphql.String},
			"coverage_period": &graphql.Field{Type: graphql.String},
			"available_claim_methods": &graphql.Field{Type: graphql.String},
			"covid_coverage": &graphql.Field{Type: graphql.Boolean},
			"start_age_from": &graphql.Field{Type: graphql.Int},
			"start_premium_from": &graphql.Field{Type: graphql.Float},
			"category": &graphql.Field{Type: ProductCategoryType},
			"insurance_type": &graphql.Field{Type: ProductInsuranceTypeType},
			"riders": &graphql.Field{Type: graphql.NewList(ProductRiderType)},
			"benefit_groups": &graphql.Field{Type: graphql.NewList(ProductBenefitGroupType)},
		},
	},
)

// ProductBenefitGroupType is the GraphQL schema for the Product type.
var ProductBenefitGroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductBenefitGroup",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
		"benefits": &graphql.Field{Type: graphql.NewList(ProductBenefitType)},
	},
})

// ProductBenefitType is the GraphQL schema for the Product type.
var ProductBenefitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductBenefit",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"icon_etc": &graphql.Field{Type: graphql.String},
		"product_id": &graphql.Field{Type: graphql.String},
		"product_benefit_group_id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"tooltip_text_description": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
	},
})

// ProductRiderType is the GraphQL schema for the Product type.
var ProductRiderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductRider",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"slug": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"is_active": &graphql.Field{Type: graphql.Boolean},
		"parent_id": &graphql.Field{Type: graphql.String},
		"summary": &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"coverage_period": &graphql.Field{Type: graphql.String},
	},
})

// ProductCategoryType is the GraphQL schema for the Product type.
var ProductCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductCategory",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// ProductInsuranceTypeType is the GraphQL schema for the Product type.
var ProductInsuranceTypeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductInsuranceType",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

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
					StartPremiumFrom: params.Args["price"].(float64),
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
							products[i].StartPremiumFrom = price
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