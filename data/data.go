package data

import (
	"time"

	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

type User struct {
	gorm.Model
	Name        string        `gorm:"size:50;not null" json:"name"`
	CI          int           `gorm:"size:8;unique;not null" json:"ci"`
	Appointment []Appointment `gorm:"foreignkey:UserID" json:"-"`
}

type Appointment struct {
	gorm.Model
	UserID      int
	DoctorID    int
	Appointment time.Time
}

type Doctor struct {
	gorm.Model
	Name        string        `gorm:"size:50;not null" json:"name"`
	CI          int           `gorm:"size:8;unique;not null" json:"ci"`
	Appointment []Appointment `gorm:"foreignkey:DoctorID" json:"-"`
}
