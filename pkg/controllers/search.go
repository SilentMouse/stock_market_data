package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/models"
	"github.com/Sirupsen/logrus"
)

func (c *Controller) Search(p graphql.ResolveParams) ([]models.Company,error){

	var companies []models.Company

	logrus.Infoln(p.Args["query"].(string))

	c.Models.Search.Raw("SELECT * FROM test1 WHERE MATCH(?)", p.Args["query"].(string)).Scan(&companies)

	return companies, nil
}
