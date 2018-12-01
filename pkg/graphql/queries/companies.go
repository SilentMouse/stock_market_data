package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/types"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
)

func (fields *GraphqlFields) AddCompanies(c *controllers.Controller) *GraphqlFields{

	fields.Fields["companies"] = &graphql.Field{
		Type:        graphql.NewList(schema_types.CompanyType(c)),
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
			return c.GetCompanies(p), nil
		},
	}


	return fields

}