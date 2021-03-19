package routes

import (
	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/user"
)

func UserRoutes(app *fiber.App) {

	// user endpoint
	app.Get("/api/v1/user", user.GetUsers)
	app.Get("/api/v1/user/:ci", user.GetUser)
	app.Post("/api/v1/user", user.AddUser)
	app.Put("/api/v1/user/:ci", user.UpdateUser)
	app.Delete("/api/v1/user/:ci", user.DeleteUser)

}
