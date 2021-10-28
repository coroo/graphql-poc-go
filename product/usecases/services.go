package usecases

import (
	"fmt"
	"errors"
	"encoding/json"

	"github.com/graphql-go/graphql"
	"graphql-poc-go/product/entities"
	"graphql-poc-go/app/utils"
	dtoEntities "graphql-poc-go/app/dto/entities"
)

type ProductService interface {
	GetAllProducts() (*[]entities.Product, error)
	GetProduct(params graphql.ResolveParams) (*entities.Product, error)
	SaveProduct(params graphql.ResolveParams) (*entities.Product, error)
	UpdateProduct(params graphql.ResolveParams) (*entities.Product, error)
	DeleteProduct(params graphql.ResolveParams) (*dtoEntities.Delete, error)
}

type productService struct {}

func NewProductService() ProductService {
	return &productService{}
}

func (usecase *productService) GetAllProducts() (*[]entities.Product, error) {
	products := new([]entities.Product)
	err := utils.GetJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (usecase *productService) GetProduct(params graphql.ResolveParams) (*entities.Product, error) {
	slug, ok := params.Args["slug"].(string)
	if ok {
		product := new(entities.Product)
		err := utils.GetJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/slug/"+slug, product)
		if err != nil {
			return nil, err
		}
		if string(product.Slug) == slug {
			return product, nil
		}
	}
	return nil, nil
}

func (usecases *productService) SaveProduct(params graphql.ResolveParams) (*entities.Product, error) {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"slug":    			params.Args["slug"].(string),
		"name":    			params.Args["name"].(string),
		"doc_name":    		params.Args["name"].(string),
		"product_type":    	params.Args["name"].(string),
	})
	//Leverage Go's HTTP Post function to make request

	product := new(entities.Product)
	err := utils.PostJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", product, postBody, "")
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (usecases *productService) UpdateProduct(params graphql.ResolveParams) (*entities.Product, error) {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"slug":    			params.Args["slug"].(string),
		"name":    			params.Args["name"].(string),
		"doc_name":    		params.Args["name"].(string),
		"product_type":    	params.Args["name"].(string),
	})
	//Leverage Go's HTTP Post function to make request

	product := new(entities.Product)
	err := utils.UpdateJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", product, postBody, "")
	if err != nil {
		return nil, err
	}
	if string(product.Id) == "" {
		return nil, errors.New("Error Occured")
	}
	return product, nil
}

func (usecases *productService) DeleteProduct(params graphql.ResolveParams) (*dtoEntities.Delete, error) {
	postBody, _ := json.Marshal(map[string]string{
		"id":    			params.Args["id"].(string),
	})
	response := new(interface{})
	err := utils.DeleteJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"products/", response, postBody, "")
	if err != nil {
		return nil, err
	}
	mapData, _ := utils.InterfaceToMap(response)

	if(mapData["message"] == nil){
		res := &dtoEntities.Delete{
			Message		: "Error",
			Description	: fmt.Sprintf("%v", mapData["detail"]),
		}
		return res, nil
	}
	res := &dtoEntities.Delete{
		Message		: fmt.Sprintf("%v", mapData["message"]),
	}
	return res, nil
}