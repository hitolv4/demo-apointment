package data

func GetAppointment() ([]Appointment, error) {
	db := DBConn
	var appointments []Appointment
	if err := db.Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

func AddAppointment(appointment Appointment) (*Appointment, error) {
	db := DBConn
	if err := db.Create(&appointment).Error; err != nil {

		return nil, err
	}
	return &appointment, nil
}
