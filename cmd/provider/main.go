package main

import (
	"engdb/provider"
	"engdb/provider/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many! Please wait!",
			})
		},
	}))

	provider.NewProviderHandler(app.Group("/api"))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(mock.SuccessProvider)
	})

	log.Fatal(app.Listen(":6000"))
}

//logger.Start()
//defer logger.Sync()
//logger.Info("run quote", "infrastructure/quotePmid")
//
//quoteProviderRepository := quoteProvider.NewQuoteProviderRepository("fictitious")
//quoteProviderService := quoteProvider.NewQuoteProviderService(quoteProviderRepository)
