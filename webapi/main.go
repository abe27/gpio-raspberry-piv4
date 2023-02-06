package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abe27/webapi/configs"
	"github.com/abe27/webapi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Error loading .env.local file")
	}

	// Initial ENV
	configs.API_NAME = os.Getenv("API_NAME")
	configs.API_DESCRIPTION = os.Getenv("API_DESCRIPTION")
	configs.API_PORT = os.Getenv("API_PORT")
	configs.API_DBNAME = os.Getenv("API_DBNAME")

	configs.Store, err = gorm.Open(sqlite.Open(configs.API_DBNAME), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		SkipDefaultTransaction:                   true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tbt_", // table name prefix, table for `User` would be `t_users`
			SingularTable: false,  // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,  // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// After connect db is completed
	configs.SetDB()
}

func main() {
	// Create config variable
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  configs.API_DESCRIPTION, // add custom server header
		AppName:       configs.API_NAME,
		BodyLimit:     10 * 1024 * 1024, // this is the default limit of 10MB
	}

	app := fiber.New(config)
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New())
	app.Static("/", "./public")
	routes.SetUpRouter(app)
	app.Listen(fmt.Sprintf(":%s", configs.API_PORT))
}
