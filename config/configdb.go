package config

import (
	"MS1/constants"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", constants.DbUserName, constants.DbPassword,
		constants.DbHost, constants.DbPort, constants.DbName)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("connected successfully.")
	return db
}
