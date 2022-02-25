package adapter

import (
	"engdb/models"
	"github.com/gofiber/fiber/v2"
)

//TODO request to provider
type ProviderImpl struct {
}

type Provider interface {
	GetRGroup(c *fiber.Ctx, rgroup models.RGroup) error
}

func (p ProviderImpl) GetRGroup(c *fiber.Ctx, rgroup models.RGroup) error {
	//TODO log header

	// http.newrequest
	// log c.GetReqHeaders()
	// check statuscode

	// transform error msg
	return c.Status(fiber.StatusOK).JSON(rgroup.Status)
}
