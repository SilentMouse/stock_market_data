package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/types"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
)

func (fields *GraphqlFields) AddSearch(c *controllers.Controller) *GraphqlFields{

	fields.Fields["search"] = &graphql.Field{
		Type:        graphql.NewList(schema_types.CompanyType(c)),
		Description: "search",
		Args: graphql.FieldConfigArgument{
			"query": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return c.Search(p)
		},
	}


	return fields

}