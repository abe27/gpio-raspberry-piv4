package controllers

import (
	"fmt"

	"github.com/abe27/webapi/configs"
	"github.com/abe27/webapi/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllLineToken(c *fiber.Ctx) error {
	var r models.Response
	var obj []models.LineToken
	if err := configs.Store.Where("is_active=?", true).Find(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = "Show All LineToken"
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func ShowLineTokenByID(c *fiber.Ctx) error {
	var r models.Response
	var obj models.LineToken
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	r.Message = fmt.Sprintf("Show By ID: %s", c.Params("id"))
	r.Data = &obj
	return c.Status(fiber.StatusOK).JSON(&r)
}

func CreateLineToken(c *fiber.Ctx) error {
	var r models.Response
	var frm models.LineToken
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var obj models.LineToken
	obj.Token = frm.Token
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive
	if err := configs.Store.Create(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = fmt.Sprintf("Create LineToken ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func UpdateLineTokenByID(c *fiber.Ctx) error {
	var r models.Response
	var frm models.LineToken
	if err := c.BodyParser(&frm); err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotAcceptable).JSON(&r)
	}

	var obj models.LineToken
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	obj.Token = frm.Token
	obj.Description = frm.Description
	obj.IsActive = frm.IsActive
	if err := configs.Store.Save(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(&r)
	}

	r.Message = fmt.Sprintf("Update LineToken ID: %s is completed!", obj.ID)
	r.Data = &obj
	return c.Status(fiber.StatusCreated).JSON(&r)
}

func DeleteLineTokenByID(c *fiber.Ctx) error {
	var r models.Response

	var obj models.LineToken
	if err := configs.Store.Where("id=?", c.Params("id")).First(&obj).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}

	if err := configs.Store.Where("id=?", c.Params("id")).Delete(&models.LineToken{}).Error; err != nil {
		r.Message = err.Error()
		return c.Status(fiber.StatusNotFound).JSON(&r)
	}
	r.Message = fmt.Sprintf("Delete LineToken ID: %s is completed!", c.Params("id"))
	return c.Status(fiber.StatusOK).JSON(&r)
}
