package user

import (
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"size:50;not null" json:"name"`
	CI   int    `gorm:"size:8;unique;not null" json:"ci"`
}

func GetUsers(c *fiber.Ctx) {
	db := data.DBConn
	var users []User
	if err := db.Find(&users).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.JSON(users)
}

func GetUser(c *fiber.Ctx) {
	params := c.Params("ci")
	db := data.DBConn
	ci, _ := strconv.Atoi(params)
	var user User
	if err := db.Where("ci = ?", ci).First(&user).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send("User not found")
		return
	}
	c.JSON(user)
}

func AddUser(c *fiber.Ctx) {
	db := data.DBConn
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	if user.Name == "" || len(user.Name) >= 40 {
		c.Status(fiber.StatusBadRequest).Send("Name can't be empty or have more that 50 char ")
		return
	}
	if user.CI == 0 || user.CI < 1000000 || user.CI > 100000000 {
		c.Status(fiber.StatusBadRequest).Send("CI can't be empty or CI is Invalid")
		return
	}
	if err := db.Create(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}
	c.Status(fiber.StatusCreated).JSON(user)
}

func UpdateUser(c *fiber.Ctx) {
	ci := c.Params("ci")
	db := data.DBConn

	var user User

	if err := db.Where("ci = ?", ci).First(&user).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send("User not found")
		return
	}

	err := c.BodyParser(&user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}

	if len(user.Name) >= 40 {
		c.Status(fiber.StatusBadRequest).Send("Name can't be empty or have more that 50 char ")
		return
	}
	if user.CI < 1000000 || user.CI > 100000000 {
		c.Status(fiber.StatusBadRequest).Send("CI can't be empty or CI is Invalid")
		return
	}
	if err := db.Model(User{}).Where("ci = ?", ci).Updates(User{Name: user.Name, CI: user.CI}).Error; err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}
	c.Status(fiber.StatusOK).Send("Updated User: ", ci)
}

func DeleteUser(c *fiber.Ctx) {
	ci := c.Params("ci")
	db := data.DBConn

	var user User
	db.Where("ci = ?", ci).First(&user)
	if user.Name == "" {
		c.Status(fiber.StatusNotFound).Send("Not user Found with CI ", ci)
		return
	}
	db.Delete(&user)
	c.Status(fiber.StatusOK).Send("User SuccessFully deleted")

}
