package models

import (
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/garyburd/redigo/redis"
)

type DataBase struct {
	DB *gorm.DB
	Search *gorm.DB
	MySql *gorm.DB
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

	self.Search, err = gorm.Open("mysql", "tcp(sphinx:9306)/")

	if err != nil {
		logrus.Error(err)
	}

	self.MySql, err = gorm.Open("mysql", "root:a3d1sw2qwe@/stock_data?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		logrus.Error(err)
	}

	items := CreateFish()
	self.CreateTablesIfNotExists(items)
	self.Migration(items)
	items = CreateMysql()
	self.CreateTablesIfNotExistsMysql(items)

	return self
}

func CreateMysql() []interface{} {
	var items []interface{}

	items = append(items, &Company{})

	return items
}

func CreateFish() []interface{} {
	var items []interface{}
	items = append(items, &User{})
	items = append(items, &UsersUsers{})
	items = append(items, &Provider{})
	items = append(items, &ProvidersUsers{})
	items = append(items, &Ticker{})
	items = append(items, &Symbol{})
	items = append(items, &Company{})
	items = append(items, &Tag{})
	items = append(items, &Sector{})
	items = append(items, &CompaniesTags{})

	return items
}

func (self *DataBase) CreateTablesIfNotExists(items []interface{}) {
	for _, v := range items {
		if !self.DB.HasTable(v) {
			self.DB.CreateTable(v)
		}
	}
}

func (self *DataBase) CreateTablesIfNotExistsMysql(items []interface{}) {
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

	self.DB.Model(&Symbol{}).AddIndex("symbol_addr_idx", "abbr")
	self.DB.Model(&Symbol{}).AddIndex("symbol_iex_id_idx", "iex_id")

	self.DB.Model(&Company{}).AddIndex("company_name_idx", "name")
	self.DB.Model(&Company{}).AddIndex("company_sector_id_idx", "sector_id")
	self.DB.Model(&Company{}).AddIndex("company_symbol_id_idx", "symbol_id")

	self.DB.Model(&CompaniesTags{}).AddIndex("ct_company_id_idx", "company_id")
	self.DB.Model(&CompaniesTags{}).AddIndex("ct_tag_id_idx", "tag_id")

}

func (self *DataBase) Close() {
	self.DB.Close()
}



