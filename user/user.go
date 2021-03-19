package user

import (
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
)

func GetUsers(c *fiber.Ctx) {

	users, err := data.GetUsers()
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.JSON(users)
}
func GetUser(c *fiber.Ctx) {
	params := c.Params("ci")

	ci, _ := strconv.Atoi(params)
	user, err := data.GetUser(ci)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send("User not found")
		return
	}
	c.JSON(user)
}

func AddUser(c *fiber.Ctx) {
	user := new(data.User)
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
	newUser, err := data.AddUser(*user)
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}

	c.Status(fiber.StatusCreated).JSON(newUser)
}

func UpdateUser(c *fiber.Ctx) {
	params := c.Params("ci")

	ci, _ := strconv.Atoi(params)
	user, err := data.GetUser(ci)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	err = c.BodyParser(&user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	err = data.UpdateUser(ci, *user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	c.Status(fiber.StatusOK).Send("Updated User: ", ci)

}
func DeleteUser(c *fiber.Ctx) {
	params := c.Params("ci")

	ci, _ := strconv.Atoi(params)
	user, err := data.GetUser(ci)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	err = data.DeleteUser(*user)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.Status(fiber.StatusOK).Send("User SuccessFully deleted")

}
