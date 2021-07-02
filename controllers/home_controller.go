package controllers

import "github.com/gofiber/fiber"

func Home(ctx *fiber.Ctx) {
	ctx.SendString("Hello World")
}
