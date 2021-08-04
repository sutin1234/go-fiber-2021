package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/sutin1234/go-fiber-2021/database"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user User
	db.Find(&user, id)
	return c.JSON(user)
}
func NewUser(c *fiber.Ctx) error {
	db := database.DBConn
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		c.Status(503).SendString("no body")
		return nil
	}

	db.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user User
	if user.FirstName == "" {
		c.Status(500).SendString("No User found")
		return nil
	}
	db.Update(&user, id)
	return c.JSON(user)

}
func DeleteUser(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var user User
	db.First(&user, id)
	if user.Email == "" {
		c.Status(500).SendString("No User found")
		return nil
	}
	db.Delete(&user)
	return c.SendString("User successfully deleted")
}
