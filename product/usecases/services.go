package usecases

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/graphql-go/graphql"
	"graphql-poc-go/product/entities"
	"graphql-poc-go/app/utils"
	dtoEntities "graphql-poc-go/app/dto/entities"
)

type ProductService interface {
	GetAllProducts() *[]entities.Product
	GetProduct(params graphql.ResolveParams) *entities.Product
	SaveProduct(params graphql.ResolveParams) *entities.Product
	UpdateProduct(params graphql.ResolveParams) *entities.Product
	DeleteProduct(params graphql.ResolveParams) *dtoEntities.Delete
}

type productService struct {}

func NewProductService() ProductService {
	return &productService{}
}

func (usecase *productService) GetAllProducts() *[]entities.Product {
	products := new([]entities.Product)
	utils.GetJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", products)
	return products
}

func (usecase *productService) GetProduct(params graphql.ResolveParams) *entities.Product {
	slug, ok := params.Args["slug"].(string)
	if ok {
		product := new(entities.Product)
		utils.GetJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/slug/"+slug, product)
		if string(product.Slug) == slug {
			return product
		}
	}
	return nil
}

func (usecases *productService) SaveProduct(params graphql.ResolveParams) *entities.Product {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"slug":    			params.Args["slug"].(string),
		"name":    			params.Args["name"].(string),
		"doc_name":    		params.Args["name"].(string),
		"product_type":    	params.Args["name"].(string),
	})
	//Leverage Go's HTTP Post function to make request

	product := new(entities.Product)
	utils.PostJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", product, postBody, "")
	if string(product.Id) == "" {
		log.Println("Error Occured")
	}
	return product
}

func (usecases *productService) UpdateProduct(params graphql.ResolveParams) *entities.Product {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"slug":    			params.Args["slug"].(string),
		"name":    			params.Args["name"].(string),
		"doc_name":    		params.Args["name"].(string),
		"product_type":    	params.Args["name"].(string),
	})
	//Leverage Go's HTTP Post function to make request

	product := new(entities.Product)
	utils.UpdateJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", product, postBody, "")
	if string(product.Id) == "" {
		log.Println("Error Occured")
	}
	return product
}

func (usecases *productService) DeleteProduct(params graphql.ResolveParams) *dtoEntities.Delete {
	postBody, _ := json.Marshal(map[string]string{
		"id":    			params.Args["id"].(string),
	})
	response := new(interface{})
	utils.DeleteJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", response, postBody, "")
	mapData, _ := utils.InterfaceToMap(response)

	if(mapData["message"] == nil){
		res := &dtoEntities.Delete{
			Message		: "Error",
			Description	: fmt.Sprintf("%v", mapData["detail"]),
		}
		return res
	}
	res := &dtoEntities.Delete{
		Message		: fmt.Sprintf("%v", mapData["message"]),
	}
	return res
}