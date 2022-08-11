package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(loadDbConfig(dbConf))
	if err != nil {
		log.Fatal("can't connect to the database")
	}
	DB = db

	if !DB.HasTable(&ResponseTimeLog{}) {
		DB.CreateTable(&ResponseTimeLog{})
	} else {
		db.Where("id > 0").Delete(ResponseTimeLog{})
	}
	return db
}
