package st_schema

import (
	"github.com/Sirupsen/logrus"
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/queries"
	"github.com/SilentMouse/stock_market_data/pkg/graphql/mutations"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
)

var (
	StSchema graphql.Schema
)

func InitSchema(c *controllers.Controller) {

	logrus.Infoln("Comeon!")

	mutation_fields := mutations.GraphqlFields{Fields: graphql.Fields{}}

	mutation_fields.AddParser(c)

	var rootMutation = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: mutation_fields.Fields,
	})

	fields := queries.GraphqlFields{Fields: graphql.Fields{}}

	fields.AddSymbols(c).AddTicker(c).AddUsers(c).AddCompanies(c).AddSearch(c)

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: fields.Fields,
	})

	StSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: rootMutation,
	})
}