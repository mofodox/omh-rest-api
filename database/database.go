package database

import (
	"fmt"
	"log"
	"os"

	"github.com/mofodox/omh-rest-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("unable to connect to the database")
	} else {
		log.Println("connected to the database successfully")
	}

	DB = connection

	connection.Debug().Migrator().DropTable(models.Property{})
	connection.Debug().AutoMigrate(models.Property{}, models.Country{})
}
