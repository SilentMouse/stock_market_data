package controllers

import (
	"github.com/graphql-go/graphql"
	//"net/http"
	"github.com/Sirupsen/logrus"
	"github.com/SilentMouse/stock_market_data/pkg/models"

)

func (c* Controller) ParseSymbols(p graphql.ResolveParams) *[]models.Symbol{

	symbols,err := c.IexClien.GetSymbols()

	for _, v := range *symbols{
		c.Models.DB.Create(&v)
	}

	//client := iex.NewClient(&http.Client{})

	//symbols, err := client.GetSymbols()

	if err != nil {
		logrus.Error("ee", err)
	}

	return symbols


}


func (c* Controller) ParseCompanies(p graphql.ResolveParams) {

	var symbols []models.Symbol

	c.Models.DB.Raw("select * from symbols").Scan(&symbols)

	for _, v := range symbols{
		company, err := c.IexClien.GetCompanies(v.Abbr)

		if err != nil{
			logrus.Error(err)
		}

		company.SymbolID = v.ID

		c.Models.DB.Create(&company)

		var sector models.Sector

		c.Models.DB.Raw("select * from sectors where name = ?", company.Sector).Scan(&sector)

		if sector.ID == 0{
			sector.Name = company.Sector
			c.Models.DB.Create(&sector)
		}

		company.SectorID = sector.ID

		c.Models.DB.Save(&company)

		for _, t := range company.Tags{

			var tag models.Tag
			c.Models.DB.Raw("select * from tags where name = ?", t).Scan(&tag)

			if tag.ID == 0{
				tag.Name = t
				c.Models.DB.Create(&tag)
			}

			var ct models.CompaniesTags

			ct.CompanyID = company.ID
			ct.TagID = tag.ID

			c.Models.DB.Create(&ct)

		}

		logrus.Infoln(company)

	}

}

func (c* Controller) AddCompaniesToSphinx(p graphql.ResolveParams) {

	var companies []models.Company

	c.Models.DB.Raw("select * from companies").Scan(&companies)

	for _, v := range companies{
		c.Models.MySql.Save(&v)
	}

}
