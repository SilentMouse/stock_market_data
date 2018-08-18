package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/Sirupsen/logrus"
	"github.com/SilentMouse/stock_market_data/pkg/models"
)

func (c *Controller) GetSymbols(p graphql.ResolveParams) []models.Symbol{

	logrus.Infoln("symbols")

	var symbols []models.Symbol

	limit := 10
	offset := 0

	if p.Args["limit"] != nil{
		limit = p.Args["limit"].(int)
	}

	if p.Args["offset"] != nil{
		limit = p.Args["offset"].(int)
	}

	c.Models.DB.Raw("select * from symbols limit ? offset ?",
		limit,
		offset,
	).Scan(&symbols)

	return symbols
}
