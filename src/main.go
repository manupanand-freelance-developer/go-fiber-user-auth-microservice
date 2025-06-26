package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/manupanand/01-fiber/model"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {


	app := fiber.New()
    // ✅ Apply CORS middleware globally
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",   // here -origin-frontend Allow all origins - change to specific domains in production
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: false, // for specific witthout wild card set -true
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is manupanands testing api...")
	})
	app.Get("/:value", func(c *fiber.Ctx) error {
		fmt.Println(model.Name)
		return c.SendString("value is:" + c.Params("value"))
		
	})
	app.Get("/health",func(c *fiber.Ctx) error {
		return c.SendString("Health is ok..")
	})

	// ✅ Start server
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
