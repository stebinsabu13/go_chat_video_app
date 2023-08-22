package utils

type BodySignUpuser struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	MobileNum       string `json:"mobilenum" binding:"required,min=10,max=10"`
	Password        string `json:"password" binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
}

type BodyLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
