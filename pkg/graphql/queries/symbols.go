package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/types"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
)

func (fields *GraphqlFields) AddSymbols(c *controllers.Controller) *GraphqlFields{

	fields.Fields["symbols"] = &graphql.Field{
		Type:        graphql.NewList(schema_types.SymbolType()),
		Description: "Symbol",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return c.GetSymbols(p), nil
		},
	}


	return fields

}