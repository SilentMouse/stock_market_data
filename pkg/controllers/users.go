package controllers

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/models"
)



func (c *Controller) GetUser(p graphql.ResolveParams) models.User{

	var user models.User
	return user
}



func (c *Controller) GetUsers(p graphql.ResolveParams) []models.User{

	var users []models.User
	return users
}
