package config

import (
	"MS1/constants"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

func InitDB() *gorm.DB {
	log.Infoln("InitDB() starting......")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", constants.DbUserName, constants.DbPassword,
		constants.DbHost, constants.DbPort, constants.DbName)
	log.Infoln("URL: ", URL)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Infoln("Database connected successfully.....")
	log.Infoln("InitDB() function ended.....")
	return db
}
