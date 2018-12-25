package app

import (
	"errors"

	"github.com/ArijeetBaruah/MyBlog/app/models"
	"github.com/graphql-go/graphql"
	"golang.org/x/crypto/bcrypt"
)

func (app *App) GetSchema() (graphql.Schema, error) {
	userSchema := app.GetUserSchema()

	fields := graphql.Fields{
		"ping": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "pong", nil
			},
		},
		"user": &graphql.Field{
			Type: userSchema,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type:         graphql.String,
					Description:  "User ID",
					DefaultValue: "",
				},
				"password": &graphql.ArgumentConfig{
					Type:         graphql.String,
					Description:  "User Password",
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				password := p.Args["password"].(string)
				user, err := models.UserByEmail(app.DB, email)
				if err != nil {
					return nil, err
				}
				if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
					return nil, errors.New("Incorrect Username or Password")
				}

				userRet := make(map[string]interface{})
				userRet["id"] = user.ID
				userRet["email"] = user.Email
				if user.FirstName.Valid {
					userRet["first_name"] = user.FirstName.String
				}
				if user.LastName.Valid {
					userRet["last_name"] = user.LastName.String
				}
				return userRet, nil
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
