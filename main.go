package main

import (
	utiler "deneme/varaibles"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var student utiler.StudentDTO = utiler.CreateStudent(15, "Dumy", "Dum", 1)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(student)
	})

	app.Listen(":3000")

}
