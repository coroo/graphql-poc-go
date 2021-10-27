package types

import (
	"github.com/graphql-go/graphql"
)

var DeleteType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Delete",
		Fields: graphql.Fields{
			"message": &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
		},
	},
)