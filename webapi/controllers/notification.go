package controllers

import (
	"fmt"

	"github.com/abe27/webapi/configs"
	"github.com/abe27/webapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllNotification(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.Notification
	if err := configs.Store.Preload("Device").Preload("LineToken").Where("is_active=?", true).Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Show All Notification"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowNotificationByID(c *fiber.Ctx) error {
	var r models.Response
	var obj models.Notification
	if err := configs.Store.Preload("Device").Preload("LineToken").Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = fmt.Sprintf("Show By ID: %s", c.Params("id"))
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateNotification(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Notification
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var device models.Device
	if err := configs.Store.Where("name=?", frm.DeviceID).First(&device).Error; err != nil {
		r.Message = fmt.Sprintf("Device ID: %s is %s", frm.ID, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var line models.LineToken
	if err := configs.Store.Where("token=?", frm.LineTokenID).First(&line).Error; err != nil {
		r.Message = fmt.Sprintf("Device ID: %s is %s", frm.ID, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var obj models.Notification
	obj.DeviceID = device.ID
	obj.LineTokenID = line.ID
	obj.IsAccept = frm.IsAccept
	obj.IsActive = frm.IsActive
	if err := configs.Store.Create(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}
	obj.Device = device
	obj.LineToken = line
	r.Message = fmt.Sprintf("Create Notification ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func UpdateNotificationByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.Notification
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var obj models.Notification
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var device models.Device
	if err := configs.Store.Where("name=?", frm.DeviceID).First(&device).Error; err != nil {
		r.Message = fmt.Sprintf("Device ID: %s is %s", frm.DeviceID, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	var line models.LineToken
	if err := configs.Store.Where("token=?", frm.LineTokenID).First(&line).Error; err != nil {
		r.Message = fmt.Sprintf("Token ID: %s is %s", frm.LineTokenID, err.Error())
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	obj.DeviceID = device.ID
	obj.LineTokenID = line.ID
	obj.IsAccept = frm.IsAccept
	obj.IsActive = frm.IsActive
	if err := configs.Store.Save(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	obj.Device = device
	obj.LineToken = line
	r.Message = fmt.Sprintf("Update Notification ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DeleteNotificationByID(c *fiber.Ctx) error {
	var r models.Response

	var obj models.Notification
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Where("id=?", c.Params("id")).Delete(&models.Notification{}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = fmt.Sprintf("Delete Notification ID: %s is completed!", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
