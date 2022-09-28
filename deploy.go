package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New(".", ".html")
	app := fiber.New(fiber.Config{Views: engine,})

	var handler = func(c *fiber.Ctx) error { return c.Render("index", fiber.Map{}) }

	app.Get("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	
	app.Static("/static", "./style")

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
