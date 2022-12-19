package routes

import (
	"github.com/abe27/webapi/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetUpRouter(c *fiber.App) {
	c.Get("/", controllers.Hello)

	r := c.Group("/api/v1")
	r.Get("", controllers.Hello)

	// Device
	rDevice := r.Group("/device")
	rDevice.Get("", controllers.GetAllDevice)
	rDevice.Get("/:id", controllers.ShowDeviceByID)
	rDevice.Post("", controllers.CreateDevice)
	rDevice.Put("/:id", controllers.UpdateDeviceByID)
	rDevice.Delete("/:id", controllers.DeleteDeviceByID)

	// Line Token
	rLineToken := r.Group("/token")
	rLineToken.Get("", controllers.GetAllLineToken)
	rLineToken.Get("/:id", controllers.ShowLineTokenByID)
	rLineToken.Post("", controllers.CreateLineToken)
	rLineToken.Put("/:id", controllers.UpdateLineTokenByID)
	rLineToken.Delete("/:id", controllers.DeleteLineTokenByID)

	// Notification
	rNotification := r.Group("/notification")
	rNotification.Get("", controllers.GetAllNotification)
	rNotification.Get("/:id", controllers.ShowNotificationByID)
	rNotification.Post("", controllers.CreateNotification)
	rNotification.Put("/:id", controllers.UpdateNotificationByID)
	rNotification.Delete("/:id", controllers.DeleteNotificationByID)

	// Temp
	rTemp := r.Group("/temp")
	rTemp.Get("", controllers.GetAllTempData)
	rTemp.Get("/:id", controllers.ShowTempDataByID)
	rTemp.Post("", controllers.CreateTempData)
	rTemp.Put("/:id", controllers.UpdateTempDataByID)
	rTemp.Delete("/:id", controllers.DeleteTempDataByID)
}
