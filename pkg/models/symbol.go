package models

import (
	"github.com/jinzhu/gorm"
)

type Symbol struct {
	gorm.Model
	Abbr      string `json:"symbol"`
	Name      string `json:"name"`
	Date      string `json:"date"`
	IsEnabled bool `json:"isEnabled"`
	Type      string `json:"isEnabled"`
	IexID     string `json:"iexId"`
}
