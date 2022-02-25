package main

import (
	"engdb/svc"
	"engdb/svc/adapter"
	"github.com/go-playground/validator/v10"
)

func main() {

	provider := adapter.ProviderImpl{}
	validate := validator.New()
	service := svc.NewService(provider, validate)

	app := service.NewServiceHandler()

	// log.StdSetup

	app.Listen(":3000")
}
