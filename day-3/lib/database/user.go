package database

import (
	"alterra-agmc-day-3/config"
	"alterra-agmc-day-3/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserById(id int) (models.User, error) {
	var user models.User

	if e := config.DB.First(&user, id).Error; e != nil {
		return models.User{}, e
	}
	return user, nil
}

func CreateUser(user models.User) (models.User, error) {
	if e := config.DB.Create(&user).Error; e != nil {
		return models.User{}, e
	}
	return user, nil
}

func UpdateUser(user models.User) error {
	if e := config.DB.Model(user).Updates(user).Error; e != nil {
		return e
	}
	return nil
}

func DeleteUser(id int) error {
	if e := config.DB.Delete(&models.User{}, id).Error; e != nil {
		return e
	}
	return nil
}

func LoginUser(user models.User) (models.User, error) {
	if e := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; e != nil {
		return models.User{}, e
	}
	return user, nil
}
