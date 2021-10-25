package invoice

import(
	"github.com/graphql-go/graphql"
	"graphql-poc-go/app/utils"
	"fmt"
)

func GetInvoiceQuery() *graphql.Field {
	return &graphql.Field{
		Type:        InvoiceType,
		Description: "Get invoice by proposal_group",
		Args: graphql.FieldConfigArgument{
			"proposal_group": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			proposal_group, ok := p.Args["proposal_group"].(string)
			if ok {
				// Find invoice
				invoice := new(Invoice)
				utils.GetJson("https://staging-api.superyou.co.id/api/v1/invoice/proposal-group/"+proposal_group, invoice)
				fmt.Println(invoice)
				if string(invoice.PolicyGroupNumber) == proposal_group {
					return invoice, nil
				}
			}
			return nil, nil
		},
	}
}

func GetInvoicesQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(InvoiceType),
		Description: "Get invoice list",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			invoices := new([]Invoice)
			utils.GetJson("http://0.0.0.0:8016/api/v1/invoices/", invoices)
			return invoices, nil
		},
	}
}