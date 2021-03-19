package data

import "errors"

func GetDoctors() ([]Doctor, error) {
	db := DBConn
	var doctors []Doctor
	if err := db.Find(&doctors).Error; err != nil {
		return nil, err
	}
	return doctors, nil
}

func GetDoctor(ci int) (*Doctor, error) {
	db := DBConn
	var doctor = Doctor{}
	if err := db.Where("ci = ?", ci).First(&doctor).Error; err != nil {
		return nil, ErrDoctorNotFount
	}
	return &doctor, nil
}

func AddDoctor(doctor Doctor) (*Doctor, error) {
	db := DBConn
	if err := db.Create(&doctor).Error; err != nil {

		return nil, err
	}
	return &doctor, nil
}

func UpdateDoctor(ci int, doctor Doctor) error {
	db := DBConn

	if err := db.Model(&Doctor{}).Where("ci = ?", ci).Updates(Doctor{Name: doctor.Name, CI: doctor.CI}).Error; err != nil {

		return err
	}
	return nil
}

func DeleteDoctor(doctor Doctor) error {
	db := DBConn
	err := db.Select("Appointment").Delete(&doctor).Error
	if err != nil {
		return err
	}
	return nil
}

var ErrDoctorNotFount = errors.New("Doctor not found")
