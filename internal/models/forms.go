package models

type SignInForm struct {
	Email		string	`json:"email" validate:"nonzero,max=20,regexp=^(?=[a-zA-Z0-9._]$)(?!.*[_.]{2})[^_.].*[^_.]$"`
	Password	string	`json:"password" validate:"nonzero,max=20"`
}

type SignUpForm struct {
	FirstName		string	`json:"first_name" validate:"nonzero,max=20"`
	LastName		string	`json:"last_name" validate:"nonzero,max=20"`
	Email			string	`json:"email" validate:"nonzero,max=20,regexp=^(?=[a-zA-Z0-9._]$)(?!.*[_.]{2})[^_.].*[^_.]$"`
	Password		string	`json:"password" validate:"nonzero,max=20"`
	PasswordConf	string	`json:"password_confirmation" validate:"nonzero,max=20"`
}
