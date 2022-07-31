package services

import (
	"goauth/models"

	"github.com/jinzhu/gorm"
)

var dbconn *gorm.DB

func SetDB(db *gorm.DB) {
	dbconn = db
	var user = models.GetUser()
	dbconn.AutoMigrate(&user)
}

func GetDb() *gorm.DB {
	return dbconn
}
