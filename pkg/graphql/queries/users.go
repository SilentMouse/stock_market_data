package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/types"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
)

func (fields *GraphqlFields) AddUsers(c *controllers.Controller) *GraphqlFields{

	fields.Fields["users"] = &graphql.Field{
		Type:        graphql.NewList(schema_types.SymbolType()),
		Description: "users",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return c.GetUsers(p), nil
		},
	}

	fields.Fields["user"] = &graphql.Field{
		Type:        schema_types.UserType(),
		Description: "user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return c.GetUser(p), nil
		},
	}


	return fields

}