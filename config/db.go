package config

import (
	"kp-golang-restfull-api/structs"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/kp-golang-restful-api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		println(err)
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.User{})
	return db
}
