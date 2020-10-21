package main

import (
	"github.com/JoelTinx/test-go/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api.SetupRoutes(app)
	app.Listen(":3000")
}