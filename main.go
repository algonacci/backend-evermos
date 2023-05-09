package main

import (
	"github.com/algonacci/backend-evermos/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Routes
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)

	shop := app.Group("/shops")
	shop.Get("/:id", handler.GetShop)
	shop.Post("/", handler.CreateShop)
	shop.Put("/:id", handler.UpdateShop)
	shop.Delete("/:id", handler.DeleteShop)

	app.Listen(":3000")
}
