package repo

import (
	"authentication/models"
	"authentication/utils"
	"authentication/utils/db"
)

func CreateUser(user models.User) error { //This is for SignUp
	query := `INSERT INTO users (username, password, email, phone_number, pancard) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.DB.Exec(query, user.UserName, user.Password, user.Email, user.PhoneNumber, user.PanCard)
	if err != nil {
		utils.ErrorLogger.Println("CreateUser error:", err)
	}
	return err
}

func ValidateUser(username, password string) (bool, error) { // This is for SignIn
	var dbPassword string
	query := `SELECT password FROM users WHERE username=$1`
	err := db.DB.QueryRow(query, username).Scan(&dbPassword)
	if err != nil {
		utils.ErrorLogger.Println("ValidateUser error:", err)
		return false, err
	}
	return dbPassword == password, nil
}

func IsUsernameTaken(username string) (bool, error) { // To check if the username is taken or not
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)`
	err := db.DB.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
