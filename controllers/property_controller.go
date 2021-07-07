package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/mofodox/omh-rest-api/database"
	"github.com/mofodox/omh-rest-api/models"
)

func AddProperty(ctx *fiber.Ctx) {
	var property models.Property

	if err := ctx.BodyParser(&property); err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "Unable to process your JSON request",
		})
	}

	database.DB.Create(&property)

	if property.ID != 0 {
		err := database.DB.Debug().Model(&models.Country{}).Where("id = ?", property.CountryID).Take(&property.Country).Error
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code": fiber.StatusInternalServerError,
				"message": "Create porperty: server error",
			})
		}
	}

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
		database.DB.Preload("Country").Find(&properties)
	}

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Found all properties",
		"properties": properties,
		"query": modeQuery,
	})
}

func ShowProperty(ctx *fiber.Ctx) {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "Unable to process your JSON request",
		})
	}

	var property models.Property
	db := database.DB.Preload("Country").Find(&property, id)

	log.Printf("GET Property - Rows affected: %v\n", db.RowsAffected)

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Found property",
		"property": property,
	})
}

func UpdateProperty(ctx *fiber.Ctx) {
	property := new(models.Property)

	if err := ctx.BodyParser(property); err != nil {
		ctx.Status(422).JSON(fiber.Map{
			"code": 422,
			"message": "Unable to parse json values",
		})

		return
	}

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "error converting the id to int64",
		})
	}

	var propertyDB models.Property
	database.DB.First(&propertyDB, id)
	db := database.DB.Model(&propertyDB).Updates(map[string]interface{}{
		"Address": property.Address,
		"Bedroom": property.Bedroom,
		"Bathroom": property.Bathroom,
		"Price": property.Price,
		"Sqm": property.Sqm,
		"Mode": property.Mode,
		"HomeType": property.HomeType,
	})

	log.Printf("PUT - Rows affected: %v\n", db.RowsAffected)

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Property updated successfully",
		"property": propertyDB,
	})
}

func DeleteProperty(ctx *fiber.Ctx) {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "error converting the id to int64",
		})
	}
	
	var property models.Property

	database.DB.Find(&property, id)
	database.DB.Delete(&property)

	ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Property successfully deleted",
	})
}
