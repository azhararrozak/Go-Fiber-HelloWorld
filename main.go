package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.JSON(fiber.Map{
    //         "message": "Hello World!",
    //     })
    // })

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

    app.Listen(":4000")
}