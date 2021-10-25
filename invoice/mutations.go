package invoice

import (
	"log"
	"encoding/json"

	"github.com/graphql-go/graphql"
	"graphql-poc-go/app/utils"
)

func CreateInvoiceMutation() *graphql.Field {
	return &graphql.Field{
		Type:        InvoiceType,
		Description: "Create new invoice",
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
			// invoice := Invoice{
			// 	Slug:    params.Args["slug"].(string),
			// 	Name:  params.Args["name"].(string),
			// 	// Description:  params.Args["description"].(string),
			// 	// Summary:  params.Args["summary"].(string),
			// 	// StartPremiumFrom: params.Args["price"].(float64),
			// }
			// fmt.Println(invoice)
			// invoices := []Invoice{}
			// invoices = append(invoices, invoice)


			//Encode the data
			postBody, _ := json.Marshal(map[string]string{
				"slug":    			params.Args["slug"].(string),
				"name":    			params.Args["name"].(string),
				"doc_name":    		params.Args["name"].(string),
				"invoice_type":    	params.Args["name"].(string),
			})
			//Leverage Go's HTTP Post function to make request

			invoice := new(Invoice)
			utils.PostJson("http://0.0.0.0:8016/api/v1/invoices/", invoice, postBody)
			if string(invoice.Id) == "" {
				log.Println("Error Occured")
			}
			return invoice, nil
			// return invoice, nil
		},
	}
}

// func UpdateInvoiceMutation() *graphql.Field {
// 	return &graphql.Field{
// 		Type:        InvoiceType,
// 		Description: "Update invoice by slug",
// 		Args: graphql.FieldConfigArgument{
// 			"slug": &graphql.ArgumentConfig{
// 				Type: graphql.NewNonNull(graphql.String),
// 			},
// 			"name": &graphql.ArgumentConfig{
// 				Type: graphql.String,
// 			},
// 			"description": &graphql.ArgumentConfig{
// 				Type: graphql.String,
// 			},
// 			"price": &graphql.ArgumentConfig{
// 				Type: graphql.Float,
// 			},
// 		},
// 		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 			slug, _ := params.Args["slug"].(string)
// 			name, nameOk := params.Args["name"].(string)
// 			description, descriptionOk := params.Args["description"].(string)
// 			summary, summaryOk := params.Args["summary"].(string)
// 			price, priceOk := params.Args["price"].(float64)
// 			invoice := Invoice{}
// 			invoices := []Invoice{}
// 			utils.GetJson("http://0.0.0.0:8016/api/v1/invoices/", invoices)
// 			for i, p := range invoices {
// 				if string(slug) == p.Slug {
// 					if nameOk {
// 						invoices[i].Name = name
// 					}
// 					if descriptionOk {
// 						invoices[i].Description = description
// 					}
// 					if summaryOk {
// 						invoices[i].Summary = summary
// 					}
// 					if priceOk {
// 						invoices[i].StartPremiumFrom = price
// 					}
// 					invoice = invoices[i]
// 					break
// 				}
// 			}
// 			return invoice, nil
// 		},
// 	}
// }

// func DeleteInvoiceMutation() *graphql.Field {
// 	return &graphql.Field{
// 		Type:        InvoiceType,
// 		Description: "Delete invoice by slug",
// 		Args: graphql.FieldConfigArgument{
// 			"slug": &graphql.ArgumentConfig{
// 				Type: graphql.NewNonNull(graphql.String),
// 			},
// 		},
// 		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
// 			slug, _ := params.Args["slug"].(string)
// 			invoice := Invoice{}
// 			invoices := []*Invoice{}
// 			utils.GetJson("http://0.0.0.0:8016/api/v1/invoices/", invoices)
// 			for i, p := range invoices {
// 				if string(slug) == p.Slug {
// 					// invoice = invoices[i]
// 					// Remove from invoice list
// 					invoices = append(invoices[:i], invoices[i+1:]...)
// 				}
// 			}

// 			return invoice, nil
// 		},
// 	}
// }