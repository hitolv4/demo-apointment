package doctor

import (
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
)

func GetDoctors(c *fiber.Ctx) {

	doctors, err := data.GetDoctors()
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.JSON(doctors)
}
func GetDoctor(c *fiber.Ctx) {
	params := c.Params("ci")

	ci, _ := strconv.Atoi(params)
	doctor, err := data.GetDoctor(ci)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.JSON(doctor)
}

func AddDoctor(c *fiber.Ctx) {
	doctor := new(data.Doctor)
	if err := c.BodyParser(doctor); err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	if doctor.Name == "" || len(doctor.Name) >= 40 {
		c.Status(fiber.StatusBadRequest).Send("Name can't be empty or have more that 50 char ")
		return
	}
	if doctor.CI == 0 || doctor.CI < 1000000 || doctor.CI > 100000000 {
		c.Status(fiber.StatusBadRequest).Send("CI can't be empty or CI is Invalid")
		return
	}
	newDoctor, err := data.AddDoctor(*doctor)
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}

	c.Status(fiber.StatusCreated).JSON(newDoctor)
}

func UpdateDoctor(c *fiber.Ctx) {
	params := c.Params("ci")

	ci, _ := strconv.Atoi(params)
	doctor, err := data.GetDoctor(ci)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	err = c.BodyParser(&doctor)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	err = data.UpdateDoctor(ci, *doctor)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	c.Status(fiber.StatusOK).Send("Updated Doctor: ", ci)

}
func DeleteDoctor(c *fiber.Ctx) {
	params := c.Params("ci")

	ci, _ := strconv.Atoi(params)
	doctor, err := data.GetDoctor(ci)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	err = data.DeleteDoctor(*doctor)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.Status(fiber.StatusOK).Send("Doctor SuccessFully deleted")

}
