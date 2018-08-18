package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Symbol struct {
	gorm.Model
	Abbr      string `json: "symbol"`
	Name      string `json: "name"`
	Date      time.Time `json: "date`
	IsEnabled bool `json: "isEnabled"`
	Type      string `json: "isEnabled"`
	IexId     string `json: "iexID"`
}
