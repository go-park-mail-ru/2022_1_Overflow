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
}

type ChangePasswordForm struct {
	OldPassword string `json:"password_old" validate:"max=45"`
	NewPassword string `json:"password_new" validate:"max=45"`
	NewPasswordConf string `json:"password_new_confirmation" validate:"max=45"`
}

type ReadMailForm struct {
	Id int32 `json:"id"`
	IsRead bool `json:"isread"`
}

type DeleteMailForm struct {
	Id int32 `json:"id"`
}

type AddFolderForm struct {
	FolderName string `json:"folder_name"`
}

type AddMailToFolderForm struct {
	FolderName string `json:"folder_name"`
	MailId int32 `json:"mail_id"`
	Move bool `json:"move"`
}

type ChangeFolderForm struct {
	FolderName string `json:"folder_name"`
	NewFolderName string `json:"new_folder_name"`
}

type DeleteFolderForm struct {
	FolderName string `json:"folder_name"`
}

type DeleteFolderMailForm struct {
	FolderName string `json:"folder_name"`
	MailId int32 `json:"mail_id"`
	Restore bool `json:"restore"`
}