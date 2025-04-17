package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/books", func(c *fiber.Ctx) error {
		log.Println("GET /books called")
		return c.JSON(books)
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		book := new(Book)
		if err := c.BodyParser(book); err != nil {
			log.Println("Failed to parse body:", err)
			return err
		}
		books = append(books, *book)
		fmt.Printf("Book added: %+v\n", *book)
		return c.JSON(book)
	})

	log.Println("Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
