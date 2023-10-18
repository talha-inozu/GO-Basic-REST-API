package main

import (
	models "deneme/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var student models.StudentDTO = models.CreateStudent(15, "Dumy", "Dum", 1)
	app := fiber.New()

	app.Get("/getStudent", func(c *fiber.Ctx) error {
		return c.JSON(student)
	})

	app.Get("/shutdown", func(c *fiber.Ctx) error {
		app.Shutdown()
		return nil
	})

	app.Listen(":3000")

}
