package db

import (
	"github.com/mstolin/wishlist-inviter/utils/models"
)

func (dbHandler DatabaseHandler) CreateUser(user *models.User) (models.User, error) {
	if err := dbHandler.DB.Create(&user).Error; err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}

func (dbHandler DatabaseHandler) GetUserById(id string) (models.User, error) {
	user := models.User{}
	if err := dbHandler.DB.Preload("Items").Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (dbHandler DatabaseHandler) DeleteUserById(id string) (models.User, error) {
	user, err := dbHandler.GetUserById(id)
	if err != nil {
		return user, err
	}

	if err = dbHandler.DB.Delete(&user).Error; err != nil {
		return user, err
	} else {
		return user, nil
	}
}
