package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/manupanand/01-fiber/model"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	app.Get("/:value", func(c *fiber.Ctx) error {
		fmt.Println(model.Name)
		return c.SendString("value is:" + c.Params("value"))
		
	})

	app.Listen(":3000")
	
}
