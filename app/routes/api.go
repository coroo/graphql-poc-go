package routes

import (
	"github.com/graphql-go/graphql"
	"graphql-poc-go/product"
	"graphql-poc-go/product/usecases"
	"graphql-poc-go/invoice"
)

var (
	productService    usecases.ProductService = usecases.NewProductService()
)

// GetRootQueryFields returns all the available queries.
func GetRootQueryFields() graphql.Fields {
	// return product.NewProductQuery(productService).GetProductsQuery()
	return graphql.Fields{
		"product": product.NewProductQuery(productService).GetProductQuery(),
		"products": product.NewProductQuery(productService).GetProductsQuery(),

		"invoice": invoice.GetInvoiceQuery(),
		"invoices": invoice.GetInvoicesQuery(),
	}
}

// GetRootMutationFields returns all the available mutations.
func GetRootMutationFields() graphql.Fields {
	return graphql.Fields{
		"createProduct": product.NewProductMutation(productService).CreateProductMutation(),
		"updateProduct": product.NewProductMutation(productService).UpdateProductMutation(),
		"deleteProduct": product.NewProductMutation(productService).DeleteProductMutation(),

		// "createInvoice": invoice.CreateInvoiceMutation(),
		// "updateInvoice": invoice.UpdateInvoiceMutation(),
		// "deleteInvoice": invoice.DeleteInvoiceMutation(),
	}
}