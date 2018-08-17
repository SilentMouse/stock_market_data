package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/graphql-go/graphql"
	"net/http"
	"github.com/Sirupsen/logrus"
	iex "github.com/SilentMouse/stock_market_data/pkg/go-iexclient"
)

type Ticker struct {
	gorm.Model
	Symbol string
        CompanyName string
        PrimaryExchange string
        LatestPrice float64
	CompanyID uint
}

type Symbol struct{
	gorm.Model
	Name string
	Type string
}

func (self *DataBase) GetTicker(p graphql.ResolveParams) Ticker{
	client := iex.NewClient(&http.Client{})

	quote, err := client.GetBatch(p.Args["symbol"].(string),"quote","1m")

	if err != nil {
		panic(err)
	}

	//for _, quote := range quotes {
	//	logrus.Printf("%v: bid $%.02f (%v shares), ask $%.02f (%v shares) [as of %v]\n",
	//		quote.Symbol, quote.LastSalePrice, quote.LastSaleSize,
	//		quote.AskPrice, quote.AskSize, quote.LastUpdated)
	//}

	logrus.Infoln(quote.Quote)

	ticker := Ticker{Symbol: quote.Quote.Symbol,
		CompanyName: quote.Quote.CompanyName,
		LatestPrice: quote.Quote.LatestPrice,
	}

	return ticker
}
