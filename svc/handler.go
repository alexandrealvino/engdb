package svc

import (
	"engdb/models"
	"engdb/svc/adapter"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type Service struct {
	Provider  adapter.Provider
	Validator *validator.Validate
}

func NewService(provider adapter.Provider, validate *validator.Validate) *Service {
	return &Service{
		Provider:  provider,
		Validator: validate,
	}
}

func (s *Service) NewServiceHandler() *fiber.App {
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

	route := app.Group("/api")

	route.Get("/v1/r-group", s.getGroup)

	return app
}

func (s *Service) getGroup(c *fiber.Ctx) error {
	// statusdomain A:A
	// getting parameters from request
	group := models.RGroup{
		MessageId:        c.GetReqHeaders()["Messageid"],
		ClientId:         c.GetReqHeaders()["Clientid"],
		AuthorizationOam: c.GetReqHeaders()["Authorizationoam"],
		Msisdn:           c.FormValue("msisdn"),
		Status:           c.FormValue("status"),
	}

	validationErrors := s.Validator.Struct(group)
	if validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error: err invalid request!")
	}

	//TODO request
	err := s.Provider.GetRGroup(c, group)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(group)
}
