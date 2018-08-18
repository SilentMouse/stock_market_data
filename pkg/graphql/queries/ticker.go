package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/types"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
)

func (fields *GraphqlFields) AddTicker(c *controllers.Controller) *GraphqlFields{

	fields.Fields["ticker"] = &graphql.Field{
		Type:        schema_types.TickerType(),
		Description: "ticker",
		Args: graphql.FieldConfigArgument{
			"symbol": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return c.GetTicker(p), nil
		},
	}


	return fields

}