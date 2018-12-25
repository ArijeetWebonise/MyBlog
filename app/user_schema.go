package app

import (
	"github.com/graphql-go/graphql"
)

func (app *App) GetUserSchema() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "User Info",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"first_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
