package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/sutin1234/go-fiber-2021/book"
	"github.com/sutin1234/go-fiber-2021/database"
	"github.com/sutin1234/go-fiber-2021/user"
)

const contextAPI = "/api/v2"
const bookContextAPI = contextAPI + "/book"
const userContextAPI = contextAPI + "/user"

func main() {
	app := fiber.New()
	app.Get("/", helloWorld)

	// initialized database
	initDatabase()
	defer database.DBConn.Close()

	// setup routes
	setupBookRoutes(app)
	setupUserRoutes(app)

	// serve fiber
	app.Listen(":3000")
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully opened!")

	// database migration
	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func setupBookRoutes(app *fiber.App) {
	app.Get(bookContextAPI, book.GetBooks)
	app.Get(bookContextAPI+"/:id", book.GetBook)
	app.Post(bookContextAPI, book.NewBook)
	app.Delete(bookContextAPI+"/:id", book.DeleteBook)
	app.Put(bookContextAPI+"/:id", book.UpdateBook)
}

func setupUserRoutes(app *fiber.App) {
	app.Get(userContextAPI, user.GetUsers)
	app.Get(userContextAPI+"/:id", user.GetUser)
	app.Post(userContextAPI, user.NewUser)
	app.Delete(userContextAPI+"/:id", user.DeleteUser)
	app.Put(userContextAPI+"/:id", user.UpdateUser)
}
