package controllers

import (
	"log"
	"strconv"

	"github.com/mofodox/omh-rest-api/database"
	"github.com/mofodox/omh-rest-api/models"

	"github.com/gofiber/fiber"
)

func CreateCountry(ctx *fiber.Ctx) {
	var country models.Country

	if err := ctx.BodyParser(&country); err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "Unable to process your JSON request",
		})
	}

	database.DB.Create(&country)

	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code": fiber.StatusCreated,
		"message": "Country successfull created",
		"country": country, 
	})
}

func GetAllCountries(ctx *fiber.Ctx) {
	var countries []models.Country

	database.DB.Find(&countries)

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "All countries found",
		"countries": countries,
	})
}

func GetCountry(ctx *fiber.Ctx) {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "Unable to process your JSON request",
		})
	}

	var country models.Country

	db := database.DB.Find(&country, id)

	log.Printf("GET Country - Rows affected: %v\n", db.RowsAffected)

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Found country",
		"country": country,
	})
}

func UpdateCountry(ctx *fiber.Ctx) {
	country := new(models.Country)

	if err := ctx.BodyParser(country); err != nil {
		ctx.Status(422).JSON(fiber.Map{
			"code": 422,
			"message": "Unable to parse JSON values",
		})
	}

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "error converting the id to int64",
		})
	}

	var countryDB models.Country
	database.DB.First(&countryDB, id)
	db := database.DB.Model(&countryDB).Updates(map[string]interface{}{
		"Name": country.Name,
	})

	log.Printf("PUT - Rows affected: %v\n", db.RowsAffected)

	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
		"message": "Country updated succesfully",
		"country": country,
	})
}

func DeleteCountry(ctx *fiber.Ctx) {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		ctx.Status(442).JSON(fiber.Map{
			"code": 442,
			"message": "error converting the id to int64",
		})
	}

	var country models.Country
	database.DB.Find(&country, id)
	database.DB.Delete(&country)

	ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Country successfully deleted",
	})
}
