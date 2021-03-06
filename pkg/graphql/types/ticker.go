package schema_types


import (

	"github.com/graphql-go/graphql"
	models "github.com/SilentMouse/stock_market_data/pkg/models"
	"github.com/SilentMouse/stock_market_data/pkg/utils"
)


func TickerType() *graphql.Object {

	var TypeVar *graphql.Object

	hashgr := utils.RandToken5()

	TypeVar = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Ticker_" + hashgr ,
		Description: "ticker",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ticker, ok := p.Source.(models.Ticker); ok {
						return ticker.ID, nil
					}
					return nil, nil
				},
			},
			"symbol": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the User",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ticker, ok := p.Source.(models.Ticker); ok {
						return ticker.Symbol, nil
					}
					return nil, nil
				},
			},
			"company_name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the User",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ticker, ok := p.Source.(models.Ticker); ok {
						return ticker.CompanyName, nil
					}
					return nil, nil
				},
			},
			"latest_price": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the User",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ticker, ok := p.Source.(models.Ticker); ok {
						return ticker.LatestPrice, nil
					}
					return nil, nil
				},
			},
		},
	})

	return TypeVar
}


