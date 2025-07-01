package repository

import (
	"errors"
	"gorm.io/gorm"
	"pet_shelter_and_store/internal/db"
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/logger"
)

func GetAllUsers() (users []models.Users, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, TranslateGormError(err)
	}

	return users, nil
}

func GetUserByID(id uint) (user models.Users, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, TranslateGormError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.Users, error) {
	var user models.Users
	err := db.GetDBConn().Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrUserNotFound
		}
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return nil, err
	}
	return &user, nil
}

func UserExists(username string) (bool, error) {
	users, err := GetAllUsers()
	if err != nil {
		return false, err
	}

	var usernameExists bool
	for _, user := range users {
		if user.Username == username {
			usernameExists = true
		}
	}
	return usernameExists, nil
}

func CreateUser(user models.Users) (userDB models.Users, err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return userDB, TranslateGormError(err)
	}

	userDB = user
	return userDB, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.Users, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}

func GetUserByEmailAndPassword(password string) (user models.Users, err error) {
	err = db.GetDBConn().Where("password = ?", password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByEmailAndPassword] error getting user by email and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}

func GetUserByEmailPasswordAndUsername(username, email, password string) (user models.Users, err error) {
	err = db.GetDBConn().Where("password = ? AND username = ?", email, password, username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByEmailPasswordAndUsername] error getting user by username, email and password: %v\n", err)
		return user, TranslateGormError(err)
	}

	return user, nil
}
