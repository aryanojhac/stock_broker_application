package service

import (
	"authentication/models"
	"authentication/repo"
	"authentication/constants"
)

func RegisterUser(user models.SignUpModel) error { //I will hash the password in this function
	// Check if username already exists
	exists, err := repo.IsUsernameTaken(user.UserName)
	if err != nil {
		return constants.ErrCheckUsername
	}
	if exists {
		return constants.ErrUsernameTaken
	}

// Check if email already exists
	emailExists, err := repo.IsEmailExists(user.Email)
	if err!= nil {
		return constants.ErrCheckEmail
	}
	if emailExists {
		return constants.ErrEmailTaken
	}

	// Create user if not exists
	return repo.CreateUser(user)
}