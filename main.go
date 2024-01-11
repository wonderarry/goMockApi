package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/wonderarry/goMockApi/database"
	"github.com/wonderarry/goMockApi/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API")
}

func importRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/user", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	importRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
