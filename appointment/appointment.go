package appointment

import (
	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
)

func GetAppointment(c *fiber.Ctx) {
	appointments, err := data.GetAppointment()
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
		return
	}
	c.JSON(appointments)
}

func AddAppointment(c *fiber.Ctx) {

	appointment := new(data.Appointment)
	if err := c.BodyParser(appointment); err != nil {
		c.Status(fiber.StatusInternalServerError).Send(err)
		return
	}

	_, err := data.GetUser(appointment.UserID)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
	}

	_, err = data.GetDoctor(appointment.DoctorID)
	if err != nil {
		c.Status(fiber.StatusNotFound).Send(err)
	}

	newAppointment, err := data.AddAppointment(*appointment)
	if err != nil {
		c.Status(fiber.StatusBadRequest).Send(err)
		return
	}

	c.Status(fiber.StatusCreated).JSON(newAppointment)
}
