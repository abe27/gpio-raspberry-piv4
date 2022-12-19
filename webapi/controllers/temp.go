package controllers

import (
	"fmt"
	"time"

	"github.com/abe27/webapi/configs"
	"github.com/abe27/webapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllTempData(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.TempData
	if err := configs.Store.Where("is_active=?", true).Preload("Device").Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Show All TempData"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowTempDataByID(c *fiber.Ctx) error {
	var r models.Response
	var obj models.TempData
	if err := configs.Store.Where("id=?", c.Params("id")).Preload("Device").First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = fmt.Sprintf("Show By ID: %s", c.Params("id"))
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateTempData(c *fiber.Ctx) error {
	var r models.Response
	var frm models.TempData
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var device models.Device
	if err := configs.Store.Where("name=?", frm.DeviceID).First(&device).Error; err != nil {
		r.Message = fmt.Sprintf("Device ID: %s is %s", frm.ID, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var obj models.TempData
	obj.DeviceID = device.ID
	obj.OnDateTime = time.Now()
	obj.Temp = frm.Temp
	obj.Humidity = frm.Humidity
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive
	if err := configs.Store.Create(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	obj.Device = device
	r.Message = fmt.Sprintf("Create TempData ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func UpdateTempDataByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.TempData
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var obj models.TempData
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var device models.Device
	if err := configs.Store.Where("name=?", frm.DeviceID).First(&device).Error; err != nil {
		r.Message = fmt.Sprintf("Device ID: %s is %s", frm.ID, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	obj.DeviceID = device.ID
	// obj.OnDateTime = time.Now()
	obj.Temp = frm.Temp
	obj.Humidity = frm.Humidity
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive
	if err := configs.Store.Save(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	obj.Device = device
	r.Message = fmt.Sprintf("Update TempData ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DeleteTempDataByID(c *fiber.Ctx) error {
	var r models.Response

	var obj models.TempData
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Where("id=?", c.Params("id")).Delete(&models.TempData{}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = fmt.Sprintf("Delete TempData ID: %s is completed!", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
