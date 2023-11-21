package migrations

import (
	"go-web-echo-app/config"
	"go-web-echo-app/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	DB = config.SetupDatabaseConnection()
	DB.AutoMigrate(&models.User{})
}
