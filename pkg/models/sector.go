package models

import (
	"github.com/jinzhu/gorm"
)

type Sector struct {
	gorm.Model
	SectorID uint
	Name     string
}