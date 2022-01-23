package main

import "github.com/gofiber/fiber/v2"
import "Learngo/06_modules/externalmymodule"

// this uses only init(s) method run
// import _ "Learngo/06_modules/externalmymodule"

// this allows us to access methods and properties without typing external module name
// import . "Learngo/06_modules/externalmymodule"
// example: Like "GetMyData()" instead of "externalmymodule.GetMyData()"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	externalmymodule.GetMyData()
	app.Listen(":3000")
}
