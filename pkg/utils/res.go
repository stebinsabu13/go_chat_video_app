package utils

type ResponseUsers struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobilenumber"`
	Password     string `json:"password"`
}
