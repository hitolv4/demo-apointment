package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
	"github.com/hitolv4/apointment/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	InitDatabase()
	routes.UserRoutes(app)
	routes.DoctorRoutes(app)
	routes.AppointmentRoutes(app)
	app.Listen(3000)
}

func InitDatabase() {
	var err error
	data.DBConn, err = gorm.Open(sqlite.Open("appointment.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	data.DBConn.AutoMigrate(&data.User{})
	data.DBConn.AutoMigrate(&data.Doctor{})
	data.DBConn.AutoMigrate(&data.Appointment{})
	fmt.Println("Database Migrated")
}
