package service

import (
	"authentication/repo"
)

func LoginUser(username, password string) (bool, error){
	return repo.SignIn(username,password)
}