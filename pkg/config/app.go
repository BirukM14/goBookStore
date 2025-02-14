package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	DB *gorm.DB
)

func ConnectDB() {
	database, err := gorm.Open("mysql", "BirukM14:Bb@0939852436@simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB = database // ✅ Corrected variable name
}

func GetDB() *gorm.DB {
	return DB // ✅ Corrected variable name
}
