package schema_types


import (

	"github.com/graphql-go/graphql"
	models "github.com/SilentMouse/stock_market_data/pkg/models"
	"github.com/SilentMouse/stock_market_data/pkg/utils"
)


func SymbolType() *graphql.Object {

	var TypeVar *graphql.Object

	hashgr := utils.RandToken5()

	TypeVar = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Symbol_" + hashgr,
		Description: "symbol",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:         graphql.String,
				Description: "system id symbol",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if symbol, ok := p.Source.(models.Symbol); ok {
						return symbol.ID, nil
					}
					return nil, nil
				},
			},
			"abbr": &graphql.Field{
				Type:        graphql.String,
				Description: "abbr symbol",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if symbol, ok := p.Source.(models.Symbol); ok {
						return symbol.Abbr, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "abbr symbol",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if symbol, ok := p.Source.(models.Symbol); ok {
						return symbol.Name, nil
					}
					return nil, nil
				},
			},
			"date": &graphql.Field{
				Type:        graphql.String,
				Description: "date",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if symbol, ok := p.Source.(models.Symbol); ok {
						return symbol.Date, nil
					}
					return nil, nil
				},
			},
			"is_enable": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "is_enable",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if symbol, ok := p.Source.(models.Symbol); ok {
						return symbol.IsEnabled, nil
					}
					return nil, nil
				},
			},
			"type": &graphql.Field{
				Type:        graphql.String,
				Description: "abbr symbol",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if symbol, ok := p.Source.(models.Symbol); ok {
						return symbol.Type, nil
					}
					return nil, nil
				},
			},
			"IexId": &graphql.Field{
				Type:        graphql.String,
				Description: "iexID",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if symbol, ok := p.Source.(models.Symbol); ok {
						return symbol.IexID, nil
					}
					return nil, nil
				},
			},

		},
	})

	return TypeVar
}
