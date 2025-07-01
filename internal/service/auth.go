package service

import (
	"errors"
	"pet_shelter_and_store/internal/errs"
	"pet_shelter_and_store/internal/models"
	"pet_shelter_and_store/internal/repository"
	"pet_shelter_and_store/utils"
)

func SignIn(userDataCheck, password string) (user models.Users, accessToken string, refreshToken string, err error) {
	if userDataCheck == "" {
		return user, "", "", errs.ErrInvalidData
	}

	user, err = repository.GetUserByUsernameAndPassword(userDataCheck, password)
	if err != nil {
		if !errors.Is(err, errs.ErrRecordNotFound) {
			return user, "", "", err
		}

		return user, "", "", errs.ErrInvalidCredentials
	}

	accessToken, refreshToken, err = utils.GenerateToken(uint(user.ID), user.Role)
	if err != nil {
		return user, "", "", err
	}

	return user, accessToken, refreshToken, nil
}
