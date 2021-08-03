package user

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/sutin1234/go-fiber-2021/database"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func GetUsers(c *fiber.Ctx) {
	db := database.DBConn
	var users []User
	db.Find(&users)
	c.JSON(users)
}
func GetUser(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var user User
	db.Find(&user, id)
	c.JSON(user)
}
func NewUser(c *fiber.Ctx) {
	db := database.DBConn
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&user)
	c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var user User
	if user.FirstName == "" {
		c.Status(500).Send("No User found")
		return
	}
	db.Update(&user, id)
	c.JSON(user)

}
func DeleteUser(c *fiber.Ctx) {
	db := database.DBConn
	id := c.Params("id")
	var User User
	db.First(&User, id)
	if User.Title == "" {
		c.Status(500).Send("No User found")
		return
	}
	db.Delete(&User)
	c.Send("User successfully deleted")
}
