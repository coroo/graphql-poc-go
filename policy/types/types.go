package types

import (
	"github.com/graphql-go/graphql"
)

var PolicyType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Policy",
		Fields: graphql.Fields{
			"id": &graphql.Field{Type: graphql.String},
			"slug": &graphql.Field{Type: graphql.String},
			"name": &graphql.Field{Type: graphql.String},
			"doc_name": &graphql.Field{Type: graphql.String},
			"parent_id": &graphql.Field{Type: graphql.String},
			"is_active": &graphql.Field{Type: graphql.Boolean},
			"featured": &graphql.Field{Type: graphql.Boolean},
			"bundling_with_rider": &graphql.Field{Type: graphql.Boolean},
			"subheading": &graphql.Field{Type: graphql.String},
			"summary": &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"icon_svg": &graphql.Field{Type: graphql.String},
			"icon_etc": &graphql.Field{Type: graphql.String},
			"rip_link": &graphql.Field{Type: graphql.String},
			"policy_type": &graphql.Field{Type: graphql.String},
			"coverage_period": &graphql.Field{Type: graphql.String},
			"available_claim_methods": &graphql.Field{Type: graphql.String},
			"covid_coverage": &graphql.Field{Type: graphql.Boolean},
			"start_age_from": &graphql.Field{Type: graphql.Int},
			"start_premium_from": &graphql.Field{Type: graphql.Float},
			"category": &graphql.Field{Type: PolicyCategoryType},
			"insurance_type": &graphql.Field{Type: PolicyInsuranceTypeType},
			"riders": &graphql.Field{Type: graphql.NewList(PolicyRiderType)},
			"benefit_groups": &graphql.Field{Type: graphql.NewList(PolicyBenefitGroupType)},
		},
	},
)

// PolicyBenefitGroupType is the GraphQL schema for the Policy type.
var PolicyBenefitGroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PolicyBenefitGroup",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
		"benefits": &graphql.Field{Type: graphql.NewList(PolicyBenefitType)},
	},
})

// PolicyBenefitType is the GraphQL schema for the Policy type.
var PolicyBenefitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PolicyBenefit",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"icon_etc": &graphql.Field{Type: graphql.String},
		"policy_id": &graphql.Field{Type: graphql.String},
		"policy_benefit_group_id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"tooltip_text_description": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
	},
})

// PolicyRiderType is the GraphQL schema for the Policy type.
var PolicyRiderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PolicyRider",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"slug": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"is_active": &graphql.Field{Type: graphql.Boolean},
		"parent_id": &graphql.Field{Type: graphql.String},
		"summary": &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"coverage_period": &graphql.Field{Type: graphql.String},
	},
})

// PolicyCategoryType is the GraphQL schema for the Policy type.
var PolicyCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PolicyCategory",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// PolicyInsuranceTypeType is the GraphQL schema for the Policy type.
var PolicyInsuranceTypeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "PolicyInsuranceType",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})