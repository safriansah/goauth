package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	pass = "postgres"
	dbse = "postgres"
	ssl  = "disable"
)

func GetConnection() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, dbse, ssl)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		log.Println(err)
		panic("cannot connect database")
	}

	log.Println("success connect databse")
	return db
}
