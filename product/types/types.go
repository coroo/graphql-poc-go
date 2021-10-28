package types

import (
	"github.com/graphql-go/graphql"
)

var ProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
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
			"product_type": &graphql.Field{Type: graphql.String},
			"coverage_period": &graphql.Field{Type: graphql.String},
			"available_claim_methods": &graphql.Field{Type: graphql.String},
			"covid_coverage": &graphql.Field{Type: graphql.Boolean},
			"start_age_from": &graphql.Field{Type: graphql.Int},
			"start_premium_from": &graphql.Field{Type: graphql.Float},
			"category": &graphql.Field{Type: ProductCategoryType},
			"insurance_type": &graphql.Field{Type: ProductInsuranceTypeType},
			"riders": &graphql.Field{Type: graphql.NewList(ProductRiderType)},
			"benefit_groups": &graphql.Field{Type: graphql.NewList(ProductBenefitGroupType)},
			"tnc": &graphql.Field{Type: graphql.NewList(ProductTncType)},
			"faq": &graphql.Field{Type: graphql.NewList(ProductFaqType)},
			"not_coverage": &graphql.Field{Type: graphql.NewList(ProductNotCoverageType)},
		},
	},
)

// ProductBenefitGroupType is the GraphQL schema for the Product type.
var ProductBenefitGroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductBenefitGroup",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
		"benefits": &graphql.Field{Type: graphql.NewList(ProductBenefitType)},
	},
})

// ProductBenefitType is the GraphQL schema for the Product type.
var ProductBenefitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductBenefit",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"icon_etc": &graphql.Field{Type: graphql.String},
		"product_id": &graphql.Field{Type: graphql.String},
		"product_benefit_group_id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"tooltip_text_description": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
	},
})

// ProductTncType is the GraphQL schema for the Product type.
var ProductTncType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductTnc",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"icon_etc": &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"product_id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
	},
})

// ProductFaqType is the GraphQL schema for the Product type.
var ProductFaqType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductFaq",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"title": &graphql.Field{Type: graphql.String},
		"body": &graphql.Field{Type: graphql.String},
		"product_id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
	},
})

// ProductNotCoverageType is the GraphQL schema for the Product type.
var ProductNotCoverageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductNotCoverage",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"icon_etc": &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"product_id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
	},
})

// ProductRiderType is the GraphQL schema for the Product type.
var ProductRiderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductRider",
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

// ProductCategoryType is the GraphQL schema for the Product type.
var ProductCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductCategory",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// ProductInsuranceTypeType is the GraphQL schema for the Product type.
var ProductInsuranceTypeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductInsuranceType",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})