package models

import (
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/garyburd/redigo/redis"
)

type DataBase struct {
	DB *gorm.DB
}

var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle: 5,
	Wait: true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	},
}

func InitDB() DataBase {
	var addr string
	val, ok := os.LookupEnv("DB_HOST")
	if !ok {
		addr = "postgresql://root@localhost:26257/stock_market_data?sslmode=disable"
	} else {
		addr = val
	}
	var self DataBase
	var err error
	self.DB, err = gorm.Open("postgres", addr)
	if err != nil {
		logrus.Fatal(err)
	}

	items := CreateFish()
	self.CreateTablesIfNotExists(items)
	self.Migration(items)

	return self
}

func CreateFish() []interface{} {
	var items []interface{}
	items = append(items, &User{})
	items = append(items, &UsersUsers{})
	items = append(items, &Provider{})
	items = append(items, &ProvidersUsers{})

	return items
}

func (self *DataBase) CreateTablesIfNotExists(items []interface{}) {
	for _, v := range items {
		if !self.DB.HasTable(v) {
			self.DB.CreateTable(v)
		}
	}
}

func (self *DataBase) Migration(items []interface{}) {

	self.DB.Model(&User{}).AddIndex("users_id_idx", "id")
	self.DB.Model(&User{}).AddIndex("users_entity_id_idx", "entity_id")
	self.DB.Model(&User{}).AddIndex("users_token_idx", "token")

	self.DB.Model(&Provider{}).AddIndex("providers_id_idx", "id")
	self.DB.Model(&Provider{}).AddIndex("providers_p_id_idx", "p_id")

	self.DB.Model(&ProvidersUsers{}).AddIndex("pr_users_id_idx", "id")
	self.DB.Model(&ProvidersUsers{}).AddIndex("pr_users_id_idx", "provider_id")
	self.DB.Model(&ProvidersUsers{}).AddIndex("pr_users_id_idx", "user_id")

	self.DB.Model(&UsersUsers{}).AddIndex("us_us_id_idx", "id")
	self.DB.Model(&UsersUsers{}).AddIndex("us_us_first_user_id_idx", "first_user_id")
	self.DB.Model(&UsersUsers{}).AddIndex("us_us_second_user_id_idx", "second_user_id")
}

func (self *DataBase) Close() {
	self.DB.Close()
}



