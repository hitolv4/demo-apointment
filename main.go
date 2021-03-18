package main

import (
	"fmt"

	"github.com/hitolv4/apointment/doctor"
	"github.com/hitolv4/apointment/user"

	"github.com/gofiber/fiber"
	"github.com/hitolv4/apointment/data"
	"github.com/hitolv4/apointment/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()
	InitDatabase()
	routes.SetupRoutes(app)
	app.Listen(3000)
}

func InitDatabase() {
	var err error
	data.DBConn, err = gorm.Open(sqlite.Open("appointment.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	data.DBConn.AutoMigrate(&user.User{})
	data.DBConn.AutoMigrate(&doctor.Doctor{})
	fmt.Println("Database Migrated")
}
