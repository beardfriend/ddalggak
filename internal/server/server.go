package server

import (
	"fmt"

	_ "github.com/beardfriend/ddalggak/docs"
	"github.com/beardfriend/ddalggak/pkg/validatorx"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/favicon"
)

// @title DDALGGAK API
// @version 1.0
// @description DDALGGAK API
// @contact.name SEHUN PARK
// @contact.email beardfriend21@gmail.com
func Run() {
	app := fiber.New(fiber.Config{
		StructValidator: validatorx.NewValidatorx().
			AddUrlValidation("url").
			AddPhoneNumValidation("phoneNum").
			Init(),
	})
	app.Use(favicon.New())

	if err := app.Listen(fmt.Sprintf(":%d", 4000)); err != nil {
		panic(err)
	}
}
