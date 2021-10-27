package routes

import (
	"github.com/graphql-go/graphql"
	"graphql-poc-go/product"
	productUsecases "graphql-poc-go/product/usecases"
	"graphql-poc-go/policy"
	policyUsecases "graphql-poc-go/policy/usecases"
)

var (
	productService    productUsecases.ProductService = productUsecases.NewProductService()
	policyService    policyUsecases.PolicyService = policyUsecases.NewPolicyService()
)

// GetRootQueryFields returns all the available queries.
func GetRootQueryFields() graphql.Fields {
	// return product.NewProductQuery(productService).GetProductsQuery()
	return graphql.Fields{
		"product": product.NewProductQuery(productService).GetProductQuery(),
		"products": product.NewProductQuery(productService).GetProductsQuery(),

		"policy": policy.NewPolicyQuery(policyService).GetPolicyQuery(),
		"policies": policy.NewPolicyQuery(policyService).GetPoliciesQuery(),
	}
}

// GetRootMutationFields returns all the available mutations.
func GetRootMutationFields() graphql.Fields {
	return graphql.Fields{
		"createProduct": product.NewProductMutation(productService).CreateProductMutation(),
		"updateProduct": product.NewProductMutation(productService).UpdateProductMutation(),
		"deleteProduct": product.NewProductMutation(productService).DeleteProductMutation(),

		"createPolicy": policy.NewPolicyMutation(policyService).CreatePolicyMutation(),
		"updatePolicy": policy.NewPolicyMutation(policyService).UpdatePolicyMutation(),
		"deletePolicy": policy.NewPolicyMutation(policyService).DeletePolicyMutation(),
	}
}