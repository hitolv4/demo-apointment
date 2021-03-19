package appointment

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
)

func GetAppointment(c *fiber.Ctx) {
	db := data.DBConn
	var appointment []data.Appointment
	if err := db.Find(&appointment).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.JSON(appointment)
}

func AddAppointment(c *fiber.Ctx) {

	db := data.DBConn
	appointment := new(data.Appointment)
	if err := c.BodyParser(appointment); err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}
	var user data.User
	if err := db.Where("id = ?", appointment.UserID).First(&user).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send("Must enter a Active user")
		return
	}

	var doctor data.Doctor
	if err := db.Where("id = ?", appointment.DoctorID).First(&doctor).Error; err != nil {
		c.Status(fiber.StatusNotFound).Send("Must enter a Active Doctor")
		return
	}

	if time.Now() == appointment.Appointment {
		c.Status(fiber.StatusNotFound).Send("you can't set the apointment the same date")
		return
	}

	if err := db.Create(&appointment).Error; err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}

	c.Status(fiber.StatusCreated).JSON(appointment)
}
