package controllers

import (
	"github.com/graphql-go/graphql"
	//"net/http"
	"github.com/Sirupsen/logrus"
	"github.com/SilentMouse/stock_market_data/pkg/models"
)

func (c* Controller) ParseSymbols(p graphql.ResolveParams) []models.Symbol{

	logrus.Infoln("parse symbols")

	var symbols []models.Symbol

	//client := iex.NewClient(&http.Client{})

	//symbols, err := client.GetSymbols()

	//if err != nil {
	//	logrus.Error(err)
	//}
	return symbols
}
