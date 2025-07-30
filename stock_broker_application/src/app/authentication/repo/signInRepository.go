package repo
import (
	"authentication/utils"
	"authentication/utils/db"
)


func SignIn(username, password string) (bool, error) { 
	var dbPassword string
	query := `SELECT password FROM users WHERE username=$1`
	err := db.DB.QueryRow(query, username).Scan(&dbPassword)
	if err != nil {
		utils.ErrorLogger.Println("ValidateUser error:", err)
		return false, err
	}
	return dbPassword == password, nil
}