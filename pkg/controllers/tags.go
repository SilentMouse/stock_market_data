package controllers

import (
	"github.com/graphql-go/graphql"
	"github.com/SilentMouse/stock_market_data/pkg/models"
)

func (c *Controller) GetCompanyTags(company_id uint,p graphql.ResolveParams) []string{

	var tags []models.Tag

	c.Models.DB.Raw("select * from tags join companies_tags " +
		"on companies_tags.tag_id = tags.id " +
		"where companies_tags.company_id = ?", company_id).Scan(&tags)

	var str_tags []string

	for _, t := range tags{
		str_tags = append(str_tags,t.Name)
	}

	return str_tags
}
