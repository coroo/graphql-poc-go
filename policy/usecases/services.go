package usecases

import (
	"fmt"
	"log"
	"encoding/json"

	"github.com/graphql-go/graphql"
	"graphql-poc-go/policy/entities"
	"graphql-poc-go/app/utils"
	dtoEntities "graphql-poc-go/app/dto/entities"
)

type PolicyService interface {
	GetAllPolicies() *[]entities.Policy
	GetPolicy(params graphql.ResolveParams) *entities.Policy
	SavePolicy(params graphql.ResolveParams) *entities.Policy
	UpdatePolicy(params graphql.ResolveParams) *entities.Policy
	DeletePolicy(params graphql.ResolveParams) *dtoEntities.Delete
}

type policyService struct {}

func NewPolicyService() PolicyService {
	return &policyService{}
}

func (usecase *policyService) GetAllPolicies() *[]entities.Policy {
	policies := new([]entities.Policy)
	utils.GetJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"policies/", policies)
	return policies
}

func (usecase *policyService) GetPolicy(params graphql.ResolveParams) *entities.Policy {
	slug, ok := params.Args["slug"].(string)
	if ok {
		policy := new(entities.Policy)
		utils.GetJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"policies/slug/"+slug, policy)
		if string(policy.Slug) == slug {
			return policy
		}
	}
	return nil
}

func (usecases *policyService) SavePolicy(params graphql.ResolveParams) *entities.Policy {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"slug":    			params.Args["slug"].(string),
		"name":    			params.Args["name"].(string),
		"doc_name":    		params.Args["name"].(string),
		"policy_type":    	params.Args["name"].(string),
	})
	//Leverage Go's HTTP Post function to make request

	policy := new(entities.Policy)
	utils.PostJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"policies/", policy, postBody, "")
	if string(policy.Id) == "" {
		log.Println("Error Occured")
	}
	return policy
}

func (usecases *policyService) UpdatePolicy(params graphql.ResolveParams) *entities.Policy {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"slug":    			params.Args["slug"].(string),
		"name":    			params.Args["name"].(string),
		"doc_name":    		params.Args["name"].(string),
		"policy_type":    	params.Args["name"].(string),
	})
	//Leverage Go's HTTP Post function to make request

	policy := new(entities.Policy)
	utils.UpdateJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"policies/", policy, postBody, "")
	if string(policy.Id) == "" {
		log.Println("Error Occured")
	}
	return policy
}

func (usecases *policyService) DeletePolicy(params graphql.ResolveParams) *dtoEntities.Delete {
	postBody, _ := json.Marshal(map[string]string{
		"id":    			params.Args["id"].(string),
	})
	response := new(interface{})
	utils.DeleteJson(utils.EnvVariable("MICRO_PRODUCT_LINK")+"policies/", response, postBody, "")
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