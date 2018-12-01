package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/models"
	"github.com/Sirupsen/logrus"
)

func (c *Controller) GetSymbols(p graphql.ResolveParams) []models.Symbol{

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


func (c *Controller) GetSymbol(p graphql.ResolveParams) models.Symbol{

	var symbol models.Symbol

	c.Models.DB.First(&symbol,p.Args["id"].(string))

	logrus.Infoln("P", symbol)

	return symbol
}
