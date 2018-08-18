package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/types"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
)

func (fields *GraphqlFields) AddParser(c *controllers.Controller) *GraphqlFields{

	fields.Fields["parseSymbols"] = &graphql.Field{
		Type:        graphql.NewList(schema_types.SymbolType()),
		Description: "parse symbols",
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			symbols := c.ParseSymbols(params)
			return symbols, nil
		},
	}

	return fields

}