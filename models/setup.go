package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)
var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/interview?parseTime=true"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Activity{})
	database.AutoMigrate(&Todo{})

	DB = database
}