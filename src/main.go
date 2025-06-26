package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/manupanand/01-fiber/model"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "User service",
		AppName:       "User authentication App v1.0.1",
	})
	// ✅ Apply CORS middleware globally
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // here -origin-frontend Allow all origins - change to specific domains in production
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: false, // for specific witthout wild card set -true
	}))
	//Rate limit middleware
	app.Use(limiter.New(limiter.Config{
			Max:        10,                    // Max 10 requests
			Expiration: 1 * time.Minute,       // Per 1 minute
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.IP()                  // Rate-limit per IP (default)
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(429).SendString("Rate limit exceeded. Try again later.")
			},
		}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is manupanands testing api...")
	})
	app.Get("/id/:value", func(c *fiber.Ctx) error {
		fmt.Println(model.Name)
		return c.SendString("id  is:" + c.Params("value"))

	})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Health is ok..")
	})
	app.Static("/doc", "./public", fiber.Static{
		Compress:      true,             // Serve compressed files
		ByteRange:     true,             // Enable video/audio seeking
		Browse:        false,             // Allow file browsing if no index.html
		Index:         "index.html",     // Serve index.html for /
		CacheDuration: 10 * time.Second, // Fiber internal caching
		MaxAge:        300,             // 5min browser cache
	})

	// ✅ Start server
	app.Server().MaxConnsPerIP=1
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
