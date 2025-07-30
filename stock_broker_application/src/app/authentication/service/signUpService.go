package service

import (
	"errors"
	"authentication/models"
	"authentication/repo"
)

func RegisterUser(user models.SignUpModel) error { //I will hash the password in this function
	// Check if username already exists
	exists, err := repo.IsUsernameTaken(user.UserName)
	if err != nil {
		return err // Database error
	}
	if exists {
		return errors.New("username already taken")
	}

	// Create user if not exists
	return repo.CreateUser(user)
}
