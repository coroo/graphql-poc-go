package main

import (
	"net/http"
	"log"
	"context"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/graphql"
	"graphql-poc-go/app/routes"
	"graphql-poc-go/app/utils"
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
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			ctx = context.WithValue(r.Context(), "header", r.Header)
			rootObject := make(map[string]interface{}, len(r.Header))
			for k, v := range r.Header {
				rootObject[k] = v
			}
			return rootObject
		},
	})

	http.Handle("/graphql", h)
	log.Println("Server ready at "+utils.EnvVariable("MAIN_SCHEMES")+"://"+utils.EnvVariable("MAIN_URL")+"/"+utils.EnvVariable("MAIN_PORT")+"/graphql")
	http.ListenAndServe(":"+utils.EnvVariable("MAIN_PORT")+"", nil)
}