package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/Sirupsen/logrus"
	"github.com/SilentMouse/stock_market_data/pkg/models"
)

func (c *Controller) GetTicker(p graphql.ResolveParams) models.Ticker{

	quote, err := c.IexClien.GetBatch(p.Args["symbol"].(string),"quote","1m")

	if err != nil {
		panic(err)
	}

	//for _, quote := range quotes {
	//	logrus.Printf("%v: bid $%.02f (%v shares), ask $%.02f (%v shares) [as of %v]\n",
	//		quote.Symbol, quote.LastSalePrice, quote.LastSaleSize,
	//		quote.AskPrice, quote.AskSize, quote.LastUpdated)
	//}

	logrus.Infoln(quote.Quote)

	ticker := models.Ticker{Symbol: quote.Quote.Symbol,
		CompanyName: quote.Quote.CompanyName,
		LatestPrice: quote.Quote.LatestPrice,
	}

	logrus.Infoln("ticker")

	return ticker
}