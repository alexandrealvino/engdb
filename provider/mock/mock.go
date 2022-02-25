package mock

import "github.com/gofiber/fiber/v2"

var SuccessProvider = map[string]interface{}{
	"membros": &fiber.Map{
		"msisdn":         "string",
		"status":         "Ativo",
		"tipo":           "M",
		"ativo_desde":    "string",
		"apelido_membro": "string",
		"docnumber":      "string",
		"profile_name":   "string",
		"cod_plan":       "string",
	},
	"status":  "Ativo",
	"qtdDisp": 0,
}

var ErrorProvider = map[string]interface{}{
	"timestamp":    0,
	"status":       0,
	"error":        "string",
	"exception":    "string",
	"errors":       []fiber.Map{},
	"message":      "string",
	"path":         "string",
	"internalCode": "string",
}
