package models

type SignInForm struct {
	Username string `json:"username" validate:"nonzero,max=45,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password string `json:"password" validate:"nonzero,max=45"`
}

type SignUpForm struct {
	Firstname string `json:"first_name" validate:"nonzero,max=45"`
	Lastname string `json:"last_name" validate:"nonzero,max=45"`
	Username string `json:"username" validate:"nonzero,max=45,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password string `json:"password" validate:"nonzero,max=45"`
	PasswordConfirmation string `json:"password_confirmation" validate:"nonzero,max=45"`
}

type MailForm struct {
	Addressee string `json:"addressee" validate:"max=45"`
	Theme     string `json:"theme" validate:"max=45"`
	Text      string `json:"text"`
	Files     string `json:"files"`
}

type ProfileSettingsForm struct {
	Firstname string `json:"first_name" validate:"max=45"`
	Lastname string `json:"last_name" validate:"max=45"`
	Password string `json:"password" validate:"max=45"`
}