package app

import (
	"github.com/graphql-go/graphql"
)

func (app *App) GetSchema() (graphql.Schema, error) {
	fields := graphql.Fields{
		"ping": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "pong", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	return graphql.NewSchema(schemaConfig)
}
