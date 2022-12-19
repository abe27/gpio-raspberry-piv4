package controllers

import (
	"time"

	"github.com/abe27/temp/webapi/models"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	var r models.Response
	r.Message = "Hello, World"
	r.OnTime = time.Now()
	return c.Status(fiber.StatusOK).JSON(&r)
}
