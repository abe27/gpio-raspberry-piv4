package configs

import "gorm.io/gorm"

var (
	Store           *gorm.DB
	API_NAME        string
	API_DESCRIPTION string
	API_PORT        int
	API_DBNAME      string
)

func SetUp() {

}
