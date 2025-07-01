package service

import (
	"fmt"
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/repository"
	"pet_shelter_and_store/logger"
	"pet_shelter_and_store/utils"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.Users, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (uint, error) {
	usernameExists, err := repository.UserExists(user.Username)
	if err != nil {
		return 0, fmt.Errorf("failed to check existing user: %w", err)
	}

	if user.Password == "" || user.Username == "" {
		return 0, errs.ErrInvalidData
	}

	if usernameExists {
		logger.Error.Printf("user with username %s already exists", user.Username)
		return 0, errs.ErrUsernameUniquenessFailed
	}

	user.Password = utils.GenerateHash(user.Password)

	var userDB models.User

	if userDB, err = repository.CreateUser(user); err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return uint(userDB.ID), nil
}
