package database

import (
	"log"

	"github.com/mofodox/omh-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3306)/omh_backend?charset=utf8mb4&parseTime=True&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("unable to connect to the database")
	} else {
		log.Println("connected to the database successfully")
	}

	DB = connection

	connection.Debug().Migrator().DropTable(models.Property{}, models.Country{})
	connection.Debug().AutoMigrate(models.Property{}, models.Country{})
}
