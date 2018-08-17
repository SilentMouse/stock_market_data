package models

import (
	"github.com/jinzhu/gorm"
)


type Company struct{
	gorm.Model
	Name string
}

