package models

type SignInForm struct {
	Username string `json:"username" validate:"nonzero,max=45,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password string `json:"password" validate:"nonzero,max=45"`
}

type SignUpForm struct {
	FirstName    string `json:"first_name" validate:"nonzero,max=45"`
	LastName     string `json:"last_name" validate:"nonzero,max=45"`
	Username     string `json:"username" validate:"nonzero,max=45,regexp=^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"`
	Password     string `json:"password" validate:"nonzero,max=45"`
	PasswordConf string `json:"password_confirmation" validate:"nonzero,max=45"`
}

type SettingsForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type MailForm struct {
	Addressee string `json:"addressee"`
	Theme     string `json:"theme"`
	Text      string `json:"text"`
	Files     string `json:"files"`
}