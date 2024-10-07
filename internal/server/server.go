package server

import (
	"fmt"

	_ "github.com/beardfriend/ddalggak/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/swagger"
)

// @title DDALGGAK API
// @version 1.0
// @description DDALGGAK API
// @contact.name SEHUN PARK
// @contact.email beardfriend21@gmail.com
func Run() {
	app := fiber.New()
	app.Use(favicon.New())
	app.Get("/swagger/*", swagger.New(swagger.ConfigDefault))

	if err := app.Listen(fmt.Sprintf(":%d", 4000)); err != nil {
		panic(err)
	}
}
