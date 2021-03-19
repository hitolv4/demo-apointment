package doctor

import (
	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
)

func GetDoctors(c *fiber.Ctx) {
	db := data.DBConn
	var doctors []data.Doctor
	if err := db.Find(&doctors).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send("User ID doesn't")
		return
	}
	c.JSON(doctors)
}

func GetDoctor(c *fiber.Ctx) {
	ci := c.Params("ci")
	db := data.DBConn
	var doctor data.Doctor
	if err := db.Where("ci = ?", ci).First(&doctor).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send("User not found")
		return
	}
	c.JSON(doctor)
}

func AddDoctor(c *fiber.Ctx) {
	db := data.DBConn
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

	if err := db.Create(&doctor).Error; err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}
	c.Status(fiber.StatusCreated).JSON(doctor)
}

func UpdateDoctor(c *fiber.Ctx) {
	ci := c.Params("ci")
	db := data.DBConn

	var doctor data.Doctor
	if err := db.Where("ci = ?", ci).First(&doctor).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send("User not found")
		return
	}

	err := c.BodyParser(&doctor)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	if len(doctor.Name) >= 40 {
		c.Status(fiber.StatusBadRequest).Send("Name can't be empty or have more that 50 char ")
		return
	}
	if doctor.CI < 1000000 || doctor.CI > 100000000 {
		c.Status(fiber.StatusBadRequest).Send("CI can't be empty or CI is Invalid")
		return
	}
	if err := db.Model(data.Doctor{}).Where("ci = ?", ci).Updates(data.Doctor{Name: doctor.Name, CI: doctor.CI}).Error; err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}

	c.Status(fiber.StatusOK).Send("Updated Doctor: ", ci)
}

func DeleteDoctor(c *fiber.Ctx) {
	ci := c.Params("ci")
	db := data.DBConn

	var doctor data.Doctor
	db.Where("ci = ?", ci).First(&doctor)
	if doctor.Name == "" {
		c.Status(fiber.StatusNotFound).Send("Not doctor Found with CI ", ci)
		return
	}
	db.Delete(&doctor)
	c.Status(fiber.StatusOK).Send("Doctor SuccessFully deleted")

}
