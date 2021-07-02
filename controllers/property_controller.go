package controllers

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/mofodox/omh-rest-api/database"
	"github.com/mofodox/omh-rest-api/models"
	"gorm.io/gorm"
)

func AddProperty(ctx *fiber.Ctx) {
	var property models.Property

	if err := ctx.BodyParser(&property); err != nil {
		log.Fatalln(err)
	}

	database.DB.Create(&property)

	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code": fiber.StatusCreated,
		"message": "Property successfully created",
		"property": property,
	})
}

func ShowAllProperties(ctx *fiber.Ctx) {
	modeQuery := ctx.Query("mode")

	var properties []models.Property

	if modeQuery == "Rental" {
		database.DB.Model(&models.Property{}).Where("Mode = ?", modeQuery).Find(&properties)
	} else if modeQuery == "Sale" {
		database.DB.Model(&models.Property{}).Where("Mode = ?", modeQuery).Find(&properties)
	} else {
		database.DB.Find(&properties)
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Found all properties",
		"properties": properties,
		"query": modeQuery,
	})
}

func ShowProperty(ctx *fiber.Ctx) {
	id := ctx.Params("ID")

	var property models.Property
	database.DB.Find(&property, id)

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Found property",
		"property": property,
	})
}


// TODO: Update handler for property still a work in progress
func UpdateProperty(ctx *fiber.Ctx) {
	id, _ := strconv.Atoi(ctx.Params("ID"))

	property := new(models.Property)
	
	err := ctx.BodyParser(&property)
	if err != nil {
		log.Fatalln(err)
	}

	err = database.DB.Model(&property).Where("ID = ?", id).Updates(&property).Find(&property, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code": fiber.StatusNotFound,
			"message": "Property not found",
		})
	} else {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code": fiber.StatusInternalServerError,
			"message": "Internal server error",
		})
	}
	
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Property updated successfully",
		"property": property,
	})
}

func DeleteProperty(ctx *fiber.Ctx) {
	id := ctx.Params("ID")
	
	var property models.Property

	database.DB.Find(&property, id)
	database.DB.Delete(&property)

	ctx.Status(fiber.StatusSeeOther).JSON(fiber.Map{
		"message": "Property successfully deleted",
	})
}
