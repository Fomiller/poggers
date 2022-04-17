package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// app.Static("/", "./public")

	// api := app.Group("/api", apiHandler)
	// v1 := api.Group("/v1", v1Handler)

	// v1.Get("/*", func(c *fiber.Ctx) error {
	// 	return c.SendString("API path: " + c.Params("*"))
	// })

	// app.Get("/:name?", func(c *fiber.Ctx) error {
	// 	if c.Params("name") != "" {
	// 		return c.SendString("name: " + c.Params("name"))
	// 	}
	// 	return c.SendString("You have no name?")
	// })

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Word": "Hello, World!",
		})
	})

	app.Listen(":3000")
}

// func apiHandler(c *fiber.Ctx) error {
// 	return c.SendString("API")
// }

// func v1Handler(c *fiber.Ctx) error {
// 	return c.SendString("v1")
// }
