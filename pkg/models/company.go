package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	SymbolID    uint  `json:"-"`
	Name        string `json:"companyName"`
	Industry    string `json:"industry"`
	Site        string `json:"website"`
	Description string `json:"description"`
	CEO         string `json:"ceo"`
	IssueType   string `json:"issueType"`
	Sector      string `json:"sector" gorm:"-"`
	SectorID    uint   `json:"-"`
	Tags        []string  `json:"tags" gorm:"-"`
}

type CompaniesTags struct {
	gorm.Model
	CompanyID uint
	TagID     uint
}
