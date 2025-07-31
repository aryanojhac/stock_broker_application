package repo

import (
	"authentication/models"
	"authentication/utils"
	"authentication/utils/db"
)

func CreateUser(user models.SignUpModel) error { //This is for SignUp
	query := `INSERT INTO users (username, password, email, phone_number, pancard) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.DB.Exec(query, user.UserName, user.Password, user.Email, user.PhoneNumber, user.PanCard)
	if err != nil {
		utils.ErrorLogger.Println("CreateUser error:", err)
	}
	return err
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

func IsEmailExists(email string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)`
	err := db.DB.QueryRow(query, email).Scan(&exists)
	if err != nil {
		utils.ErrorLogger.Println("IsEmailExists error:", err)
		return false, err
	}
	return exists, nil
}
