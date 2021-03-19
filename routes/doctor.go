package routes

import (
	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/doctor"
)

func DoctorRoutes(app *fiber.App) {

	// doctor endpoint
	app.Get("/api/v1/doctor", doctor.GetDoctors)
	app.Get("/api/v1/doctor/:ci", doctor.GetDoctor)
	app.Post("/api/v1/doctor", doctor.AddDoctor)
	app.Put("/api/v1/doctor/:ci", doctor.UpdateDoctor)
	app.Delete("/api/v1/doctor/:ci", doctor.DeleteDoctor)

}
