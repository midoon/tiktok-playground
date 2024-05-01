package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDBOpenConenction() *gorm.DB {
	GormUrl := "root:root@tcp(127.0.0.1:3306)/http_code?charset=utf8mb4&parseTime=True&loc=Local"
	dialect := mysql.Open(GormUrl)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Fatalf("failed create connection to DB : %s", err.Error())
	}
	return db
}
