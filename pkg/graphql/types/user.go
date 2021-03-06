package schema_types


import (

	"github.com/graphql-go/graphql"
	models "github.com/SilentMouse/stock_market_data/pkg/models"
	"github.com/SilentMouse/stock_market_data/pkg/utils"
)


func UserType() *graphql.Object {

	var userTypeVar *graphql.Object

	hashgr := utils.RandToken5()

	userTypeVar = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User_" + hashgr,
		Description: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the User",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.Name, nil
					}
					return nil, nil
				},
			},
			"main_image": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the User",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.MainImage, nil
					}
					return nil, nil
				},
			},
			"first_name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the User",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.FirstName, nil
					}
					return nil, nil
				},
			},
			"second_name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the User",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.SecondName, nil
					}
					return nil, nil
				},
			},
			"phone": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.Phone, nil
					}
					return nil, nil
				},
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.Email, nil
					}
					return nil, nil
				},
			},
			"token": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.Token, nil
					}
					return nil, nil
				},
			},
			"created_at": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the user",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(models.User); ok {
						return user.CreatedAt, nil
					}
					return nil, nil
				},
			},
		},
	})

	return userTypeVar
}


