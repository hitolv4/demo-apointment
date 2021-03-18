package routes

import (
	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/doctor"
	"github.com/hitolv4/apointment/user"
)

func SetupRoutes(app *fiber.App) {

	// user endpoint
	app.Get("/api/v1/user", user.GetUsers)
	app.Get("/api/v1/user/:ci", user.GetUser)
	app.Post("/api/v1/user", user.AddUser)
	app.Put("/api/v1/user/:ci", user.UpdateUser)
	app.Delete("/api/v1/user/:ci", user.DeleteUser)

	// doctor endpoint
	app.Get("/api/v1/doctor", doctor.GetDoctors)
	app.Get("/api/v1/doctor/:ci", doctor.GetDoctor)
	app.Post("/api/v1/doctor", doctor.AddDoctor)
	app.Put("/api/v1/doctor/:ci", doctor.UpdateDoctor)
	app.Delete("/api/v1/doctor/:ci", doctor.DeleteDoctor)
}
