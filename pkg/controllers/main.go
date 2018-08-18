package controllers

import (
	"github.com/SilentMouse/stock_market_data/pkg/models"
	"github.com/SilentMouse/stock_market_data/pkg/go-iexclient"
	"net/http"
)

type Controller struct{
	Models *models.DataBase
	IexClien *go_iexclient.Client
}


func InitController(db *models.DataBase) Controller{
	var self Controller

	self.Models = db
	self.IexClien = go_iexclient.NewClient(&http.Client{})

	return self
}
