package models

import (
	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{}, &Order{})
}
