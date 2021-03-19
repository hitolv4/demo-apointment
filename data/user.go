package data

import "errors"

func GetUsers() ([]User, error) {
	db := DBConn
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(ci int) (*User, error) {
	db := DBConn
	var user = User{}
	if err := db.Where("ci = ?", ci).First(&user).Error; err != nil {
		return nil, ErrNotFount
	}
	return &user, nil
}

func AddUser(user User) (*User, error) {
	db := DBConn
	if err := db.Create(&user).Error; err != nil {

		return nil, err
	}
	return &user, nil
}

func UpdateUser(ci int, user User) error {
	db := DBConn

	if err := db.Model(&User{}).Where("ci = ?", ci).Updates(User{Name: user.Name, CI: user.CI}).Error; err != nil {

		return err
	}
	return nil
}

func DeleteUser(user User) error {
	db := DBConn
	err := db.Select("Appointment").Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

var ErrNotFount = errors.New("User not found")
