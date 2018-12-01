package schema_types


import (

	"github.com/graphql-go/graphql"
	models "github.com/SilentMouse/stock_market_data/pkg/models"
	"github.com/SilentMouse/stock_market_data/pkg/utils"
	"github.com/SilentMouse/stock_market_data/pkg/controllers"
	"strconv"
)


func CompanyType(c *controllers.Controller) *graphql.Object {

	var TypeVar *graphql.Object

	hashgr := utils.RandToken5()

	TypeVar = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Company_" + hashgr,
		Description: "company",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:         graphql.String,
				Description: "id company",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:         graphql.String,
				Description: "name",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.Name, nil
					}
					return nil, nil
				},
			},
			"industry": &graphql.Field{
				Type:         graphql.String,
				Description: "industry",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.Industry, nil
					}
					return nil, nil
				},
			},
			"site": &graphql.Field{
				Type:         graphql.String,
				Description: "site",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.Site, nil
					}
					return nil, nil
				},
			},
			"description": &graphql.Field{
				Type:         graphql.String,
				Description: "description",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.Description, nil
					}
					return nil, nil
				},
			},
			"ceo": &graphql.Field{
				Type:         graphql.String,
				Description: "ceo",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.CEO, nil
					}
					return nil, nil
				},
			},
			"IssueType": &graphql.Field{
				Type:         graphql.String,
				Description: "itype",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.IssueType, nil
					}
					return nil, nil
				},
			},
			"sector": &graphql.Field{
				Type:         graphql.String,
				Description: "sector",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return ent.Sector, nil
					}
					return nil, nil
				},
			},
			"tags": &graphql.Field{
				Type:         graphql.NewList(graphql.String),
				Description: "tags",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						return c.GetCompanyTags(ent.ID,p), nil
					}
					return nil, nil
				},
			},
			"symbol": &graphql.Field{
				Type:         graphql.String,
				Description: "tags",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if ent, ok := p.Source.(models.Company); ok {
						ent_id := strconv.FormatUint(uint64(ent.SymbolID),10)
						p.Args["id"] = ent_id
						return c.GetSymbol(p).Abbr, nil
					}
					return nil, nil
				},
			},



		},
	})

	return TypeVar
}