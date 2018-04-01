package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Provider struct {
	gorm.Model
	Name      string
	FirstName string
	LastName  string
	Nickname  string
	PID       string
	Email     string
	Phone     string
	MainImage string
}

type ProvidersUsers struct {
	gorm.Model
	ProviderID uint64
	Provider   Provider
	UserID     uint64
	User       User
}
