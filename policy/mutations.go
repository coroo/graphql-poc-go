package policy

import (
	"github.com/graphql-go/graphql"
	"graphql-poc-go/app/utils"
	"graphql-poc-go/policy/types"
	"graphql-poc-go/policy/entities"
	"graphql-poc-go/policy/usecases"
	dtoTypes "graphql-poc-go/app/dto/types"
)
type PolicyMutation interface {
	CreatePolicyMutation() *graphql.Field
	UpdatePolicyMutation() *graphql.Field
	DeletePolicyMutation() *graphql.Field
}

type policyMutation struct {
	usecases		usecases.PolicyService
}

func NewPolicyMutation(usecase usecases.PolicyService) PolicyMutation {
	return &policyMutation{
		usecases: usecase,
	}
}

func (mutations *policyMutation) CreatePolicyMutation() *graphql.Field {
	return &graphql.Field{
		Type:        types.PolicyType,
		Description: "Create new policy",
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
			policy := usecases.PolicyService.SavePolicy(mutations.usecases, params)
			return policy, nil
		},
	}
}

func (mutations *policyMutation) UpdatePolicyMutation() *graphql.Field {
	return &graphql.Field{
		Type:        types.PolicyType,
		Description: "Update policy by slug",
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
			policy := entities.Policy{}
			policies := []entities.Policy{}
			utils.GetJson("http://0.0.0.0:8016/api/v1/policies/", policies)
			for i, p := range policies {
				if string(slug) == p.Slug {
					if nameOk {
						policies[i].Name = name
					}
					if descriptionOk {
						policies[i].Description = description
					}
					if summaryOk {
						policies[i].Summary = summary
					}
					if priceOk {
						policies[i].StartPremiumFrom = price
					}
					policy = policies[i]
					break
				}
			}
			return policy, nil
		},
	}
}

func (mutations *policyMutation) DeletePolicyMutation() *graphql.Field {
	return &graphql.Field{
		Type:        dtoTypes.DeleteType,
		Description: "Delete policy by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			response := usecases.PolicyService.DeletePolicy(mutations.usecases, params)
			return response, nil
		},
	}
}