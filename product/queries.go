package product

import(
	"github.com/graphql-go/graphql"
	"graphql-poc-go/product/types"
	"graphql-poc-go/product/usecases"
)
type ProductQuery interface {
	GetProductsQuery() *graphql.Field
	GetProductQuery() *graphql.Field
}

type productQuery struct {
	usecases		usecases.ProductService
}

func NewProductQuery(usecase usecases.ProductService) ProductQuery {
	return &productQuery{
		usecases: usecase,
	}
}

func (queries *productQuery) GetProductsQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.ProductType),
		Description: "Get product list",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			products := usecases.ProductService.GetAllProducts(queries.usecases)
			return products, nil
		},
	}
}

func (queries *productQuery) GetProductQuery() *graphql.Field {
	return &graphql.Field{
		Type:        types.ProductType,
		Description: "Get product by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			product := usecases.ProductService.GetProduct(queries.usecases, params)
			return product, nil
		},
	}
}