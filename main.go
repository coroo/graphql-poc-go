package main

import (
	"net/http"
	"log"
	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
	// "graphql-poc-go/product"
	"graphql-poc-go/app/routes"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: routes.GetRootQueryFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: routes.GetRootMutationFields(),
		}),
	},
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Println("Server ready at http://localhost/8080/graphql")
	http.ListenAndServe(":8080", nil)
}