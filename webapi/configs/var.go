package configs

import (
	"github.com/abe27/webapi/models"
	"gorm.io/gorm"
)

var (
	Store           *gorm.DB
	API_NAME        string
	API_DESCRIPTION string
	API_PORT        string
	API_DBNAME      string
)

func SetDB() {
	Store.AutoMigrate(&models.Device{})
	Store.AutoMigrate(&models.LineToken{})
	Store.AutoMigrate(&models.Notification{})
	Store.AutoMigrate(&models.TempData{})
}
