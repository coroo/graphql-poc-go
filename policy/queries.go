package policy

import(
	"github.com/graphql-go/graphql"
	"graphql-poc-go/policy/types"
	"graphql-poc-go/policy/usecases"
)
type PolicyQuery interface {
	GetPoliciesQuery() *graphql.Field
	GetPolicyQuery() *graphql.Field
}

type policyQuery struct {
	usecases		usecases.PolicyService
}

func NewPolicyQuery(usecase usecases.PolicyService) PolicyQuery {
	return &policyQuery{
		usecases: usecase,
	}
}

func (queries *policyQuery) GetPoliciesQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(types.PolicyType),
		Description: "Get policy list",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			policies := usecases.PolicyService.GetAllPolicies(queries.usecases)
			return policies, nil
		},
	}
}

func (queries *policyQuery) GetPolicyQuery() *graphql.Field {
	return &graphql.Field{
		Type:        types.PolicyType,
		Description: "Get policy by slug",
		Args: graphql.FieldConfigArgument{
			"slug": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			policy := usecases.PolicyService.GetPolicy(queries.usecases, params)
			return policy, nil
		},
	}
}