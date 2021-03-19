package routes

import (
	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/appointment"
)

func AppointmentRoutes(app *fiber.App) {

	// Appointment Endpoint
	app.Get("/api/v1/appointment", appointment.GetAppointment)
	//app.Get("/api/v1/appointment/:ci", doctor.GetDoctor)
	app.Post("/api/v1/appointment", appointment.AddAppointment)
	//app.Put("/api/v1/appointment/:ci", doctor.UpdateDoctor)
	//app.Delete("/api/v1/appointment/:ci", doctor.DeleteDoctor)
}
