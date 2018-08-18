package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Ticker struct {
	gorm.Model
	Symbol string
        CompanyName string
        PrimaryExchange string
        LatestPrice float64
	CompanyID uint
}