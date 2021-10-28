package config

import (
	"os"
	"test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() {
	connection := os.Getenv("CONNECTION")

	var err error
	Db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitMigrate() {

	Db.Migrator().DropTable(&models.User{})
	Db.Migrator().DropTable(&models.Product{})

	Db.AutoMigrate(&models.User{})
	Db.AutoMigrate(&models.Product{})

}
