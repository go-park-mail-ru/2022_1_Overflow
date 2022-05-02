package models

type SignInForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpForm struct {
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Username string `json:"username"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type MailForm struct {
	Addressee string `json:"addressee"`
	Theme     string `json:"theme"`
	Text      string `json:"text"`
	Files     string `json:"files"`
}

type ProfileSettingsForm struct {
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Password string `json:"password"`
}