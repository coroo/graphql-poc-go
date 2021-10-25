package invoice

import (
	"github.com/graphql-go/graphql"
)

var InvoiceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Invoice",
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
			"invoice_type": &graphql.Field{Type: graphql.String},
			"coverage_period": &graphql.Field{Type: graphql.String},
			"available_claim_methods": &graphql.Field{Type: graphql.String},
			"covid_coverage": &graphql.Field{Type: graphql.Boolean},
			"start_age_from": &graphql.Field{Type: graphql.Int},
			"start_premium_from": &graphql.Field{Type: graphql.Float},
			"category": &graphql.Field{Type: InvoiceCategoryType},
			"insurance_type": &graphql.Field{Type: InvoiceInsuranceTypeType},
			"riders": &graphql.Field{Type: graphql.NewList(InvoiceRiderType)},
			"benefit_groups": &graphql.Field{Type: graphql.NewList(InvoiceBenefitGroupType)},
		},
	},
)

// InvoiceBenefitGroupType is the GraphQL schema for the Invoice type.
var InvoiceBenefitGroupType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceBenefitGroup",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
		"benefits": &graphql.Field{Type: graphql.NewList(InvoiceBenefitType)},
	},
})

// InvoiceBenefitType is the GraphQL schema for the Invoice type.
var InvoiceBenefitType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceBenefit",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"icon_svg": &graphql.Field{Type: graphql.String},
		"icon_etc": &graphql.Field{Type: graphql.String},
		"invoice_id": &graphql.Field{Type: graphql.String},
		"invoice_benefit_group_id": &graphql.Field{Type: graphql.String},
		"tooltip_text": &graphql.Field{Type: graphql.String},
		"tooltip_text_description": &graphql.Field{Type: graphql.String},
		"order": &graphql.Field{Type: graphql.Int},
	},
})

// InvoiceRiderType is the GraphQL schema for the Invoice type.
var InvoiceRiderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceRider",
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

// InvoiceCategoryType is the GraphQL schema for the Invoice type.
var InvoiceCategoryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceCategory",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// InvoiceInsuranceTypeType is the GraphQL schema for the Invoice type.
var InvoiceInsuranceTypeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "InvoiceInsuranceType",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})