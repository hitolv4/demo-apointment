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
	UserID      int       `gorm:"size:3;not null" json:"userdId"`
	DoctorID    int       `gorm:"size:3;not null" json:"doctordId"`
	Appointment time.Time `gorm:"not null" json:"appointment"`
}

type Doctor struct {
	gorm.Model
	Name        string        `gorm:"size:50;not null" json:"name"`
	CI          int           `gorm:"size:8;unique;not null" json:"ci"`
	Appointment []Appointment `gorm:"foreignkey:DoctorID" json:"-"`
}
