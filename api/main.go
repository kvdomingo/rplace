package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kvdomingo/rplace/pkg/redis"
)

func main() {
	app := fiber.New(
		fiber.Config{
			AppName:           "r/place",
			ServerHeader:      "r/place",
			EnablePrintRoutes: true,
		},
	)

	api := app.Group("/api")

	api.Get(
		"/", func(c *fiber.Ctx) error {
			return c.JSON(
				&map[string]string{
					"api": "ok",
					"db":  redis.Health(),
				},
			)
		},
	)

	log.Fatal(app.Listen(":8000"))
}
