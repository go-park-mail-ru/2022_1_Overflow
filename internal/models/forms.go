package models

type SignInForm struct {
	Email		string	`json:"email" validate:"nonzero,max=20,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password	string	`json:"password" validate:"nonzero,max=20"`
}

type SignUpForm struct {
	FirstName		string	`json:"first_name" validate:"nonzero,max=20"`
	LastName		string	`json:"last_name" validate:"nonzero,max=20"`
	Email			string	`json:"email" validate:"nonzero,max=20,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password		string	`json:"password" validate:"nonzero,max=20"`
	PasswordConf	string	`json:"password_confirmation" validate:"nonzero,max=20"`
}

type SettingsForm struct {
	FirstName	string	`json:"first_name"`	
	LastName	string	`json:"last_name"`
	Password	string	`json:"password"`
}

type MailForm struct {
	Addressee string    `json:"addressee"`
	Theme     string    `json:"theme"`
	Text      string    `json:"text"`
	Files     string    `json:"files"`
}