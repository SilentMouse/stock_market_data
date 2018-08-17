package st_schema

import (
	"github.com/Sirupsen/logrus"
	"github.com/graphql-go/graphql"
	schema_types "github.com/SilentMouse/stock_market_data/pkg/graphql/types"
	models "github.com/SilentMouse/stock_market_data/pkg/models"
)

var (
	StSchema graphql.Schema
)

func InitSchema(db *models.DataBase) {

	logrus.Infoln("Comeon!")

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			/*
			   curl -g 'http://localhost:8080/graphql?query=mutation+_{createSphere(name:"My+new+todo"){name}}'
			*/
			"AddUser": &graphql.Field{
				Type:        schema_types.UserType(db),
				Description: "create new comment",
				Args: graphql.FieldConfigArgument{
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"score_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					user := models.User{}
					return user, nil
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        schema_types.UserType(db),
				Description: "list of users",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the sphere",
						Type:        graphql.String,
					},
					"provider_id": &graphql.ArgumentConfig{
						Description: "id of the sphere",
						Type:        graphql.String,
					},
					"token": &graphql.ArgumentConfig{
						Description: "id of the sphere",
						Type:        graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return db.GetUser(p), nil
				},
			},
			"ticker": &graphql.Field{
				Type:        schema_types.TickerType(db),
				Description: "list of users",
				Args: graphql.FieldConfigArgument{
					"symbol": &graphql.ArgumentConfig{
						Description: "symbol",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return db.GetTicker(p), nil
				},
			},
		},
	})


	StSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: rootMutation,
	})
}