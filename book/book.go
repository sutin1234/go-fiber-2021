package book

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/sutin1234/go-fiber-2021/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}
func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).SendString("no body to create")
		return err
	}

	db.Create(&book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	if book.Title == "" {
		c.Status(500).SendString("No book found")
		return nil
	}
	db.Update(&book, id)
	return c.JSON(book)

}
func DeleteBook(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).SendString("No book found")
		return nil
	}
	db.Delete(&book)
	return c.SendString("Book successfully deleted")
}
