package services

import (
	"errors"
	"time"

	"tma.com.vn/lngochuy/rest/models"
	"tma.com.vn/lngochuy/rest/resources"
)

func GetUserList() ([]resources.UserResponse, error) {
	var users = make([]resources.UserResponse, 0)

	err := models.GetDB().
		Model(&models.User{}).
		Select("id, email, date_of_birth").
		Find(&users).Error

	if err != nil {
		return users, errors.New("can't get all users")
	}

	return users, nil
}

func GetUser(id uint) (*resources.UserResponse, error) {
	var user = new(resources.UserResponse)

	err := models.GetDB().
		Model(&models.User{}).
		Select("ID, email, date_of_birth").
		First(&user, id).Error

	if err != nil {
		return user, errors.New("can't get user")
	}

	return user, nil
}

func CreateUser(email string, dateOfBirth time.Time) (uint, error) {
	user := &models.User{
		Email:       email,
		DateOfBirth: dateOfBirth,
	}

	if err := models.GetDB().Create(&user).Error; err != nil {
		return 0, errors.New("can't create user")
	}

	return user.ID, nil
}

func UpdateUser(id uint, email string, dateOfBirth time.Time) error {
	if models.GetDB().Model(&models.User{ID: id}).
		Updates(models.User{
			Email:       email,
			DateOfBirth: dateOfBirth,
		}).RowsAffected == 0 {
		return errors.New("can't update a void object")
	}

	return nil
}

func RemoveUser(id uint) error {
	if models.GetDB().Delete(&models.User{ID: id}).RowsAffected == 0 {
		return errors.New("can't delete a void object")
	}

	return nil
}
