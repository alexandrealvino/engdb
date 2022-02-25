package provider

import (
	"engdb/provider/mock"
	"github.com/gofiber/fiber/v2"
)

func NewProviderHandler(route fiber.Router) {
	route.Get("/service-sgg", get)
}

///api/service-sgg/sgg/getGroup/msisdn/{msisdn}/status/{status}

func get(c *fiber.Ctx) error {
	//return c.Status(fiber.StatusInternalServerError).JSON(ErrorProvider)

	// getting request parameters
	//header := c.Request().Header.Header()

	status := c.FormValue("status")
	msisdn := c.FormValue("msisdn")
	//println("header", string(header))
	println("status:", status, "msisdn:", msisdn)

	messageId := c.GetReqHeaders()["Messageid"]

	println("messageId: ", messageId)

	return c.Status(fiber.StatusOK).JSON(mock.SuccessProvider)
}

//quoteProvider := &QuoteProvider{}
//
//customContext, cancel := context.WithCancel(context.Background())
//defer cancel()
//
//err := h.quoteProviderService.BuildQuoteProvider(customContext, quoteProvider)
//if err != nil {
//	return c.Status(fiber.StatusInternalServerError).JSON(ErrorInternal)
//}
//
//err = h.quoteProviderService.BuildQuoteProvider(customContext, quoteProvider)
//if err != nil {
//	return c.Status(fiber.StatusInternalServerError).JSON(ErrorProvider)
//}
