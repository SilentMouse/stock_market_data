package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/Sirupsen/logrus"
	"github.com/SilentMouse/stock_market_data/pkg/models"
)

func (c *Controller) GetCompanies(p graphql.ResolveParams) []models.Company{

	logrus.Infoln("companies")

	var companies []models.Company

	limit := 10
	offset := 0

	if p.Args["limit"] != nil{
		limit = p.Args["limit"].(int)
	}

	if p.Args["offset"] != nil{
		offset = p.Args["offset"].(int)
	}

	c.Models.DB.Raw("select *,sectors.name sector from companies join sectors " +
		"on sectors.id = companies.sector_id limit ? offset ?",
		limit,
		offset,
	).Scan(&companies)

	return companies
}
