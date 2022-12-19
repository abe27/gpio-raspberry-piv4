package routes

import (
	"github.com/abe27/webapi/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.Hello)

	r := c.Group("/api/v1")
	r.Get("", controllers.Hello)
}
