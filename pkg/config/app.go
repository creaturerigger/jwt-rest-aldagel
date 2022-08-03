package config

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func getDatabaseEnvVars() (map[string]string, error) {
	varMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}
	return varMap, nil
}

func Connect() {
	var err error
	varMap, _ := getDatabaseEnvVars()
	usrName := varMap["DB_USER"]
	pass := varMap["DB_USER_PASSWORD"]
	dbName := varMap["DB_NAME"]
	servAddr := varMap["SERVER_ADDR"]
	db, err = gorm.Open(mysql.Open(usrName+":"+pass+"@tcp("+servAddr+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
