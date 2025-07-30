package models

type SignUpModel struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PanCard     string `json:"pancard"`
}
