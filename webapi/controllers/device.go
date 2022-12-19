package controllers

import (
	"fmt"

	"github.com/abe27/webapi/configs"
	"github.com/abe27/webapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllDevice(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Device
	if err := configs.Store.Where("is_active=?", true).Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Show All Device"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowDeviceByID(c *fiber.Ctx) error {
	var r models.Response
	var obj models.Device
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = fmt.Sprintf("Show By ID: %s", c.Params("id"))
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateDevice(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Device
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var obj models.Device
	obj.Name = frm.Name
	obj.OnPin = frm.OnPin
	obj.AlertOn = frm.AlertOn
	obj.IsActive = frm.IsActive
	if err := configs.Store.Create(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = fmt.Sprintf("Create Device ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func UpdateDeviceByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Device
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var obj models.Device
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	obj.Name = frm.Name
	obj.OnPin = frm.OnPin
	obj.AlertOn = frm.AlertOn
	obj.IsActive = frm.IsActive
	if err := configs.Store.Save(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = fmt.Sprintf("Update Device ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DeleteDeviceByID(c *fiber.Ctx) error {
	var r models.Response

	var obj models.Device
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Where("id=?", c.Params("id")).Delete(&models.Device{}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = fmt.Sprintf("Delete Device ID: %s is completed!", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
