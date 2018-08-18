package models

import (
	_ "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string
	FirstName   string
	SecondName  string
	Phone       string
	Email       string
	MainImage   string
	Token       string
	Friends     []User
	Subscribers []User
	EntityID    uint64
}

type UsersUsers struct {
	gorm.Model
	RelationType string
	FirstUserId  uint64
	SecondUserId uint64
}
